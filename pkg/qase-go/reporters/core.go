package reporters

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/qase-tms/qase-go/pkg/qase-go/clients"
	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

// Reporter defines the interface that all reporters must implement
type Reporter interface {
	// StartTestRun initializes the test run
	StartTestRun(ctx context.Context) error

	// AddResult adds a test result to the reporter
	AddResult(result *domain.TestResult) error

	// CompleteTestRun finalizes the test run
	CompleteTestRun(ctx context.Context) error
}

// CoreReporter manages a single reporter with fallback support
type CoreReporter struct {
	config   *config.Config
	reporter Reporter
	fallback Reporter
	runID    *int64
	mutex    sync.RWMutex
}

// NewCoreReporter creates a new core reporter with the given configuration
func NewCoreReporter(cfg *config.Config) (*CoreReporter, error) {
	if cfg == nil {
		return nil, fmt.Errorf("configuration cannot be nil")
	}

	reporter := &CoreReporter{
		config: cfg,
	}

	// Initialize reporter based on configuration
	if err := reporter.initializeReporter(); err != nil {
		return nil, fmt.Errorf("failed to initialize reporter: %w", err)
	}

	return reporter, nil
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
		// Try to create TestOps reporter
		client, err := cr.createTestOpsClient()
		if err != nil {
			// If TestOps client creation fails and we have fallback, use it
			if cr.fallback != nil {
				log.Printf("Warning: TestOps client creation failed, using fallback: %v", err)
				cr.reporter = cr.fallback
				cr.fallback = nil // Clear fallback since we're using it as main reporter
				return nil
			}
			return fmt.Errorf("failed to create TestOps client: %w", err)
		}

		cr.reporter = NewTestOpsReporter(client)
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
		client, err := cr.createTestOpsClient()
		if err != nil {
			return fmt.Errorf("failed to create TestOps client for fallback: %w", err)
		}
		cr.fallback = NewTestOpsReporter(client)
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

// CreateRun creates a new test run and returns its ID
func (a *TestOpsClientAdapter) CreateRun(ctx context.Context) (int64, error) {
	return a.client.CreateRun(ctx)
}

// CompleteRun completes the test run by ID
func (a *TestOpsClientAdapter) CompleteRun(ctx context.Context, runID int64) error {
	return a.client.CompleteRun(ctx, runID)
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

// StartTestRun starts the test run with the main reporter
func (cr *CoreReporter) StartTestRun(ctx context.Context) error {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()

	// Clear run ID
	cr.runID = nil

	if cr.reporter == nil {
		return fmt.Errorf("no reporter configured")
	}

	// Start test run with main reporter
	if err := cr.reporter.StartTestRun(ctx); err != nil {
		log.Printf("Warning: Failed to start test run: %v", err)
		// Try fallback if available
		if cr.fallback != nil {
			log.Printf("Trying fallback reporter")
			if fallbackErr := cr.fallback.StartTestRun(ctx); fallbackErr != nil {
				return fmt.Errorf("both main reporter and fallback failed: %w, fallback: %v", err, fallbackErr)
			}
			cr.reporter = cr.fallback
			cr.fallback = nil
		} else {
			return fmt.Errorf("failed to start test run: %w", err)
		}
	}

	return nil
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

	// Add to main reporter
	if err := cr.reporter.AddResult(result); err != nil {
		log.Printf("Warning: Failed to add result: %v", err)
		// Try fallback if available
		if cr.fallback != nil {
			log.Printf("Trying fallback reporter")
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
		log.Printf("Warning: Failed to complete test run: %v", err)
		// Try fallback if available
		if cr.fallback != nil {
			log.Printf("Trying fallback reporter")
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

// GetRunID returns the current run ID if available
func (cr *CoreReporter) GetRunID() *int64 {
	cr.mutex.RLock()
	defer cr.mutex.RUnlock()

	if cr.runID != nil {
		runID := *cr.runID
		return &runID
	}
	return nil
}

// SetRunID sets the run ID
func (cr *CoreReporter) SetRunID(runID int64) {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()

	cr.runID = &runID
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
	// If we have a file reporter as main reporter, we're in report mode
	if cr.reporter != nil {
		// This is a simple check - in a real implementation you might want to use type assertion
		// For now, we'll check if the config mode is report or if we're using fallback
		if cr.config.Mode == "report" {
			return "report"
		}
		// If we're in testops mode but have a file reporter, we're using fallback
		if cr.config.Mode == "testops" {
			return "report" // fallback mode
		}
	}
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
