package clients

import (
	"context"
	"fmt"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
	"github.com/qase-tms/qase-go/pkg/qase-go/logging"
)

// UnifiedClient combines v1 and v2 clients for optimal API usage
type UnifiedClient struct {
	v1Client    *V1Client // For run management
	v2Client    *V2Client // For result uploading
	config      *config.Config
	projectCode string
}

// NewUnifiedClient creates a new unified client that uses v1 for runs and v2 for results
func NewUnifiedClient(cfg *config.Config) (*UnifiedClient, error) {
	// Create client config from main config
	clientConfig := ClientConfig{
		BaseURL:  buildAPIBaseURL(cfg.TestOps.API.Host),
		APIToken: cfg.TestOps.API.Token,
		Debug:    cfg.Debug,
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
	v2Client.SetConverter(NewV2ConverterWithUploader(v1Client, cfg.TestOps.Project))

	return &UnifiedClient{
		v1Client:    v1Client,
		v2Client:    v2Client,
		config:      cfg,
		projectCode: cfg.TestOps.Project,
	}, nil
}

// UploadResults uploads test results using v2 API with batching
func (c *UnifiedClient) UploadResults(ctx context.Context, runID int64, results []*domain.TestResult) error {
	if len(results) == 0 {
		return nil
	}

	// Get batch size from config
	batchSize := c.config.TestOps.Batch.Size
	if batchSize <= 0 {
		batchSize = 50 // Default batch size
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

	// Log summary
	if failedUploads > 0 {
		logging.Info("Upload summary: %d successful, %d failed", successfulUploads, failedUploads)
		if successfulUploads == 0 {
			// If no results were uploaded at all, return error
			return fmt.Errorf("failed to upload any results: %w", lastError)
		}
		// If some results were uploaded successfully, log warning but don't fail
		logging.Warn("Warning: Some results failed to upload, but test run will continue")
	}

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
