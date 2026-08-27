package clients

import (
	"context"
	"fmt"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
	"github.com/qase-tms/qase-go/pkg/qase-go/logging"
)

// resultSender sends test results to Qase. It is the part of the v2 client
// UnifiedClient depends on, extracted so that batching can be tested without
// a live API.
type resultSender interface {
	SendResult(ctx context.Context, projectCode string, runID int64, result *domain.TestResult) error
	SendResults(ctx context.Context, projectCode string, runID int64, results []*domain.TestResult) error
}

// UnifiedClient combines v1 and v2 clients for optimal API usage
type UnifiedClient struct {
	v1Client    *V1Client    // For run management
	v2Client    resultSender // For result uploading
	config      *config.Config
	projectCode string
}

// NewUnifiedClient creates a new unified client that uses v1 for runs and v2 for results
// This function creates HostData internally for backward compatibility
func NewUnifiedClient(cfg *config.Config) (*UnifiedClient, error) {
	hostData := GetHostInfo()
	return NewUnifiedClientWithHostData(cfg, hostData)
}

// NewUnifiedClientWithHostData creates a new unified client with provided HostData
// This allows passing HostData from CoreReporter to avoid duplicate creation
func NewUnifiedClientWithHostData(cfg *config.Config, hostData *HostData) (*UnifiedClient, error) {
	if hostData == nil {
		// Fallback: create host data if not provided
		hostData = GetHostInfo()
	}

	// Create client config from main config
	clientConfig := ClientConfig{
		BaseURL:  buildAPIBaseURL(cfg.TestOps.API.Host),
		APIToken: cfg.TestOps.API.Token,
		Debug:    cfg.Debug,
		HostData: hostData,
	}

	// Create v1 client for run management
	v1Client, err := NewV1Client(clientConfig)
	if err != nil {
		return nil, err
	}

	// Create v2 client for result uploading
	v2Client, err := NewV2Client(clientConfig)
	if err != nil {
		return nil, err
	}

	// Update v2 client converter to use v1 client for attachment uploads
	// Use adapter to convert V1Client to AttachmentUploader interface
	uploader := NewV1ClientAdapter(v1Client)
	v2Client.SetConverter(NewV2ConverterWithUploader(uploader, cfg.TestOps.Project))

	return &UnifiedClient{
		v1Client:    v1Client,
		v2Client:    v2Client,
		config:      cfg,
		projectCode: cfg.TestOps.Project,
	}, nil
}

// UploadResults uploads test results using v2 API with batching.
//
// Results are sent in chunks of at most config.MaxBatchSize: the bulk endpoint
// rejects larger requests with a non-retryable HTTP 413. The caller's slice is
// never modified, so results of a failed chunk stay available to it.
func (c *UnifiedClient) UploadResults(ctx context.Context, runID int64, results []*domain.TestResult) error {
	if len(results) == 0 {
		return nil
	}

	// Get batch size from config, clamped to the limit accepted by the API
	batchSize := c.config.GetBatchSize()
	if configured := c.config.TestOps.Batch.Size; configured > batchSize {
		logging.Warn("Configured batch size %d exceeds the API limit, using %d instead", configured, batchSize)
	}

	// Track successful and failed uploads
	var successfulUploads int
	var failedUploads int
	var lastError error

	// Send results in batches
	for i := 0; i < len(results); i += batchSize {
		end := i + batchSize
		if end > len(results) {
			end = len(results)
		}

		batch := results[i:end]

		if len(batch) == 1 {
			// Send single result
			err := c.v2Client.SendResult(ctx, c.projectCode, runID, batch[0])
			if err != nil {
				logging.Warn("Warning: Failed to send single result '%s': %v", batch[0].Title, err)
				failedUploads++
				lastError = err
			} else {
				successfulUploads++
			}
		} else {
			// Send batch of results
			err := c.v2Client.SendResults(ctx, c.projectCode, runID, batch)
			if err != nil {
				logging.Warn("Warning: Failed to send batch of %d results: %v", len(batch), err)
				failedUploads += len(batch)
				lastError = err
			} else {
				successfulUploads += len(batch)
			}
		}
	}

	// Report every lost result: a silent partial upload makes a run look
	// complete while results are missing
	if failedUploads > 0 {
		logging.Warn("Upload summary: %d of %d results uploaded, %d lost", successfulUploads, len(results), failedUploads)
		return fmt.Errorf("failed to upload %d of %d test results: %w", failedUploads, len(results), lastError)
	}

	logging.Debug("Upload summary: %d results uploaded", successfulUploads)

	return nil
}

// GetProjectCode returns the project code
func (c *UnifiedClient) GetProjectCode() string {
	return c.projectCode
}

// GetConfig returns the configuration
func (c *UnifiedClient) GetConfig() *config.Config {
	return c.config
}
