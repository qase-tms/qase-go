package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
	"github.com/qase-tms/qase-go/pkg/qase-go/logging"
	api_v2_client "github.com/qase-tms/qase-go/qase-api-v2-client"
)

// V2Client implements Client interface using Qase API v2
type V2Client struct {
	client    *api_v2_client.APIClient
	config    ClientConfig
	converter *V2Converter
}

// NewV2Client creates a new API v2 client
func NewV2Client(config ClientConfig) (*V2Client, error) {
	cfg := api_v2_client.NewConfiguration()

	if config.BaseURL != "" {
		// Override default base URL if provided
		cfg.Servers = api_v2_client.ServerConfigurations{
			{
				URL:         config.BaseURL,
				Description: "Custom API server",
			},
		}
		if config.Debug {
			// Debug logging removed for clean code
		}
	} else {
		// Use default configuration from the generated client
		if config.Debug {
			// Debug logging removed for clean code
		}
	}

	client := api_v2_client.NewAPIClient(cfg)

	return &V2Client{
		client:    client,
		config:    config,
		converter: NewV2Converter(),
	}, nil
}

// logResultPretty logs a result in a readable JSON format
func (c *V2Client) logResultPretty(prefix string, result *api_v2_client.ResultCreate) {
	if !c.config.Debug {
		return
	}

	execution := result.GetExecution()

	// Create a simplified structure for logging
	logData := map[string]interface{}{
		"title": result.GetTitle(),
	}

	// Add execution object
	executionData := map[string]interface{}{
		"status": execution.GetStatus(),
	}
	if execution.StartTime.IsSet() {
		executionData["start_time"] = execution.GetStartTime()
	}
	if execution.EndTime.IsSet() {
		executionData["end_time"] = execution.GetEndTime()
	}
	if execution.Duration.IsSet() {
		executionData["duration_ms"] = execution.GetDuration()
	}
	if execution.Stacktrace.IsSet() {
		executionData["stacktrace"] = execution.GetStacktrace()
	}
	if execution.Thread.IsSet() {
		executionData["thread"] = execution.GetThread()
	}
	logData["execution"] = executionData

	// Convert to JSON for pretty printing
	if jsonData, err := json.MarshalIndent(logData, "", "  "); err == nil {
		logging.Info("%s: %s", prefix, string(jsonData))
	} else {
		logging.Info("%s: %+v", prefix, logData)
	}
}

// logBatchRequestPretty logs a batch request in a readable format
func (c *V2Client) logBatchRequestPretty(batchRequest *api_v2_client.CreateResultsRequestV2) {
	if !c.config.Debug {
		return
	}

	results := batchRequest.GetResults()

	// Create structure for the entire batch request
	batchLogData := map[string]interface{}{
		"total_results": len(results),
		"results":       make([]map[string]interface{}, 0, len(results)),
	}

	for i, result := range results {
		execution := result.GetExecution()

		// Create detailed structure for each result
		resultData := map[string]interface{}{
			"index": i,
			"title": result.GetTitle(),
		}

		// Add all available fields
		if result.HasId() {
			resultData["id"] = result.GetId()
		}
		if result.HasSignature() {
			resultData["signature"] = result.GetSignature()
		}
		if result.TestopsId.IsSet() {
			resultData["testops_id"] = result.GetTestopsId()
		}
		if len(result.GetTestopsIds()) > 0 {
			resultData["testops_ids"] = result.GetTestopsIds()
		}

		// Add execution object
		executionData := map[string]interface{}{
			"status": execution.GetStatus(),
		}
		if execution.StartTime.IsSet() {
			executionData["start_time"] = execution.GetStartTime()
		}
		if execution.EndTime.IsSet() {
			executionData["end_time"] = execution.GetEndTime()
		}
		if execution.Duration.IsSet() {
			executionData["duration_ms"] = execution.GetDuration()
		}
		if execution.Stacktrace.IsSet() {
			executionData["stacktrace"] = execution.GetStacktrace()
		}
		if execution.Thread.IsSet() {
			executionData["thread"] = execution.GetThread()
		}
		resultData["execution"] = executionData

		// Steps
		if len(result.GetSteps()) > 0 {
			stepsData := make([]map[string]interface{}, 0, len(result.GetSteps()))
			for j, step := range result.GetSteps() {
				stepData := map[string]interface{}{
					"step_index": j,
				}

				// Step data
				if step.Data != nil {
					stepData["data"] = step.Data
				}

				// Step execution
				if step.Execution != nil {
					exec := step.Execution
					stepExec := map[string]interface{}{}

					if exec.StartTime.IsSet() {
						stepExec["start_time"] = exec.GetStartTime()
					}
					if exec.EndTime.IsSet() {
						stepExec["end_time"] = exec.GetEndTime()
					}
					if exec.Duration.IsSet() {
						stepExec["duration_ms"] = exec.GetDuration()
					}
					stepExec["status"] = exec.GetStatus()
					if exec.HasComment() {
						stepExec["comment"] = exec.GetComment()
					}
					if len(exec.GetAttachments()) > 0 {
						stepExec["attachments"] = exec.GetAttachments()
					}

					stepData["execution"] = stepExec
				}

				// Nested steps
				if len(step.GetSteps()) > 0 {
					stepData["nested_steps"] = step.GetSteps()
				}

				stepsData = append(stepsData, stepData)
			}
			resultData["steps"] = stepsData
		}

		// Steps type
		if result.StepsType.IsSet() {
			resultData["steps_type"] = result.GetStepsType()
		}

		// Parameters
		if result.GetParams() != nil {
			resultData["params"] = result.GetParams()
		}

		// Parameter groups
		if len(result.GetParamGroups()) > 0 {
			resultData["param_groups"] = result.GetParamGroups()
		}

		// Relations
		if result.Relations.IsSet() {
			resultData["relations"] = result.GetRelations()
		}

		// Message
		if result.Message.IsSet() {
			resultData["message"] = result.GetMessage()
		}

		// Defect
		if result.HasDefect() {
			resultData["defect"] = result.GetDefect()
		}

		// Fields
		if result.HasFields() {
			resultData["fields"] = result.GetFields()
		}

		// Attachments
		if len(result.GetAttachments()) > 0 {
			resultData["attachments"] = result.GetAttachments()
		}

		batchLogData["results"] = append(batchLogData["results"].([]map[string]interface{}), resultData)
	}

	// Output the entire batch request in one JSON
	if jsonData, err := json.MarshalIndent(batchLogData, "", "  "); err == nil {
		logging.Info("Batch request:\n%s", string(jsonData))
	} else {
		logging.Info("Batch request: %+v", batchLogData)
	}
}

// SetConverter sets a custom converter for the V2Client
func (c *V2Client) SetConverter(converter *V2Converter) {
	c.converter = converter
}

// SendResult sends a single test result to Qase using API v2
func (c *V2Client) SendResult(ctx context.Context, projectCode string, runID int64, result *domain.TestResult) error {
	logging.Info("Sending result to Qase API v2: project=%s, run=%d, result=%s", projectCode, runID, result.Title)
	logging.Info("Domain result: %s", result.String())

	// Convert domain model to API v2 model
	apiResult, err := c.converter.ConvertTestResult(result)
	if err != nil {
		// Log warning but don't fail - try to send with partial data
		logging.Warn("Warning: Result '%s' had conversion issues: %v, attempting to send with partial data", result.Title, err)

		// Try to create a minimal result with basic information
		if apiResult == nil {
			// Create a minimal result with just title and status
			execution := api_v2_client.NewResultExecution(string(result.Execution.Status))
			minimalResult := api_v2_client.NewResultCreate(result.Title, *execution)
			if result.ID != "" {
				minimalResult.SetId(result.ID)
			}
			if result.Signature != "" {
				minimalResult.SetSignature(result.Signature)
			}
			apiResult = minimalResult
		}

		// If we still don't have a result, return error
		if apiResult == nil {
			return fmt.Errorf("failed to convert domain model and could not create minimal result: %w", err)
		}
	}

	if c.config.Debug {
		c.logResultPretty("Converted API result", apiResult)
	}

	// Set API token in context
	authCtx := context.WithValue(ctx, api_v2_client.ContextAPIKeys, map[string]api_v2_client.APIKey{
		"TokenAuth": {
			Key: c.config.APIToken,
		},
	})

	// Send result to API
	httpResp, err := c.client.ResultsAPI.CreateResultV2(authCtx, projectCode, runID).ResultCreate(*apiResult).Execute()
	if err != nil {
		if c.config.Debug {
			logging.Error("API request failed with error: %v", err)
			if httpResp != nil {
				logging.Info("HTTP response status: %s", httpResp.Status)
				logging.Info("HTTP response headers: %v", httpResp.Header)
				if body, readErr := io.ReadAll(httpResp.Body); readErr == nil {
					logging.Info("HTTP response body: %s", string(body))
				}
			}
		}
		return fmt.Errorf("failed to send result to Qase API v2: %w", err)
	}

	logging.Info("Successfully sent result: HTTP status=%s", httpResp.Status)

	return nil
}

// SendResults sends multiple test results to Qase using API v2 batch endpoint
func (c *V2Client) SendResults(ctx context.Context, projectCode string, runID int64, results []*domain.TestResult) error {
	logging.Info("Sending batch results to Qase API v2: count=%d, project=%s, run=%d", len(results), projectCode, runID)
	for i, result := range results {
		logging.Info("Result %d: %s", i, result.String())
	}

	// Convert all domain models to API v2 models, attempting to convert all results
	var apiResults []api_v2_client.ResultCreate
	var conversionWarnings []string

	for _, result := range results {
		apiResult, err := c.converter.ConvertTestResult(result)
		if err != nil {
			// Log warning but don't skip the result - try to send with partial data
			logging.Warn("Warning: Result '%s' had conversion issues: %v, attempting to send with partial data", result.Title, err)
			conversionWarnings = append(conversionWarnings, fmt.Sprintf("'%s': %v", result.Title, err))

			// Try to create a minimal result with basic information
			if apiResult == nil {
				// Create a minimal result with just title and status
				execution := api_v2_client.NewResultExecution(string(result.Execution.Status))
				minimalResult := api_v2_client.NewResultCreate(result.Title, *execution)
				if result.ID != "" {
					minimalResult.SetId(result.ID)
				}
				if result.Signature != "" {
					minimalResult.SetSignature(result.Signature)
				}
				apiResult = minimalResult
			}
		}

		if apiResult != nil {
			apiResults = append(apiResults, *apiResult)
		}
	}

	// Log summary of conversion results
	if len(conversionWarnings) > 0 {
		logging.Warn("Warning: %d results had conversion issues: %v", len(conversionWarnings), conversionWarnings)
	}

	if len(apiResults) == 0 {
		logging.Warn("Warning: No results could be converted at all, skipping batch upload")
		return fmt.Errorf("no results could be converted for batch upload")
	}

	logging.Info("Converted API results count: %d (warnings: %d)", len(apiResults), len(conversionWarnings))
	for i, result := range apiResults {
		c.logResultPretty(fmt.Sprintf("Result %d", i), &result)
	}

	// Create batch request
	batchRequest := api_v2_client.NewCreateResultsRequestV2()
	batchRequest.SetResults(apiResults)

	if c.config.Debug {
		c.logBatchRequestPretty(batchRequest)
	}

	// Set API token in context
	authCtx := context.WithValue(ctx, api_v2_client.ContextAPIKeys, map[string]api_v2_client.APIKey{
		"TokenAuth": {
			Key: c.config.APIToken,
		},
	})

	// Send batch results to API
	httpResp, err := c.client.ResultsAPI.CreateResultsV2(authCtx, projectCode, runID).CreateResultsRequestV2(*batchRequest).Execute()
	if err != nil {
		if c.config.Debug {
			logging.Error("API request failed with error: %v", err)
			if httpResp != nil {
				logging.Info("HTTP response status: %s", httpResp.Status)
				logging.Info("HTTP response headers: %v", httpResp.Header)
				if body, readErr := io.ReadAll(httpResp.Body); readErr == nil {
					logging.Info("HTTP response body: %s", string(body))
				}
			}
		}
		return fmt.Errorf("failed to send batch results to Qase API v2: %w", err)
	}

	logging.Info("Successfully sent batch results: HTTP status=%s", httpResp.Status)

	return nil
}

// UploadAttachment is not supported in API v2 client, use V1Client instead
func (c *V2Client) UploadAttachment(ctx context.Context, projectCode string, file []*os.File) (string, error) {
	return "", fmt.Errorf("UploadAttachment is not supported in API v2 client. Please use V1Client instead")
}
