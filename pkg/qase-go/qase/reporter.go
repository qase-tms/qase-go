package qase

import (
	"context"
	"fmt"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
	"github.com/qase-tms/qase-go/pkg/qase-go/reporters"
)

// Reporter wraps the core reporter for use in the framework
type Reporter struct {
	coreReporter *reporters.CoreReporter
}

// NewReporter creates a new reporter wrapper
func NewReporter(coreReporter *reporters.CoreReporter) *Reporter {
	return &Reporter{
		coreReporter: coreReporter,
	}
}

// StartTestRun starts the test run
func (r *Reporter) StartTestRun(ctx context.Context) error {
	if r.coreReporter == nil {
		return fmt.Errorf("core reporter is not initialized")
	}
	return r.coreReporter.StartTestRun(ctx)
}

// AddResult adds a test result
func (r *Reporter) AddResult(result *domain.TestResult) error {
	if r.coreReporter == nil {
		return fmt.Errorf("core reporter is not initialized")
	}
	return r.coreReporter.AddResult(result)
}

// CompleteTestRun completes the test run
func (r *Reporter) CompleteTestRun(ctx context.Context) error {
	if r.coreReporter == nil {
		return fmt.Errorf("core reporter is not initialized")
	}
	return r.coreReporter.CompleteTestRun(ctx)
}

// GetCoreReporter returns the underlying core reporter
func (r *Reporter) GetCoreReporter() *reporters.CoreReporter {
	return r.coreReporter
}
