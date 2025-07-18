package reporters

import (
	"context"
	"fmt"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

// TestOpsClient defines the interface for test operations client
type TestOpsClient interface {
	// CreateRun creates a new test run and returns its ID
	CreateRun(ctx context.Context) (int64, error)
	
	// CompleteRun completes the test run by ID
	CompleteRun(ctx context.Context, runID int64) error
	
	// UploadResults uploads test results to the specified run
	UploadResults(ctx context.Context, runID int64, results []*domain.TestResult) error
}

// TestOpsReporter handles sending test results to Qase TestOps with a simple 3-method API
type TestOpsReporter struct {
	client  TestOpsClient
	results []*domain.TestResult
	runID   *int64
}


// NewTestOpsReporter creates a new TestOps reporter with the given TestOps client
func NewTestOpsReporter(client TestOpsClient) *TestOpsReporter {
	return &TestOpsReporter{
		client:  client,
		results: make([]*domain.TestResult, 0),
	}
}


// StartTestRun creates a new test run and stores the run ID
func (r *TestOpsReporter) StartTestRun(ctx context.Context) error {
	runID, err := r.client.CreateRun(ctx)
	if err != nil {
		return fmt.Errorf("failed to create test run: %w", err)
	}
	
	r.runID = &runID
	return nil
}

// AddResult adds a test result to the internal collection
func (r *TestOpsReporter) AddResult(result *domain.TestResult) error {
	if result == nil {
		return fmt.Errorf("result cannot be nil")
	}
	
	if result.Title == "" {
		return fmt.Errorf("result title cannot be empty")
	}
	
	// Set run ID if not already set
	if result.RunID == nil && r.runID != nil {
		result.SetRunID(*r.runID)
	}
	
	r.results = append(r.results, result)
	return nil
}

// CompleteTestRun sends all accumulated results and completes the test run
func (r *TestOpsReporter) CompleteTestRun(ctx context.Context) error {
	if r.runID == nil {
		return fmt.Errorf("no test run started - call StartTestRun first")
	}
	
	// Upload results using unified client (handles batching internally)
	if len(r.results) > 0 {
		err := r.client.UploadResults(ctx, *r.runID, r.results)
		if err != nil {
			return fmt.Errorf("failed to upload results: %w", err)
		}
	}
	
	// Complete the test run
	err := r.client.CompleteRun(ctx, *r.runID)
	if err != nil {
		return fmt.Errorf("failed to complete test run: %w", err)
	}
	
	return nil
}

// CreateSimpleResult creates a simple test result with minimal required fields
func CreateSimpleResult(title string, status domain.TestResultStatus) *domain.TestResult {
	result := domain.NewTestResult(title)
	result.Execution.Status = status
	
	now := time.Now().Unix()
	result.Execution.StartTime = &now
	result.Execution.EndTime = &now
	
	return result
}

// CreateResultWithSteps creates a test result with steps
func CreateResultWithSteps(title string, status domain.TestResultStatus, steps []domain.TestStep) *domain.TestResult {
	result := CreateSimpleResult(title, status)
	
	for _, step := range steps {
		result.AddStep(step)
	}
	
	return result
}

// CreateStep creates a simple test step
func CreateStep(action string, status domain.StepStatus) domain.TestStep {
	step := *domain.NewTestStep(action)
	step.Execution.Status = status
	
	now := time.Now().Unix()
	step.Execution.StartTime = &now
	step.Execution.EndTime = &now
	
	return step
}
