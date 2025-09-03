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
	uploadResultsCalled bool
	uploadResultsError  error
	uploadedResults     []*domain.TestResult
	uploadedRunID       int64
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

	runID := int64(123)
	reporter := NewTestOpsReporter(client, runID)

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
	if reporter.runID != runID {
		t.Errorf("expected runID %d, got %d", runID, reporter.runID)
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
			mock := &mockClient{}
			runID := int64(123)
			reporter := NewTestOpsReporter(mock, runID)

			initialCount := len(reporter.results)
			err := reporter.AddResult(tt.result)

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
				if tt.result.RunID == nil || *tt.result.RunID != runID {
					t.Error("result should have run ID set")
				}
			}
		})
	}
}

func TestTestOpsReporter_CompleteTestRun(t *testing.T) {
	tests := []struct {
		name        string
		uploadError error
		expectError bool
		hasResults  bool
	}{
		{
			name:        "successful completion with results",
			uploadError: nil,
			expectError: false,
			hasResults:  true,
		},
		{
			name:        "upload error",
			uploadError: fmt.Errorf("upload failed"),
			expectError: true,
			hasResults:  true,
		},
		{
			name:        "no results to upload",
			uploadError: nil,
			expectError: false,
			hasResults:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockClient{
				uploadResultsError: tt.uploadError,
			}

			runID := int64(123)
			reporter := NewTestOpsReporter(mock, runID)
			ctx := context.Background()

			if tt.hasResults {
				// Add a test result
				result := domain.NewTestResult("Test 1")
				err := reporter.AddResult(result)
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

			if tt.hasResults {
				if !mock.uploadResultsCalled {
					t.Error("UploadResults should have been called")
				}

				// Verify uploaded data
				if mock.uploadedRunID != runID {
					t.Errorf("expected uploaded run ID %d, got %d", runID, mock.uploadedRunID)
				}
				if len(mock.uploadedResults) != 1 {
					t.Errorf("expected 1 uploaded result, got %d", len(mock.uploadedResults))
				}
			} else {
				// If no results, UploadResults should not be called
				if mock.uploadResultsCalled {
					t.Error("UploadResults should not have been called when no results")
				}
			}
		})
	}
}

func TestTestOpsReporter_CompleteTestRun_EmptyResults(t *testing.T) {
	mock := &mockClient{}
	runID := int64(123)
	reporter := NewTestOpsReporter(mock, runID)
	ctx := context.Background()

	// Complete test run without adding results
	err := reporter.CompleteTestRun(ctx)
	if err != nil {
		t.Errorf("CompleteTestRun should not fail with empty results: %v", err)
	}

	// UploadResults should not be called if there are no results
	if mock.uploadResultsCalled {
		t.Error("UploadResults should not have been called with empty results")
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

	runID := int64(123)
	reporter := NewTestOpsReporter(client, runID)
	ctx := context.Background()

	// Add results
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

	// This will fail due to no real API, but we test the workflow
	err = reporter.CompleteTestRun(ctx)
	if err == nil {
		t.Log("CompleteTestRun succeeded (unexpected but possible)")
	} else {
		t.Logf("CompleteTestRun failed as expected: %v", err)
	}
}

func TestCreateSimpleResult(t *testing.T) {
	title := "Test Simple Result"
	status := domain.StatusPassed

	result := CreateSimpleResult(title, status)

	if result == nil {
		t.Fatal("result should not be nil")
	}

	if result.Title != title {
		t.Errorf("expected title %q, got %q", title, result.Title)
	}

	if result.Execution.Status != status {
		t.Errorf("expected status %v, got %v", status, result.Execution.Status)
	}

	if result.Execution.StartTime == nil {
		t.Error("StartTime should be set")
	}

	if result.Execution.EndTime == nil {
		t.Error("EndTime should be set")
	}

	if result.Execution.StartTime != nil && result.Execution.EndTime != nil {
		if *result.Execution.StartTime != *result.Execution.EndTime {
			t.Error("StartTime and EndTime should be equal for simple result")
		}
	}
}

func TestCreateResultWithSteps(t *testing.T) {
	title := "Test With Steps"
	status := domain.StatusFailed

	steps := []domain.TestStep{
		CreateStep("Step 1", domain.StepStatusPassed),
		CreateStep("Step 2", domain.StepStatusFailed),
	}

	result := CreateResultWithSteps(title, status, steps)

	if result == nil {
		t.Fatal("result should not be nil")
	}

	if result.Title != title {
		t.Errorf("expected title %q, got %q", title, result.Title)
	}

	if result.Execution.Status != status {
		t.Errorf("expected status %v, got %v", status, result.Execution.Status)
	}

	if len(result.Steps) != len(steps) {
		t.Errorf("expected %d steps, got %d", len(steps), len(result.Steps))
	}

	for i, step := range result.Steps {
		if step.Data.Action != steps[i].Data.Action {
			t.Errorf("step %d: expected action %q, got %q", i, steps[i].Data.Action, step.Data.Action)
		}
		if step.Execution.Status != steps[i].Execution.Status {
			t.Errorf("step %d: expected status %v, got %v", i, steps[i].Execution.Status, step.Execution.Status)
		}
	}
}

func TestCreateStep(t *testing.T) {
	action := "Click button"
	status := domain.StepStatusPassed

	step := CreateStep(action, status)

	if step.Data.Action != action {
		t.Errorf("expected action %q, got %q", action, step.Data.Action)
	}

	if step.Execution.Status != status {
		t.Errorf("expected status %v, got %v", status, step.Execution.Status)
	}

	if step.Execution.StartTime == nil {
		t.Error("StartTime should be set")
	}

	if step.Execution.EndTime == nil {
		t.Error("EndTime should be set")
	}

	if step.Execution.StartTime != nil && step.Execution.EndTime != nil {
		if *step.Execution.StartTime != *step.Execution.EndTime {
			t.Error("StartTime and EndTime should be equal for simple step")
		}
	}
}
