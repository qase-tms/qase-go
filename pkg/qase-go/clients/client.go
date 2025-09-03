package clients

import (
	"context"
	"fmt"
	"os"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

// Client represents a Qase API client interface
type Client interface {
	// SendResult sends a single test result to Qase
	SendResult(ctx context.Context, projectCode string, runID int64, result *domain.TestResult) error
	
	// SendResults sends multiple test results to Qase (batch operation)
	SendResults(ctx context.Context, projectCode string, runID int64, results []*domain.TestResult) error
	
	// UploadAttachment uploads attachments to Qase
	UploadAttachment(ctx context.Context, projectCode string, file []*os.File) (string, error)
}

// RunInfo contains information about a test run
type RunInfo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// ClientConfig contains configuration for API clients
type ClientConfig struct {
	// BaseURL is the Qase API base URL (optional, defaults will be used)
	BaseURL string
	
	// APIToken is the Qase API token
	APIToken string
	
	// Debug enables debug logging
	Debug bool
	
	// UseAPIv2 determines whether to use API v2 (default: true)
	UseAPIv2 bool
}

// NewClient creates a new Qase client based on configuration
func NewClient(config ClientConfig) (Client, error) {
	if config.APIToken == "" {
		return nil, fmt.Errorf("API token is required")
	}
	
	if config.UseAPIv2 {
		return NewV2Client(config)
	}
	
	return NewV1Client(config)
}

// NewClientFromConfig creates a new Qase client from config.Config
func NewClientFromConfig(cfg *config.Config, useAPIv2 bool) (Client, error) {
	clientConfig := ClientConfig{
		BaseURL:  buildAPIBaseURL(cfg.TestOps.API.Host),
		APIToken: cfg.TestOps.API.Token,
		Debug:    cfg.Debug,
		UseAPIv2: useAPIv2,
	}
	
	return NewClient(clientConfig)
}

// buildAPIBaseURL constructs the API base URL from host
func buildAPIBaseURL(host string) string {
	if host == "" || host == "qase.io" {
		return "" // Use default
	}
	return "https://api." + host
}
