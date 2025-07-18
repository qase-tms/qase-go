package clients

import (
	"context"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

// UnifiedClient combines v1 and v2 clients for optimal API usage
type UnifiedClient struct {
	v1Client    *V1Client  // For run management
	v2Client    *V2Client  // For result uploading
	config      *config.Config
	projectCode string
}

// NewUnifiedClient creates a new unified client that uses v1 for runs and v2 for results
func NewUnifiedClient(cfg *config.Config) (*UnifiedClient, error) {
	// Create client config from main config
	clientConfig := ClientConfig{
		BaseURL:  cfg.TestOps.API.Host,
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

// CreateRun creates a new test run using v1 API and returns the run ID
func (c *UnifiedClient) CreateRun(ctx context.Context) (int64, error) {
	// Get run title from config or use default
	title := c.config.TestOps.Run.Title
	if title == "" {
		title = "Automated Test Run"
	}
	
	// Get run description from config
	description := c.config.TestOps.Run.Description
	if description == "" && c.config.Environment != "" {
		description = "Test run in " + c.config.Environment + " environment"
	}
	
	// Create run using v1 client
	runInfo, err := c.v1Client.CreateRun(ctx, c.projectCode, title, description)
	if err != nil {
		return 0, err
	}
	
	return runInfo.ID, nil
}

// CompleteRun completes the test run using v1 API
func (c *UnifiedClient) CompleteRun(ctx context.Context, runID int64) error {
	return c.v1Client.CompleteRun(ctx, c.projectCode, runID)
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
				return err
			}
		} else {
			// Send batch of results
			err := c.v2Client.SendResults(ctx, c.projectCode, runID, batch)
			if err != nil {
				return err
			}
		}
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