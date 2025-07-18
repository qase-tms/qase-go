package reporters

import (
	"context"
	"fmt"
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/clients"
	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

// mockClient implements the TestOpsClient interface for testing
type mockClient struct {
	createRunCalled    bool
	createRunRunID     int64
	createRunError     error
	completeRunCalled  bool
	completeRunError   error
	uploadResultsCalled bool
	uploadResultsError error
	uploadedResults    []*domain.TestResult
	uploadedRunID      int64
}

func (m *mockClient) CreateRun(ctx context.Context) (int64, error) {
	m.createRunCalled = true
	return m.createRunRunID, m.createRunError
}

func (m *mockClient) CompleteRun(ctx context.Context, runID int64) error {
	m.completeRunCalled = true
	return m.completeRunError
}

func (m *mockClient) UploadResults(ctx context.Context, runID int64, results []*domain.TestResult) error {
	m.uploadResultsCalled = true
	m.uploadedRunID = runID
	m.uploadedResults = results
	return m.uploadResultsError
}

func TestNewTestOpsReporter(t *testing.T) {
	cfg := &config.Config{
		TestOps: config.TestOpsConfig{
			API: config.APIConfig{
				Token: "test-token",
				Host:  "qase.io",
			},
			Project: "TEST",
		},
	}
	
	client, err := clients.NewUnifiedClient(cfg)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	
	reporter := NewTestOpsReporter(client)
	
	if reporter == nil {
		t.Fatal("reporter should not be nil")
	}
	if reporter.client != client {
		t.Error("client should be stored correctly")
	}
	if reporter.results == nil {
		t.Error("results slice should be initialized")
	}
	if len(reporter.results) != 0 {
		t.Error("results slice should be empty initially")
	}
	if reporter.runID != nil {
		t.Error("runID should be nil initially")
	}
}

func TestTestOpsReporter_StartTestRun(t *testing.T) {
	tests := []struct {
		name          string
		createRunID   int64
		createRunErr  error
		expectError   bool
	}{
		{
			name:        "successful run creation",
			createRunID: 123,
			createRunErr: nil,
			expectError: false,
		},
		{
			name:        "failed run creation",
			createRunID: 0,
			createRunErr: fmt.Errorf("API error"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockClient{
				createRunRunID: tt.createRunID,
				createRunError: tt.createRunErr,
			}
			
			reporter := NewTestOpsReporter(mock)
			ctx := context.Background()
			
			err := reporter.StartTestRun(ctx)
			
			if tt.expectError && err == nil {
				t.Error("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			
			if !mock.createRunCalled {
				t.Error("CreateRun should have been called")
			}
			
			if !tt.expectError && reporter.runID != nil && *reporter.runID != tt.createRunID {
				t.Errorf("expected runID %d, got %d", tt.createRunID, *reporter.runID)
			}
		})
	}
}

func TestTestOpsReporter_AddResult(t *testing.T) {
	tests := []struct {
		name        string
		result      *domain.TestResult
		expectError bool
	}{
		{
			name:        "valid result",
			result:      domain.NewTestResult("Test 1"),
			expectError: false,
		},
		{
			name:        "nil result",
			result:      nil,
			expectError: true,
		},
		{
			name:        "empty title",
			result:      domain.NewTestResult(""),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockClient{createRunRunID: 123}
			reporter := NewTestOpsReporter(mock)
			
			// Start test run first
			ctx := context.Background()
			err := reporter.StartTestRun(ctx)
			if err != nil {
				t.Fatalf("failed to start test run: %v", err)
			}
			
			initialCount := len(reporter.results)
			err = reporter.AddResult(tt.result)
			
			if tt.expectError && err == nil {
				t.Error("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			
			if !tt.expectError {
				if len(reporter.results) != initialCount+1 {
					t.Errorf("expected %d results, got %d", initialCount+1, len(reporter.results))
				}
				
				// Check that run ID was set
				if tt.result.RunID == nil || *tt.result.RunID != 123 {
					t.Error("result should have run ID set")
				}
			}
		})
	}
}

func TestTestOpsReporter_AddResult_WithoutStartingRun(t *testing.T) {
	mock := &mockClient{}
	reporter := NewTestOpsReporter(mock)
	
	result := domain.NewTestResult("Test 1")
	err := reporter.AddResult(result)
	
	if err != nil {
		t.Errorf("AddResult should not fail without started run: %v", err)
	}
	
	// Run ID should not be set
	if result.RunID != nil {
		t.Error("result should not have run ID set when no run started")
	}
}

func TestTestOpsReporter_CompleteTestRun(t *testing.T) {
	tests := []struct {
		name           string
		hasStartedRun  bool
		uploadError    error
		completeError  error
		expectError    bool
	}{
		{
			name:          "successful completion with results",
			hasStartedRun: true,
			uploadError:   nil,
			completeError: nil,
			expectError:   false,
		},
		{
			name:          "upload error",
			hasStartedRun: true,
			uploadError:   fmt.Errorf("upload failed"),
			completeError: nil,
			expectError:   true,
		},
		{
			name:          "complete error",
			hasStartedRun: true,
			uploadError:   nil,
			completeError: fmt.Errorf("complete failed"),
			expectError:   true,
		},
		{
			name:          "no run started",
			hasStartedRun: false,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockClient{
				createRunRunID:     123,
				uploadResultsError: tt.uploadError,
				completeRunError:   tt.completeError,
			}
			
			reporter := NewTestOpsReporter(mock)
			ctx := context.Background()
			
			if tt.hasStartedRun {
				err := reporter.StartTestRun(ctx)
				if err != nil {
					t.Fatalf("failed to start test run: %v", err)
				}
				
				// Add a test result
				result := domain.NewTestResult("Test 1")
				err = reporter.AddResult(result)
				if err != nil {
					t.Fatalf("failed to add result: %v", err)
				}
			}
			
			err := reporter.CompleteTestRun(ctx)
			
			if tt.expectError && err == nil {
				t.Error("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			
			if tt.hasStartedRun {
				if !mock.uploadResultsCalled {
					t.Error("UploadResults should have been called")
				}
				if tt.uploadError == nil && !mock.completeRunCalled {
					t.Error("CompleteRun should have been called")
				}
				
				// Verify uploaded data
				if mock.uploadedRunID != 123 {
					t.Errorf("expected uploaded run ID 123, got %d", mock.uploadedRunID)
				}
				if len(mock.uploadedResults) != 1 {
					t.Errorf("expected 1 uploaded result, got %d", len(mock.uploadedResults))
				}
			}
		})
	}
}

func TestTestOpsReporter_CompleteTestRun_EmptyResults(t *testing.T) {
	mock := &mockClient{createRunRunID: 123}
	reporter := NewTestOpsReporter(mock)
	ctx := context.Background()
	
	// Start test run without adding results
	err := reporter.StartTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to start test run: %v", err)
	}
	
	err = reporter.CompleteTestRun(ctx)
	if err != nil {
		t.Errorf("CompleteTestRun should not fail with empty results: %v", err)
	}
	
	// Note: UploadResults may not be called if the UnifiedClient optimizes empty result sets
	// The important thing is that CompleteRun is called
	if !mock.completeRunCalled {
		t.Error("CompleteRun should have been called")
	}
	
	// If UploadResults was called, it should have been called with empty results
	if mock.uploadResultsCalled && len(mock.uploadedResults) != 0 {
		t.Errorf("if UploadResults called, expected 0 uploaded results, got %d", len(mock.uploadedResults))
	}
}

func TestTestOpsReporter_WorkflowIntegration(t *testing.T) {
	// Test complete workflow with real UnifiedClient (will fail API calls but test structure)
	cfg := &config.Config{
		TestOps: config.TestOpsConfig{
			API: config.APIConfig{
				Token: "test-token",
				Host:  "qase.io",
			},
			Project: "TEST",
			Batch: config.BatchConfig{
				Size: 25,
			},
		},
	}
	
	client, err := clients.NewUnifiedClient(cfg)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	
	reporter := NewTestOpsReporter(client)
	ctx := context.Background()
	
	// This will fail due to no real API, but we test the workflow
	err = reporter.StartTestRun(ctx)
	if err == nil {
		t.Log("StartTestRun succeeded (unexpected but possible)")
	} else {
		t.Logf("StartTestRun failed as expected: %v", err)
	}
	
	// Add results regardless
	result1 := domain.NewTestResult("Test 1")
	result1.Execution.Status = domain.StatusPassed
	err = reporter.AddResult(result1)
	if err != nil {
		t.Errorf("AddResult should not fail: %v", err)
	}
	
	result2 := domain.NewTestResult("Test 2")
	result2.Execution.Status = domain.StatusFailed
	err = reporter.AddResult(result2)
	if err != nil {
		t.Errorf("AddResult should not fail: %v", err)
	}
	
	if len(reporter.results) != 2 {
		t.Errorf("expected 2 results, got %d", len(reporter.results))
	}
}