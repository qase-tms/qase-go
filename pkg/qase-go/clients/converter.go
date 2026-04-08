package clients

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
	"github.com/qase-tms/qase-go/pkg/qase-go/logging"
	api_v2_client "github.com/qase-tms/qase-go/qase-api-v2-client"
)

// AttachmentUploader interface for uploading attachments
// Returns a slice of hashes corresponding to the uploaded files
type AttachmentUploader interface {
	UploadAttachment(ctx context.Context, projectCode string, file []*os.File) ([]string, error)
}

// v1ClientAdapter adapts V1Client to AttachmentUploader interface
type v1ClientAdapter struct {
	client *V1Client
}

// UploadAttachment implements AttachmentUploader interface
func (a *v1ClientAdapter) UploadAttachment(ctx context.Context, projectCode string, file []*os.File) ([]string, error) {
	return a.client.uploadAttachmentsInternal(ctx, projectCode, file)
}

// NewV1ClientAdapter creates an adapter for V1Client to implement AttachmentUploader
func NewV1ClientAdapter(client *V1Client) AttachmentUploader {
	return &v1ClientAdapter{client: client}
}

// V2Converter handles conversion from domain models to API v2 models
type V2Converter struct {
	uploader    AttachmentUploader
	projectCode string
}

// NewV2Converter creates a new V2 converter instance
func NewV2Converter() *V2Converter {
	return &V2Converter{}
}

// NewV2ConverterWithUploader creates a new V2 converter instance with attachment uploader
func NewV2ConverterWithUploader(uploader AttachmentUploader, projectCode string) *V2Converter {
	return &V2Converter{
		uploader:    uploader,
		projectCode: projectCode,
	}
}

// ConvertTestResult converts domain.TestResult to api_v2_client.ResultCreate
func (c *V2Converter) ConvertTestResult(result *domain.TestResult) (*api_v2_client.ResultCreate, error) {
	if result == nil {
		return nil, fmt.Errorf("test result cannot be nil")
	}

	// Create execution object
	execution := api_v2_client.NewResultExecution(string(result.Execution.Status))

	if result.Execution.StartTime != nil {
		// Convert milliseconds to seconds with fractional part
		startTime := *result.Execution.StartTime
		var startTimeSeconds float64
		if startTime > 1e12 { // If time is in milliseconds (13+ digits)
			startTimeSeconds = float64(startTime) / 1000.0
		} else {
			startTimeSeconds = float64(startTime)
		}

		// Ensure time is not in the future (API requirement)
		now := float64(time.Now().Unix())
		if startTimeSeconds > now {
			startTimeSeconds = now - 1.0 // Set to 1 second ago
		}

		execution.SetStartTime(startTimeSeconds)
	}

	if result.Execution.EndTime != nil {
		// Convert milliseconds to seconds with fractional part
		endTime := *result.Execution.EndTime
		var endTimeSeconds float64
		if endTime > 1e12 { // If time is in milliseconds (13+ digits)
			endTimeSeconds = float64(endTime) / 1000.0
		} else {
			endTimeSeconds = float64(endTime)
		}

		// Ensure time is not in the future (API requirement)
		now := float64(time.Now().Unix())
		if endTimeSeconds > now {
			endTimeSeconds = now
		}

		execution.SetEndTime(endTimeSeconds)
	}

	if result.Execution.Duration != nil {
		// Duration should remain in milliseconds
		execution.SetDuration(*result.Execution.Duration)
	}

	if result.Execution.Stacktrace != nil {
		execution.SetStacktrace(*result.Execution.Stacktrace)
	}

	if result.Execution.Thread != nil {
		execution.SetThread(*result.Execution.Thread)
	}

	// Create main result object
	apiResult := api_v2_client.NewResultCreate(result.Title, *execution)

	// Set optional fields
	if result.ID != "" {
		apiResult.SetId(result.ID)
	}

	if result.Signature != "" {
		apiResult.SetSignature(result.Signature)
	}

	// Handle TestopsID (can be single int64, []int64, or nil)
	if result.TestopsID != nil {
		if err := c.setTestopsID(apiResult, result.TestopsID); err != nil {
			// Log warning but don't fail the entire result
			logging.Warn("Warning: Failed to set TestopsID: %v, continuing without TestopsID", err)
			// Continue without TestopsID instead of failing
		}
	}

	// Convert fields
	if len(result.Fields) > 0 {
		if err := c.setFields(apiResult, result.Fields); err != nil {
			// Log warning but don't fail the entire result
			logging.Warn("Warning: Failed to set fields: %v, continuing without fields", err)
			// Continue without fields instead of failing
		}
	}

	// Convert tags
	if len(result.Tags) > 0 {
		c.setTags(apiResult, result.Tags)
	}

	// Convert attachments
	if len(result.Attachments) > 0 {
		if err := c.setAttachments(context.Background(), apiResult, result.Attachments); err != nil {
			// Log warning but don't fail the entire result
			logging.Warn("Warning: Failed to set attachments: %v, continuing without attachments", err)
			// Continue without attachments instead of failing
		}
	}

	// Convert steps
	if len(result.Steps) > 0 {
		apiSteps, err := c.convertSteps(result.Steps)
		if err != nil {
			// Log warning but don't fail the entire result
			logging.Warn("Warning: Failed to convert steps: %v, continuing without steps", err)
			// Continue without steps instead of failing
		} else {
			apiResult.SetSteps(apiSteps)

			// Set steps type based on step content
			if len(apiSteps) > 0 {
				stepsType := api_v2_client.CLASSIC
				apiResult.SetStepsType(stepsType)
			}
		}
	}

	// Convert params
	if len(result.Params) > 0 {
		c.setParams(apiResult, result.Params)
	}

	// Convert group params (flatten to param_groups format)
	if len(result.GroupParams) > 0 {
		c.setParamGroups(apiResult, result.GroupParams)
	}

	// Convert relations
	if result.Relations != nil && result.Relations.Suite != nil {
		if err := c.setRelations(apiResult, result.Relations); err != nil {
			// Log warning but don't fail the entire result
			logging.Warn("Warning: Failed to set relations: %v, continuing without relations", err)
			// Continue without relations instead of failing
		}
	}

	// Set message
	if result.Message != nil {
		apiResult.SetMessage(*result.Message)
	}

	return apiResult, nil
}

// ConvertTestStep converts domain.TestStep to api_v2_client.ResultStep
func (c *V2Converter) ConvertTestStep(step *domain.TestStep) (*api_v2_client.ResultStep, error) {
	if step == nil {
		return nil, fmt.Errorf("test step cannot be nil")
	}

	// Create step data
	stepData := api_v2_client.NewResultStepData(step.Data.Action)

	if step.Data.ExpectedResult != nil {
		stepData.SetExpectedResult(*step.Data.ExpectedResult)
	}

	if step.Data.Data != nil {
		stepData.SetInputData(*step.Data.Data)
	}

	// Convert attachments for step data
	if len(step.Attachments) > 0 {
		attachmentIDs, err := c.processAttachmentsGracefully(context.Background(), step.Attachments)
		if err != nil {
			// Log warning but don't fail the entire step
			logging.Warn("Warning: Some step attachments failed to process: %v", err)
			// Continue with available attachments
		}

		if len(attachmentIDs) > 0 {
			stepData.SetAttachments(attachmentIDs)
		}
	}

	// Create step execution
	execution := api_v2_client.NewResultStepExecution(api_v2_client.ResultStepStatus(step.Execution.Status))

	if step.Execution.StartTime != nil {
		// Convert milliseconds to seconds
		execution.SetStartTime(float64(*step.Execution.StartTime) / 1000.0)
	}

	if step.Execution.EndTime != nil {
		// Convert milliseconds to seconds
		execution.SetEndTime(float64(*step.Execution.EndTime) / 1000.0)
	}

	if step.Execution.Duration != nil {
		// Duration should remain in milliseconds
		execution.SetDuration(*step.Execution.Duration)
	}

	// Create main step object
	apiStep := api_v2_client.NewResultStep()
	apiStep.SetData(*stepData)
	apiStep.SetExecution(*execution)

	// Convert nested steps recursively
	if len(step.Steps) > 0 {
		var nestedSteps []map[string]interface{}
		var nestedStepErrors []string

		for i, nestedStep := range step.Steps {
			nestedApiStep, err := c.ConvertTestStep(&nestedStep)
			if err != nil {
				// Log warning but don't fail the entire step
				logging.Warn("Warning: Failed to convert nested step %d '%s': %v, skipping nested step", i, nestedStep.Data.Action, err)
				nestedStepErrors = append(nestedStepErrors, fmt.Sprintf("nested step %d '%s': %v", i, nestedStep.Data.Action, err))
				continue // Skip this nested step and continue with others
			}
			// Convert to map[string]interface{} as required by API
			stepMap := map[string]interface{}{
				"data":      nestedApiStep.GetData(),
				"execution": nestedApiStep.GetExecution(),
			}
			if len(nestedApiStep.GetSteps()) > 0 {
				stepMap["steps"] = nestedApiStep.GetSteps()
			}
			nestedSteps = append(nestedSteps, stepMap)
		}

		// Log summary of nested step conversion results
		if len(nestedStepErrors) > 0 {
			logging.Warn("Warning: %d nested steps had conversion errors and were skipped: %v", len(nestedStepErrors), nestedStepErrors)
		}

		// Set nested steps only if we have any successful conversions
		if len(nestedSteps) > 0 {
			apiStep.SetSteps(nestedSteps)
		}
	}

	return apiStep, nil
}

// setTestopsID handles different types of TestopsID
func (c *V2Converter) setTestopsID(apiResult *api_v2_client.ResultCreate, testopsID interface{}) error {
	switch v := testopsID.(type) {
	case int64:
		apiResult.SetTestopsId(v)
	case []int64:
		apiResult.SetTestopsIds(v)
	case float64:
		// Handle JSON unmarshaling edge case
		apiResult.SetTestopsId(int64(v))
	case int:
		// Handle int type
		apiResult.SetTestopsId(int64(v))
	case []int:
		// Handle []int type
		ids := make([]int64, len(v))
		for i, id := range v {
			ids[i] = int64(id)
		}
		apiResult.SetTestopsIds(ids)
	default:
		return fmt.Errorf("unsupported TestopsID type: %T", v)
	}
	return nil
}

// setFields sets result fields based on known field types
func (c *V2Converter) setFields(apiResult *api_v2_client.ResultCreate, fields map[string]string) error {
	resultFields := api_v2_client.NewResultCreateFields()

	// Initialize AdditionalProperties map
	additionalProperties := make(map[string]interface{})

	for key, value := range fields {
		switch key {
		case "description":
			resultFields.SetDescription(value)
		case "preconditions":
			resultFields.SetPreconditions(value)
		case "postconditions":
			resultFields.SetPostconditions(value)
		case "layer":
			resultFields.SetLayer(value)
		case "severity":
			resultFields.SetSeverity(value)
		case "priority":
			resultFields.SetPriority(value)
		case "behavior":
			resultFields.SetBehavior(value)
		case "type":
			resultFields.SetType(value)
		case "muted":
			resultFields.SetMuted(value)
		case "is_flaky":
			resultFields.SetIsFlaky(value)
		case "executed_by":
			resultFields.SetExecutedBy(value)
		case "tags":
			resultFields.SetTags(value)
		case "author":
			resultFields.SetAuthor(value)
		default:
			// Add unknown fields to AdditionalProperties
			additionalProperties[key] = value
		}
	}

	// Set AdditionalProperties if there are any
	if len(additionalProperties) > 0 {
		resultFields.AdditionalProperties = additionalProperties
	}

	apiResult.SetFields(*resultFields)
	return nil
}

// setAttachments converts domain attachments to API format
func (c *V2Converter) setAttachments(ctx context.Context, apiResult *api_v2_client.ResultCreate, attachments []domain.Attachment) error {
	logging.Debug("Converting %d attachments", len(attachments))

	attachmentIDs, err := c.processAttachmentsGracefully(ctx, attachments)
	if err != nil {
		// Log warning but don't fail the entire result
		logging.Warn("Warning: Some attachments failed to process: %v", err)
		// Continue with available attachments
	}

	if len(attachmentIDs) > 0 {
		logging.Debug("Processed attachments, got %d IDs: %v", len(attachmentIDs), attachmentIDs)
		apiResult.SetAttachments(attachmentIDs)
	} else {
		logging.Debug("No attachments could be processed, continuing without attachments")
	}

	return nil // Don't fail the entire result due to attachment issues
}

// attachmentToUpload represents an attachment that needs to be uploaded
type attachmentToUpload struct {
	attachment domain.Attachment
	file       *os.File
	isTempFile bool // true if this is a temporary file that needs cleanup
}

// processAttachmentsGracefully processes attachments, skipping problematic ones instead of failing
// Groups attachments into batches respecting API limits:
// - Up to 32 MB per file
// - Up to 128 MB per single request
// - Up to 20 files per single request
func (c *V2Converter) processAttachmentsGracefully(ctx context.Context, attachments []domain.Attachment) ([]string, error) {
	logging.Debug("Processing %d attachments gracefully", len(attachments))
	logging.Debug("Uploader available: %v, project code: %s", c.uploader != nil, c.projectCode)

	var attachmentIDs []string
	var errors []string

	// Separate attachments into those that need upload and those with existing IDs

	var toUpload []attachmentToUpload
	var existingIDs []string

	for _, attachment := range attachments {
		// If attachment has existing ID and no file path/content, use it directly
		if attachment.ID != "" && !attachment.HasFilePath() && len(attachment.Content) == 0 {
			existingIDs = append(existingIDs, attachment.ID)
			logging.Debug("Using existing attachment ID: %s (no upload)", attachment.ID)
			continue
		}

		// If no uploader available, try to use existing ID or skip
		if c.uploader == nil || c.projectCode == "" {
			if attachment.ID != "" {
				existingIDs = append(existingIDs, attachment.ID)
				logging.Debug("Using existing attachment ID: %s (no uploader available)", attachment.ID)
			} else {
				logging.Warn("Warning: Attachment '%s' has no ID and no uploader available, skipping", attachment.FileName)
				errors = append(errors, fmt.Sprintf("no ID and no uploader for %s", attachment.FileName))
			}
			continue
		}

		// Prepare file for upload
		var file *os.File
		var isTempFile bool
		var err error

		if attachment.HasFilePath() {
			// Open file from path
			file, err = os.Open(attachment.GetFilePath())
			if err != nil {
				logging.Warn("Warning: Failed to open attachment file %s: %v, skipping", attachment.GetFilePath(), err)
				errors = append(errors, fmt.Sprintf("file open failed for %s: %v", attachment.GetFilePath(), err))
				continue
			}
			isTempFile = false
		} else if len(attachment.Content) > 0 {
			// Create temporary file for content
			logging.Debug("Creating temporary file for content attachment: %s", attachment.FileName)

			// Extract file extension to preserve it
			ext := filepath.Ext(attachment.FileName)
			baseName := strings.TrimSuffix(attachment.FileName, ext)

			// Create temp file with preserved extension
			tmpFile, err := os.CreateTemp("", baseName+"_*"+ext)
			if err != nil {
				logging.Warn("Warning: Failed to create temporary file for content attachment %s: %v, skipping", attachment.FileName, err)
				errors = append(errors, fmt.Sprintf("temp file creation failed for %s: %v", attachment.FileName, err))
				continue
			}

			// Write content to temporary file
			if _, err := tmpFile.Write(attachment.Content); err != nil {
				logging.Warn("Warning: Failed to write content to temporary file for %s: %v, skipping", attachment.FileName, err)
				os.Remove(tmpFile.Name()) // Clean up
				tmpFile.Close()
				errors = append(errors, fmt.Sprintf("content write failed for %s: %v", attachment.FileName, err))
				continue
			}

			// Reset file pointer to beginning
			if _, err := tmpFile.Seek(0, 0); err != nil {
				logging.Warn("Warning: Failed to seek temporary file for %s: %v, skipping", attachment.FileName, err)
				os.Remove(tmpFile.Name()) // Clean up
				tmpFile.Close()
				errors = append(errors, fmt.Sprintf("file seek failed for %s: %v", attachment.FileName, err))
				continue
			}

			file = tmpFile
			isTempFile = true
		} else {
			// No file path, no content, but has ID - use it
			if attachment.ID != "" {
				existingIDs = append(existingIDs, attachment.ID)
				logging.Debug("Using existing attachment ID: %s", attachment.ID)
			} else {
				logging.Warn("Warning: Attachment '%s' has no file path, content, or ID, skipping", attachment.FileName)
				errors = append(errors, fmt.Sprintf("no file path, content, or ID for %s", attachment.FileName))
			}
			continue
		}

		// Check file size (32 MB limit per file)
		fileInfo, err := file.Stat()
		if err != nil {
			logging.Warn("Warning: Failed to get file info for %s: %v, skipping", attachment.FileName, err)
			if isTempFile {
				os.Remove(file.Name())
			}
			file.Close()
			errors = append(errors, fmt.Sprintf("file stat failed for %s: %v", attachment.FileName, err))
			continue
		}

		const maxFileSize = 32 * 1024 * 1024 // 32 MB
		if fileInfo.Size() > maxFileSize {
			logging.Warn("Warning: File %s exceeds 32 MB limit (%d bytes), skipping", attachment.FileName, fileInfo.Size())
			if isTempFile {
				os.Remove(file.Name())
			}
			file.Close()
			errors = append(errors, fmt.Sprintf("file %s exceeds 32 MB limit", attachment.FileName))
			continue
		}

		toUpload = append(toUpload, attachmentToUpload{
			attachment: attachment,
			file:       file,
			isTempFile: isTempFile,
		})
	}

	// Add existing IDs to result
	attachmentIDs = append(attachmentIDs, existingIDs...)

	// Upload files in batches
	if len(toUpload) > 0 {
		uploadedIDs, uploadErrors := c.uploadAttachmentsInBatches(ctx, toUpload)
		attachmentIDs = append(attachmentIDs, uploadedIDs...)
		errors = append(errors, uploadErrors...)
	}

	// Log summary
	if len(errors) > 0 {
		logging.Warn("Warning: %d attachments had errors and were skipped: %v", len(errors), errors)
	}

	logging.Debug("Final attachment IDs: %v (processed: %d, skipped: %d)", attachmentIDs, len(attachmentIDs), len(errors))

	// Return error only if no attachments could be processed at all
	if len(attachmentIDs) == 0 && len(errors) > 0 {
		return nil, fmt.Errorf("all attachments failed to process: %v", errors)
	}

	return attachmentIDs, nil
}

// uploadAttachmentsInBatches uploads attachments in batches respecting API limits
// Returns uploaded IDs and errors
func (c *V2Converter) uploadAttachmentsInBatches(ctx context.Context, toUpload []attachmentToUpload) ([]string, []string) {
	const (
		maxFileSize      = 32 * 1024 * 1024  // 32 MB per file
		maxRequestSize   = 128 * 1024 * 1024 // 128 MB per request
		maxFilesPerBatch = 20                // 20 files per request
	)

	var uploadedIDs []string
	var errors []string

	// Group files into batches
	var currentBatch []attachmentToUpload
	var currentBatchSize int64

	uploadBatch := func(batch []attachmentToUpload) {
		if len(batch) == 0 {
			return
		}

		files := make([]*os.File, 0, len(batch))
		for _, item := range batch {
			files = append(files, item.file)
		}

		logging.Debug("Uploading batch of %d files (total size: %d bytes)", len(files), currentBatchSize)

		hashes, err := c.uploader.UploadAttachment(ctx, c.projectCode, files)
		if err != nil {
			logging.Warn("Warning: Failed to upload batch of %d files: %v", len(files), err)
			for _, item := range batch {
				errors = append(errors, fmt.Sprintf("batch upload failed for %s: %v", item.attachment.FileName, err))
				if item.isTempFile {
					os.Remove(item.file.Name())
				}
				item.file.Close()
			}
			return
		}

		// Match hashes to files (API returns hashes in the same order as files were sent)
		if len(hashes) != len(batch) {
			logging.Warn("Warning: Expected %d hashes from API, got %d", len(batch), len(hashes))
			// Use available hashes, mark missing ones as errors
			for i, item := range batch {
				if i < len(hashes) {
					uploadedIDs = append(uploadedIDs, hashes[i])
					logging.Debug("Successfully uploaded attachment, got hash: %s", hashes[i])
				} else {
					errors = append(errors, fmt.Sprintf("no hash returned for %s", item.attachment.FileName))
				}
				if item.isTempFile {
					os.Remove(item.file.Name())
				}
				item.file.Close()
			}
		} else {
			uploadedIDs = append(uploadedIDs, hashes...)
			for i, hash := range hashes {
				logging.Debug("Successfully uploaded attachment %s, got hash: %s", batch[i].attachment.FileName, hash)
				if batch[i].isTempFile {
					os.Remove(batch[i].file.Name())
				}
				batch[i].file.Close()
			}
		}
	}

	// Process files and create batches
	for _, item := range toUpload {
		fileInfo, err := item.file.Stat()
		if err != nil {
			logging.Warn("Warning: Failed to get file info for %s: %v, skipping", item.attachment.FileName, err)
			if item.isTempFile {
				os.Remove(item.file.Name())
			}
			item.file.Close()
			errors = append(errors, fmt.Sprintf("file stat failed for %s: %v", item.attachment.FileName, err))
			continue
		}

		fileSize := fileInfo.Size()

		// Check if adding this file would exceed limits
		wouldExceedSize := currentBatchSize+fileSize > maxRequestSize
		wouldExceedCount := len(currentBatch) >= maxFilesPerBatch

		// If current batch is full or would exceed limits, upload it first
		if wouldExceedSize || wouldExceedCount {
			if len(currentBatch) > 0 {
				uploadBatch(currentBatch)
				currentBatch = nil
				currentBatchSize = 0
			}

			// If single file exceeds request size limit, skip it (already checked file size limit)
			if fileSize > maxRequestSize {
				logging.Warn("Warning: File %s exceeds 128 MB request limit (%d bytes), skipping", item.attachment.FileName, fileSize)
				if item.isTempFile {
					os.Remove(item.file.Name())
				}
				item.file.Close()
				errors = append(errors, fmt.Sprintf("file %s exceeds 128 MB request limit", item.attachment.FileName))
				continue
			}
		}

		// Add file to current batch
		currentBatch = append(currentBatch, item)
		currentBatchSize += fileSize
	}

	// Upload remaining batch
	if len(currentBatch) > 0 {
		uploadBatch(currentBatch)
	}

	return uploadedIDs, errors
}

// convertSteps converts all steps from domain to API format
func (c *V2Converter) convertSteps(steps []domain.TestStep) ([]api_v2_client.ResultStep, error) {
	var apiSteps []api_v2_client.ResultStep
	var stepErrors []string

	for i, step := range steps {
		apiStep, err := c.ConvertTestStep(&step)
		if err != nil {
			// Log warning but don't fail the entire result
			logging.Warn("Warning: Failed to convert step %d '%s': %v, skipping step", i, step.Data.Action, err)
			stepErrors = append(stepErrors, fmt.Sprintf("step %d '%s': %v", i, step.Data.Action, err))
			continue // Skip this step and continue with others
		}
		apiSteps = append(apiSteps, *apiStep)
	}

	// Log summary of step conversion results
	if len(stepErrors) > 0 {
		logging.Warn("Warning: %d steps had conversion errors and were skipped: %v", len(stepErrors), stepErrors)
	}

	// Return error only if no steps could be converted at all
	if len(apiSteps) == 0 && len(stepErrors) > 0 {
		return nil, fmt.Errorf("all steps failed to convert: %v", stepErrors)
	}

	return apiSteps, nil
}

// setParams converts domain params to API format
func (c *V2Converter) setParams(apiResult *api_v2_client.ResultCreate, params map[string]string) {
	apiParams := make(map[string]string)
	for k, v := range params {
		apiParams[k] = v
	}
	apiResult.SetParams(apiParams)
}

// setParamGroups converts domain group params to API format (flattened)
func (c *V2Converter) setParamGroups(apiResult *api_v2_client.ResultCreate, groupParams map[string]string) {
	var paramGroups [][]string
	for key := range groupParams {
		paramGroups = append(paramGroups, []string{key})
	}
	apiResult.SetParamGroups(paramGroups)
}

// setTags converts domain tags to API format (comma-separated string in fields)
func (c *V2Converter) setTags(apiResult *api_v2_client.ResultCreate, tags []string) {
	// Deduplicate tags
	seen := make(map[string]struct{})
	unique := make([]string, 0, len(tags))
	for _, tag := range tags {
		if _, ok := seen[tag]; !ok {
			seen[tag] = struct{}{}
			unique = append(unique, tag)
		}
	}

	tagsStr := strings.Join(unique, ",")

	// Get existing fields or create new ones
	fields := apiResult.GetFields()
	fieldsPtr := &fields
	fieldsPtr.SetTags(tagsStr)
	apiResult.SetFields(*fieldsPtr)
}

// setRelations converts domain relations to API format
func (c *V2Converter) setRelations(apiResult *api_v2_client.ResultCreate, relations *domain.Relation) error {
	if relations.Suite == nil {
		return nil
	}

	apiRelations := api_v2_client.NewResultRelations()

	var suiteItems []api_v2_client.RelationSuiteItem
	for _, suiteData := range relations.Suite.Data {
		item := api_v2_client.NewRelationSuiteItem(suiteData.Title)
		if suiteData.PublicID != nil {
			item.SetPublicId(*suiteData.PublicID)
		}
		suiteItems = append(suiteItems, *item)
	}

	suite := api_v2_client.NewRelationSuite(suiteItems)
	apiRelations.SetSuite(*suite)
	apiResult.SetRelations(*apiRelations)

	return nil
}
