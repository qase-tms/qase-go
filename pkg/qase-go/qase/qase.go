package qase

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
	"github.com/qase-tms/qase-go/pkg/qase-go/logging"
	"github.com/qase-tms/qase-go/pkg/qase-go/reporters"
)

var (
	reporter *reporters.CoreReporter
	once     sync.Once
)

func init() {
	once.Do(func() {
		// Load configuration from root directory
		cfg, err := config.LoadWithParentSearch()
		if err != nil {
			// Fallback to unsafe loading if validation fails
			logging.Warn("Warning: Configuration validation failed, using unsafe loading: %v", err)
			cfg = config.LoadUnsafeWithParentSearch()
		}

		// Initialize logging system
		loggingConfig := &logging.Config{Debug: cfg.Debug}
		if err := logging.InitFromConfig(loggingConfig); err != nil {
			logging.Warn("Warning: Failed to initialize logging system: %v", err)
		} else {
			logging.Info("Logging system initialized successfully")
		}

		// Log configuration as JSON
		cfgJSON, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			logging.Error("Failed to marshal config to JSON: %v", err)
		} else {
			logging.Info("Qase configuration loaded:\n%s", string(cfgJSON))
		}

		r, err := reporters.NewCoreReporter(cfg)
		if err != nil {
			logging.Error("Failed to create core reporter: %v", err)
			return
		}

		reporter = r
		logging.Info("Qase reporter initialized successfully")
	})
}

// Test is the main test function
func Test(t *testing.T, meta TestMetadata, fn func()) {
	// Skip test if ignore is set
	if meta.Ignore {
		t.Skip("Test is marked as ignored")
		return
	}

	// Skip test if reporter is not available
	if reporter == nil {
		t.Log("Qase reporter not available, skipping result reporting")
		fn()
		return
	}

	// Create test result
	result := domain.NewTestResult(meta.Title)
	result.Execution.Status = domain.StatusPassed
	defer func() {
		// Add result to reporter
		if err := reporter.AddResult(result); err != nil {
			t.Logf("Failed to add result to reporter: %v", err)
		}
	}()

	// Run the test function
	fn()
}

// TestWithSteps is a test function that supports test steps
func TestWithSteps(t *testing.T, meta TestMetadata, fn func(*TestStepBuilder)) {
	// Skip test if ignore is set
	if meta.Ignore {
		t.Skip("Test is marked as ignored")
		return
	}

	// Skip test if reporter is not available
	if reporter == nil {
		t.Log("Qase reporter not available, skipping result reporting")
		builder := &TestStepBuilder{}
		fn(builder)
		return
	}

	// Create test result
	result := domain.NewTestResult(meta.Title)
	result.Execution.Status = domain.StatusPassed

	// Create step builder
	builder := &TestStepBuilder{
		result: result,
	}
	defer func() {
		// Add result to reporter
		if err := reporter.AddResult(result); err != nil {
			t.Logf("Failed to add result to reporter: %v", err)
		}
	}()

	// Run the test function
	fn(builder)
}

// TestStepBuilder helps build test steps
type TestStepBuilder struct {
	result *domain.TestResult
}

// AddStep adds a step to the test result
func (b *TestStepBuilder) AddStep(action string, status domain.StepStatus) {
	step := domain.NewTestStep(action)
	step.Execution.Status = status
	b.result.AddStep(*step)
}

// AddStepWithData adds a step with additional data to the test result
func (b *TestStepBuilder) AddStepWithData(action string, status domain.StepStatus, data map[string]interface{}) {
	step := domain.NewTestStep(action)
	step.Execution.Status = status
	// Convert map to StepTextData if needed
	if data != nil {
		if text, ok := data["text"].(string); ok {
			step.Data.Data = &text
		}
	}
	b.result.AddStep(*step)
}

// GetReporter returns the current reporter instance
func GetReporter() *reporters.CoreReporter {
	return reporter
}
