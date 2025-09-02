package reporters

import (
	"context"
	"fmt"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

// TestOpsClient defines the interface for test operations client
type TestOpsClient interface {
	// UploadResults uploads test results to the specified run
	UploadResults(ctx context.Context, runID int64, results []*domain.TestResult) error
}

// TestOpsReporter handles sending test results to Qase TestOps
type TestOpsReporter struct {
	client  TestOpsClient
	results []*domain.TestResult
	runID   int64
}


// NewTestOpsReporter creates a new TestOps reporter with the given TestOps client and run ID
func NewTestOpsReporter(client TestOpsClient, runID int64) *TestOpsReporter {
	return &TestOpsReporter{
		client:  client,
		results: make([]*domain.TestResult, 0),
		runID:   runID,
	}
}




// AddResult adds a test result to the internal collection
func (r *TestOpsReporter) AddResult(result *domain.TestResult) error {
	if result == nil {
		return fmt.Errorf("result cannot be nil")
	}
	
	if result.Title == "" {
		return fmt.Errorf("result title cannot be empty")
	}
	
	// Set run ID
	result.SetRunID(r.runID)
	
	r.results = append(r.results, result)
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
