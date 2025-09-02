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

	// Use DisplayName as title, fallback to Title if DisplayName is empty
	title := meta.DisplayName
	if title == "" {
		title = meta.Title
	}
	if title == "" {
		title = t.Name() // Use test name as fallback
	}

	// Create test result
	result := domain.NewTestResult(title)
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

	// Use DisplayName as title, fallback to Title if DisplayName is empty
	title := meta.DisplayName
	if title == "" {
		title = meta.Title
	}
	if title == "" {
		title = t.Name() // Use test name as fallback
	}

	// Create test result
	result := domain.NewTestResult(title)
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

// AddMessage adds a message to the current test result
func AddMessage(message string) {
	if reporter == nil {
		return
	}
	// For now, we'll just log the message
	// In a more complex implementation, this could be stored in the current test result
	logging.Info("Test message: %s", message)
}

// Step executes a test step with metadata
func Step(t *testing.T, meta StepMetadata, fn func()) {
	// For now, just log the step and execute the function
	logging.Info("Executing step: %s", meta.Name)
	if meta.Description != "" {
		logging.Info("Step description: %s", meta.Description)
	}
	fn()
}

// AddAttachments adds file attachments to the current test result
func AddAttachments(filePaths ...string) {
	if reporter == nil {
		return
	}
	// For now, just log the attachments
	for _, path := range filePaths {
		logging.Info("Adding attachment: %s", path)
	}
}

// AttachContent adds content as an attachment to the current test result
func AttachContent(name, content, mimeType string) {
	if reporter == nil {
		return
	}
	// For now, just log the content attachment
	logging.Info("Adding content attachment: %s (type: %s)", name, mimeType)
}

// InitializeGlobal initializes qase globally (for backward compatibility)
// This function is kept for backward compatibility but doesn't do anything special
// since initialization now happens automatically in init()
func InitializeGlobal() error {
	// Initialization already happened in init(), just return success
	return nil
}

// TestMain is a placeholder function for backward compatibility
// In the simplified architecture, this is not needed
func TestMain(m *testing.M) {
	// Just run the tests
	_ = m.Run()
}

// GetReporter returns the current reporter instance
func GetReporter() *reporters.CoreReporter {
	return reporter
}
