package qase

import (
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



// AddResult adds a test result
func (r *Reporter) AddResult(result *domain.TestResult) error {
	if r.coreReporter == nil {
		return fmt.Errorf("core reporter is not initialized")
	}
	return r.coreReporter.AddResult(result)
}



// GetCoreReporter returns the underlying core reporter
func (r *Reporter) GetCoreReporter() *reporters.CoreReporter {
	return r.coreReporter
}
