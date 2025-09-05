package reporters

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

// TestOpsClient defines the interface for test operations client
type TestOpsClient interface {
	// UploadResults uploads test results to the specified run
	UploadResults(ctx context.Context, runID int64, results []*domain.TestResult) error
}

// TestOpsReporter handles sending test results to Qase TestOps
type TestOpsReporter struct {
	client       TestOpsClient
	results      []*domain.TestResult
	runID        int64
	config       *config.Config
	statusFilter []string
}

// NewTestOpsReporter creates a new TestOps reporter with the given TestOps client and run ID
func NewTestOpsReporter(client TestOpsClient, runID int64) *TestOpsReporter {
	return &TestOpsReporter{
		client:       client,
		results:      make([]*domain.TestResult, 0),
		runID:        runID,
		config:       nil,
		statusFilter: nil,
	}
}

// NewTestOpsReporterWithConfig creates a new TestOps reporter with configuration for status filtering
func NewTestOpsReporterWithConfig(client TestOpsClient, runID int64, cfg *config.Config) *TestOpsReporter {
	statusFilter := cfg.TestOps.StatusFilter
	return &TestOpsReporter{
		client:       client,
		results:      make([]*domain.TestResult, 0),
		runID:        runID,
		config:       cfg,
		statusFilter: statusFilter,
	}
}

// shouldIncludeResult checks if a result should be included based on status filter
func (r *TestOpsReporter) shouldIncludeResult(result *domain.TestResult) bool {
	// If no status filter is configured, include all results
	if len(r.statusFilter) == 0 {
		return true
	}

	// Check if the result's status should be excluded (is in the filter list)
	resultStatus := string(result.Execution.Status)
	for _, excludedStatus := range r.statusFilter {
		if strings.EqualFold(resultStatus, excludedStatus) {
			return false // Exclude this result
		}
	}

	return true // Include this result
}

// AddResult adds a test result to the internal collection
func (r *TestOpsReporter) AddResult(result *domain.TestResult) error {
	if result == nil {
		return fmt.Errorf("result cannot be nil")
	}

	if result.Title == "" {
		return fmt.Errorf("result title cannot be empty")
	}

	// Check if result should be included based on status filter
	if !r.shouldIncludeResult(result) {
		// Log that the result was filtered out (if debug is enabled)
		if r.config != nil && r.config.Debug {
			fmt.Printf("Filtered out result '%s' with status '%s' (excluded by statusFilter: %v)\n",
				result.Title, result.Execution.Status, r.statusFilter)
		}
		return nil // Silently skip filtered results
	}

	// Set run ID
	result.SetRunID(r.runID)

	r.results = append(r.results, result)
	return nil
}

// CompleteTestRun sends all collected results to TestOps
func (r *TestOpsReporter) CompleteTestRun(ctx context.Context) error {
	if len(r.results) == 0 {
		return nil // No results to send
	}

	// Send all results to TestOps
	if err := r.client.UploadResults(ctx, r.runID, r.results); err != nil {
		return fmt.Errorf("failed to upload results: %w", err)
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
