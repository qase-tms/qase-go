package reporters

import (
	"context"
	"fmt"
	"sync"

	"github.com/qase-tms/qase-go/pkg/qase-go/clients"
	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
	"github.com/qase-tms/qase-go/pkg/qase-go/logging"
)

// Reporter defines the interface that all reporters must implement
type Reporter interface {
	// AddResult adds a test result to the reporter
	AddResult(result *domain.TestResult) error

	// CompleteTestRun finalizes the test run
	CompleteTestRun(ctx context.Context) error
}

// CoreReporter manages a single reporter with fallback support
type CoreReporter struct {
	config        *config.Config
	reporter      Reporter
	fallback      Reporter
	statusMapping domain.StatusMapping
	mutex         sync.RWMutex
}

// NewCoreReporter creates a new core reporter with the given configuration
func NewCoreReporter(cfg *config.Config) (*CoreReporter, error) {
	if cfg == nil {
		return nil, fmt.Errorf("configuration cannot be nil")
	}

	reporter := &CoreReporter{
		config: cfg,
	}

	// Initialize status mapping
	if err := reporter.initializeStatusMapping(); err != nil {
		return nil, fmt.Errorf("failed to initialize status mapping: %w", err)
	}

	// Initialize reporter based on configuration
	if err := reporter.initializeReporter(); err != nil {
		return nil, fmt.Errorf("failed to initialize reporter: %w", err)
	}

	return reporter, nil
}

// initializeStatusMapping initializes the status mapping from configuration
func (cr *CoreReporter) initializeStatusMapping() error {
	statusMappingConfig := cr.config.GetStatusMapping()
	
	if len(statusMappingConfig) == 0 {
		// No status mapping configured
		cr.statusMapping = make(domain.StatusMapping)
		return nil
	}
	
	// Create status mapping from configuration
	mapping, err := domain.NewStatusMapping(statusMappingConfig)
	if err != nil {
		return fmt.Errorf("invalid status mapping configuration: %w", err)
	}
	
	cr.statusMapping = mapping
	
	// Log status mapping if debug is enabled
	if cr.config.Debug {
		logging.Info("Status mapping initialized: %s", mapping.String())
	}
	
	return nil
}

// initializeReporter sets up the main reporter and fallback based on configuration
func (cr *CoreReporter) initializeReporter() error {
	// Initialize fallback if configured
	if err := cr.initializeFallback(); err != nil {
		return fmt.Errorf("failed to initialize fallback: %w", err)
	}

	// Initialize main reporter based on mode
	switch cr.config.Mode {
	case "report":
		cr.reporter = NewFileReporter(FileReporterConfig{
			Config: cr.config,
		})
	case "testops":
		// Check if run ID is provided
		if cr.config.TestOps.Run.ID == nil {
			return fmt.Errorf("test run ID is required for TestOps mode - set QASE_TESTOPS_RUN_ID environment variable or id in run config")
		}

		// Try to create TestOps reporter
		client, err := cr.createTestOpsClient()
		if err != nil {
			// If TestOps client creation fails and we have fallback, use it
			if cr.fallback != nil {
				logging.Warn("Warning: TestOps client creation failed, using fallback: %v", err)
				cr.reporter = cr.fallback
				cr.fallback = nil // Clear fallback since we're using it as main reporter
				return nil
			}
			return fmt.Errorf("failed to create TestOps client: %w", err)
		}

		cr.reporter = NewTestOpsReporterWithConfig(client, *cr.config.TestOps.Run.ID, cr.config)
	case "off":
		return fmt.Errorf("reporting is disabled (mode: off)")
	default:
		return fmt.Errorf("unknown mode: %s", cr.config.Mode)
	}

	return nil
}

// initializeFallback sets up the fallback reporter based on configuration
func (cr *CoreReporter) initializeFallback() error {
	switch cr.config.Fallback {
	case "report":
		cr.fallback = NewFileReporter(FileReporterConfig{
			Config: cr.config,
		})
	case "testops":
		// Try to create TestOps client for fallback
		if cr.config.TestOps.Run.ID == nil {
			return fmt.Errorf("test run ID is required for TestOps fallback mode")
		}
		client, err := cr.createTestOpsClient()
		if err != nil {
			return fmt.Errorf("failed to create TestOps client for fallback: %w", err)
		}
		cr.fallback = NewTestOpsReporterWithConfig(client, *cr.config.TestOps.Run.ID, cr.config)
	case "off":
		// No fallback configured
		cr.fallback = nil
	default:
		return fmt.Errorf("unknown fallback mode: %s", cr.config.Fallback)
	}

	return nil
}

// TestOpsClientAdapter adapts the UnifiedClient to TestOpsClient interface
type TestOpsClientAdapter struct {
	client *clients.UnifiedClient
}

// UploadResults uploads test results to the specified run
func (a *TestOpsClientAdapter) UploadResults(ctx context.Context, runID int64, results []*domain.TestResult) error {
	return a.client.UploadResults(ctx, runID, results)
}

// createTestOpsClient creates a TestOps client based on configuration
func (cr *CoreReporter) createTestOpsClient() (TestOpsClient, error) {
	// Use the existing UnifiedClient from the clients package
	client, err := clients.NewUnifiedClient(cr.config)
	if err != nil {
		return nil, fmt.Errorf("failed to create unified client: %w", err)
	}

	// Create a TestOpsClientAdapter that wraps the UnifiedClient
	return &TestOpsClientAdapter{client: client}, nil
}

// AddResult adds a test result to the main reporter
func (cr *CoreReporter) AddResult(result *domain.TestResult) error {
	if result == nil {
		return fmt.Errorf("result cannot be nil")
	}

	if result.Title == "" {
		return fmt.Errorf("result title cannot be empty")
	}

	cr.mutex.Lock()
	defer cr.mutex.Unlock()

	if cr.reporter == nil {
		return fmt.Errorf("no reporter configured")
	}

	// Apply status mapping before adding to reporter
	originalStatus := result.Execution.Status
	if cr.statusMapping.ApplyMappingToResult(result) {
		// Status was mapped, log the change if debug is enabled
		if cr.config.Debug {
			logging.Info("Status mapped for test '%s': %s -> %s", 
				result.Title, originalStatus, result.Execution.Status)
		}
	}

	// Add to main reporter
	if err := cr.reporter.AddResult(result); err != nil {
		logging.Warn("Warning: Failed to add result: %v", err)
		// Try fallback if available
		if cr.fallback != nil {
			logging.Info("Trying fallback reporter")
			if fallbackErr := cr.fallback.AddResult(result); fallbackErr != nil {
				return fmt.Errorf("both main reporter and fallback failed: %w, fallback: %v", err, fallbackErr)
			}
			cr.reporter = cr.fallback
			cr.fallback = nil
		} else {
			return fmt.Errorf("failed to add result: %w", err)
		}
	}

	return nil
}

// CompleteTestRun completes the test run with the main reporter
func (cr *CoreReporter) CompleteTestRun(ctx context.Context) error {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()

	if cr.reporter == nil {
		return fmt.Errorf("no reporter configured")
	}

	// Complete test run with main reporter
	if err := cr.reporter.CompleteTestRun(ctx); err != nil {
		logging.Warn("Warning: Failed to complete test run: %v", err)
		// Try fallback if available
		if cr.fallback != nil {
			logging.Info("Trying fallback reporter")
			if fallbackErr := cr.fallback.CompleteTestRun(ctx); fallbackErr != nil {
				return fmt.Errorf("both main reporter and fallback failed: %w, fallback: %v", err, fallbackErr)
			}
			cr.reporter = cr.fallback
			cr.fallback = nil
		} else {
			return fmt.Errorf("failed to complete test run: %w", err)
		}
	}

	return nil
}

// GetResults returns all collected test results from the active reporter
func (cr *CoreReporter) GetResults() []*domain.TestResult {
	cr.mutex.RLock()
	defer cr.mutex.RUnlock()

	// Note: This is a placeholder. In a real implementation,
	// the reporter interface would need a GetResults method
	// For now, we return an empty slice since we don't store results
	return []*domain.TestResult{}
}

// GetConfig returns the current configuration
func (cr *CoreReporter) GetConfig() *config.Config {
	return cr.config
}

// GetReporterCount returns the number of active reporters
func (cr *CoreReporter) GetReporterCount() int {
	cr.mutex.RLock()
	defer cr.mutex.RUnlock()

	count := 0
	if cr.reporter != nil {
		count++
	}
	if cr.fallback != nil {
		count++
	}
	return count
}

// IsTestOpsMode returns true if the reporter is configured for TestOps mode
func (cr *CoreReporter) IsTestOpsMode() bool {
	return cr.config.Mode == "testops"
}

// IsReportMode returns true if the reporter is configured for report mode
func (cr *CoreReporter) IsReportMode() bool {
	return cr.config.Mode == "report"
}

// IsOffMode returns true if the reporter is configured for off mode
func (cr *CoreReporter) IsOffMode() bool {
	return cr.config.Mode == "off"
}

// GetCurrentMode returns the current active mode (may differ from config mode due to fallback)
func (cr *CoreReporter) GetCurrentMode() string {
	// Return the configured mode - the actual reporter type doesn't change the mode
	// Fallback is just a backup, not a mode change
	return cr.config.Mode
}

// CreateSimpleResult creates a simple test result with minimal required fields
func (cr *CoreReporter) CreateSimpleResult(title string, status domain.TestResultStatus) *domain.TestResult {
	return CreateSimpleResult(title, status)
}

// CreateResultWithSteps creates a test result with steps
func (cr *CoreReporter) CreateResultWithSteps(title string, status domain.TestResultStatus, steps []domain.TestStep) *domain.TestResult {
	return CreateResultWithSteps(title, status, steps)
}

// CreateStep creates a simple test step
func (cr *CoreReporter) CreateStep(action string, status domain.StepStatus) domain.TestStep {
	return CreateStep(action, status)
}

// GetStatusMapping returns the current status mapping
func (cr *CoreReporter) GetStatusMapping() domain.StatusMapping {
	return cr.statusMapping
}
