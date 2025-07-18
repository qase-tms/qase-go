package clients

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
	api_v1_client "github.com/qase-tms/qase-go/qase-api-client"
)

// V1Client implements Client interface using Qase API v1
type V1Client struct {
	client *api_v1_client.APIClient
	config ClientConfig
}

// NewV1Client creates a new API v1 client
func NewV1Client(config ClientConfig) (*V1Client, error) {
	cfg := api_v1_client.NewConfiguration()

	if config.BaseURL != "" {
		// Parse the base URL and set servers
		cfg.Servers = api_v1_client.ServerConfigurations{
			{
				URL:         config.BaseURL,
				Description: "Custom API server",
			},
		}
	}

	client := api_v1_client.NewAPIClient(cfg)

	return &V1Client{
		client: client,
		config: config,
	}, nil
}

// SendResult is not supported in API v1 client
func (c *V1Client) SendResult(ctx context.Context, projectCode string, runID int64, result *domain.TestResult) error {
	return fmt.Errorf("SendResult is not supported in API v1 client. Please use V2Client instead. You can create it with clients.NewV2Client()")
}

// SendResults is not supported in API v1 client
func (c *V1Client) SendResults(ctx context.Context, projectCode string, runID int64, results []*domain.TestResult) error {
	return fmt.Errorf("SendResults is not supported in API v1 client. Please use V2Client instead. You can create it with clients.NewV2Client()")
}

// CreateRun creates a new test run using API v1
func (c *V1Client) CreateRun(ctx context.Context, projectCode string, title, description string) (*RunInfo, error) {
	if c.config.Debug {
		log.Printf("Creating test run: project=%s, title=%s", projectCode, title)
	}

	// Create run request
	runCreate := api_v1_client.NewRunCreate(title)
	runCreate.SetDescription(description)

	// Set API token in context
	authCtx := context.WithValue(ctx, api_v1_client.ContextAPIKeys, map[string]api_v1_client.APIKey{
		"TokenAuth": {
			Key: c.config.APIToken,
		},
	})

	// Create run via API
	response, _, err := c.client.RunsAPI.CreateRun(authCtx, projectCode).RunCreate(*runCreate).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create test run: %w", err)
	}

	runInfo := &RunInfo{
		ID:          response.Result.GetId(),
		Title:       "", // Title not available in API v1 response
		Description: "", // Description not available in API v1 response
		URL:         fmt.Sprintf("https://app.qase.io/run/%s/dashboard/%d", projectCode, response.Result.GetId()),
	}

	if c.config.Debug {
		log.Printf("Successfully created test run: id=%d, url=%s", runInfo.ID, runInfo.URL)
	}

	return runInfo, nil
}

// CompleteRun marks a test run as completed using API v1
func (c *V1Client) CompleteRun(ctx context.Context, projectCode string, runID int64) error {
	if c.config.Debug {
		log.Printf("Completing test run: project=%s, run=%d", projectCode, runID)
	}

	// Set API token in context
	authCtx := context.WithValue(ctx, api_v1_client.ContextAPIKeys, map[string]api_v1_client.APIKey{
		"TokenAuth": {
			Key: c.config.APIToken,
		},
	})

	// Complete run via API
	_, _, err := c.client.RunsAPI.CompleteRun(authCtx, projectCode, int32(runID)).Execute()
	if err != nil {
		return fmt.Errorf("failed to complete test run: %w", err)
	}

	if c.config.Debug {
		log.Printf("Successfully completed test run: run=%d", runID)
	}

	return nil
}

// UploadAttachment uploads attachments to Qase API v1
func (c *V1Client) UploadAttachment(ctx context.Context, projectCode string, file []*os.File) (string, error) {
	const op = "client.clientv1.uploadattachment"

	if len(file) == 0 {
		return "", fmt.Errorf("no files provided")
	}

	if c.config.Debug {
		log.Printf("[%s] uploading attachment: projectCode=%s, file=%s", op, projectCode, file[0].Name())
	}

	// Set API token in context
	authCtx := context.WithValue(ctx, api_v1_client.ContextAPIKeys, map[string]api_v1_client.APIKey{
		"TokenAuth": {
			Key: c.config.APIToken,
		},
	})

	resp, r, err := c.client.AttachmentsAPI.
		UploadAttachment(authCtx, projectCode).
		File(file).
		Execute()

	if err != nil {
		return "", NewQaseApiError(err.Error(), extractBody(r))
	}

	if resp.Result == nil || len(resp.Result) == 0 {
		return "", fmt.Errorf("no attachment hash returned from API")
	}

	return *resp.Result[0].Hash, nil
}

// NewQaseApiError creates a new Qase API error
func NewQaseApiError(message string, body []byte) error {
	if len(body) > 0 {
		return fmt.Errorf("qase API error: %s, response body: %s", message, string(body))
	}
	return fmt.Errorf("qase API error: %s", message)
}

// extractBody extracts the response body for error reporting
func extractBody(resp *http.Response) []byte {
	if resp == nil || resp.Body == nil {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return body
}
