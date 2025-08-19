package clients

import (
	"context"
	"fmt"
	"os"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
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

// SetConverter sets a custom converter for the V2Client
func (c *V2Client) SetConverter(converter *V2Converter) {
	c.converter = converter
}

// SendResult sends a single test result to Qase using API v2
func (c *V2Client) SendResult(ctx context.Context, projectCode string, runID int64, result *domain.TestResult) error {
	// Convert domain model to API v2 model
	apiResult, err := c.converter.ConvertTestResult(result)
	if err != nil {
		return fmt.Errorf("failed to convert domain model: %w", err)
	}

	// Set API token in context
	authCtx := context.WithValue(ctx, api_v2_client.ContextAPIKeys, map[string]api_v2_client.APIKey{
		"TokenAuth": {
			Key: c.config.APIToken,
		},
	})

	// Send result to API
	_, err = c.client.ResultsAPI.CreateResultV2(authCtx, projectCode, runID).ResultCreate(*apiResult).Execute()
	if err != nil {
		return fmt.Errorf("failed to send result to Qase API v2: %w", err)
	}

	return nil
}

// SendResults sends multiple test results to Qase using API v2 batch endpoint
func (c *V2Client) SendResults(ctx context.Context, projectCode string, runID int64, results []*domain.TestResult) error {
	// Convert all domain models to API v2 models
	var apiResults []api_v2_client.ResultCreate
	for _, result := range results {
		apiResult, err := c.converter.ConvertTestResult(result)
		if err != nil {
			return fmt.Errorf("failed to convert domain model for result '%s': %w", result.Title, err)
		}
		apiResults = append(apiResults, *apiResult)
	}

	// Create batch request
	batchRequest := api_v2_client.NewCreateResultsRequestV2()
	batchRequest.SetResults(apiResults)

	// Set API token in context
	authCtx := context.WithValue(ctx, api_v2_client.ContextAPIKeys, map[string]api_v2_client.APIKey{
		"TokenAuth": {
			Key: c.config.APIToken,
		},
	})

	// Send batch results to API
	_, err := c.client.ResultsAPI.CreateResultsV2(authCtx, projectCode, runID).CreateResultsRequestV2(*batchRequest).Execute()
	if err != nil {
		return fmt.Errorf("failed to send batch results to Qase API v2: %w", err)
	}

	return nil
}

// CreateRun creates a new test run (placeholder - not implemented in v2 client yet)
func (c *V2Client) CreateRun(ctx context.Context, projectCode string, title, description string) (*RunInfo, error) {
	// TODO: Implement when run creation API is available in v2 client
	return nil, fmt.Errorf("CreateRun not implemented for API v2 client")
}

// CompleteRun marks a test run as completed (placeholder - not implemented in v2 client yet)
func (c *V2Client) CompleteRun(ctx context.Context, projectCode string, runID int64) error {
	// TODO: Implement when run completion API is available in v2 client
	return fmt.Errorf("CompleteRun not implemented for API v2 client")
}

// UploadAttachment is not supported in API v2 client, use V1Client instead
func (c *V2Client) UploadAttachment(ctx context.Context, projectCode string, file []*os.File) (string, error) {
	return "", fmt.Errorf("UploadAttachment is not supported in API v2 client. Please use V1Client instead")
}
