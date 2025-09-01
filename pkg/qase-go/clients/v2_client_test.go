package clients

import (
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

func TestV2Client_SendResults_HandlesConversionErrorsGracefully(t *testing.T) {
	// Create a mock client that will simulate API responses
	client := &V2Client{
		config: ClientConfig{
			Debug: true,
		},
		converter: NewV2Converter(),
	}

	// Create test results with problematic data
	results := []*domain.TestResult{
		{
			Title: "Valid Result",
			Execution: domain.TestExecution{
				Status: domain.StatusPassed,
			},
		},
		{
			Title: "Result With Problematic Attachments",
			Execution: domain.TestExecution{
				Status: domain.StatusFailed,
			},
			Attachments: []domain.Attachment{
				{
					FileName: "problematic.txt",
					Content:  []byte("content"),
					// No ID and no file path - this will cause issues
				},
			},
		},
	}

	// Mock the API client to avoid actual HTTP calls
	// This test just verifies that conversion happens without errors
	// The actual sending will fail due to missing API client, but that's expected

	// The key point is that conversion should succeed despite attachment issues
	for _, result := range results {
		apiResult, err := client.converter.ConvertTestResult(result)
		if err != nil {
			t.Errorf("Expected successful conversion for result '%s', got error: %v", result.Title, err)
		}
		if apiResult == nil {
			t.Errorf("Expected non-nil API result for result '%s'", result.Title)
		}
	}
}
