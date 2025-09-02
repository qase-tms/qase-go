package qase

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
	"github.com/qase-tms/qase-go/pkg/qase-go/logging"
	"github.com/qase-tms/qase-go/pkg/qase-go/reporters"
)

var (
	once     sync.Once
	reporter *reporters.CoreReporter
	// Global variable to store current test result for error details
	currentTestResult *domain.TestResult
	currentTestMutex  sync.RWMutex
	// Stack to track nested steps
	currentStepStack []*domain.TestStep
	stepStackMutex   sync.RWMutex
	// Add this for automatic test run completion
	completeTestRunOnce sync.Once
)

func init() {
	once.Do(func() {
		// Check if global initialization has already been done
		// Use environment variable to track global initialization across packages
		if os.Getenv("QASE_GLOBAL_INITIALIZED") == "true" {
			// Global initialization was done, skip local initialization
			return
		}

		// Fallback to local initialization if global wasn't called
		cfg, err := config.Load()
		if err != nil {
			// Fallback to unsafe loading if validation fails
			logging.Warn("Warning: Configuration validation failed, using unsafe loading: %v", err)
			cfg = config.LoadUnsafe()
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
		if err := reporter.StartTestRun(context.Background()); err != nil {
			logging.Error("Failed to start test run: %v", err)
			return
		}

		logging.Info("Qase test run started successfully (local initialization)")
	})
}

// Test is the main test function
func Test(t *testing.T, meta TestMetadata, fn func()) {
	// Skip test if ignore is set
	if meta.Ignore {
		t.Skip("Test ignored by Qase metadata")
		return
	}

	result := domain.NewTestResult(meta.DisplayName)

	// Set title
	if meta.Title != "" {
		result.Title = meta.Title
	}

	// Set description
	if meta.Description != "" {
		result.Fields["description"] = meta.Description
	}

	// Set comment
	if meta.Comment != "" {
		result.Fields["comment"] = meta.Comment
	}

	// Set test case IDs
	if len(meta.IDs) > 0 {
		if len(meta.IDs) == 1 {
			result.SetTestopsIDSingle(meta.IDs[0])
		} else {
			result.SetTestopsIDMultiple(meta.IDs)
		}
	}

	// Set suite information
	if meta.Suite != "" {
		result.AddSuiteData(meta.Suite, nil)
	}

	// Add fields (matching JavaScript fields)
	for k, v := range meta.Fields {
		result.Fields[k] = v
	}

	// Add parameters (matching JavaScript parameters)
	for k, v := range meta.Parameters {
		result.Params[k] = v
	}

	// Add group parameters (matching JavaScript groupParams)
	for k, v := range meta.GroupParameters {
		result.GroupParams[k] = v
	}

	// Add system information
	result.Fields["go_version"] = runtime.Version()
	result.Fields["os"] = runtime.GOOS
	result.Fields["arch"] = runtime.GOARCH
	result.Fields["test_file"] = getTestFile()
	result.Fields["test_line"] = getTestLine()

	// Begin test execution
	now := time.Now().UnixMilli()
	result.Execution.StartTime = &now

	// Set current test result for error details capture
	setCurrentTestResult(result)

	defer func() {
		// Clear current test result
		clearCurrentTestResult()

		if r := recover(); r != nil {
			// Panic occurred - set status to invalid and capture stacktrace
			result.Execution.Status = domain.StatusInvalid
			msg := fmt.Sprintf("panic: %v", r)
			result.Message = &msg

			// Capture stacktrace for panic
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			stacktrace := string(buf[:n])
			result.Execution.Stacktrace = &stacktrace

		} else if t.Failed() {
			// Test failed due to assertion failure
			result.Execution.Status = domain.StatusFailed

			// Try to get more detailed error information from the test
			if result.Execution.ErrorDetails == nil {
				msg := "Test failed due to assertion failure"
				result.Message = &msg
			} else {
				// Use error details for more specific message
				if result.Execution.ErrorDetails.ErrorType != "" {
					msg := fmt.Sprintf("Test failed: %s", result.Execution.ErrorDetails.ErrorType)
					result.Message = &msg
				}
			}

			// Capture stacktrace for failed test
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			stacktrace := string(buf[:n])
			result.Execution.Stacktrace = &stacktrace

		} else {
			result.Execution.Status = domain.StatusPassed
		}

		// End test execution
		end := time.Now().UnixMilli()
		result.Execution.EndTime = &end
		duration := end - now
		result.Execution.Duration = &duration

		// Add result to reporter
		if getActiveReporter() != nil {
			_ = getActiveReporter().AddResult(result)
		}
	}()

	fn()
}

// Step creates a test step
func Step(t *testing.T, meta StepMetadata, fn func()) {
	step := domain.NewTestStep(meta.Name)

	if meta.Description != "" {
		step.Data.ExpectedResult = &meta.Description
	}

	if meta.ExpectedResult != "" {
		step.Data.ExpectedResult = &meta.ExpectedResult
	}

	if meta.Data != "" {
		step.Data.Data = &meta.Data
	}

	// Begin step execution
	now := time.Now().UnixMilli()
	step.Execution.StartTime = &now

	fmt.Printf("Step: %s - %s\n", meta.Name, meta.Description)

	// Push this step onto the stack
	pushStep(step)

	// Execute the step function
	func() {
		defer func() {
			if r := recover(); r != nil {
				// Panic occurred in step - set status to failed
				step.Execution.Status = domain.StepStatusFailed
			}
		}()
		fn()
	}()

	// End step execution and set status
	end := time.Now().UnixMilli()
	step.Execution.EndTime = &end
	duration := end - now
	step.Execution.Duration = &duration

	// Set step status based on test result
	// Note: t.Failed() might not be accurate during step execution
	// We'll set it to passed by default and let the test framework handle failures
	step.Execution.Status = domain.StepStatusPassed

	// Pop this step from the stack
	popStep()

	// Add step to the appropriate parent
	if currentResult := getCurrentTestResult(); currentResult != nil {
		if currentStep := getCurrentStep(); currentStep != nil {
			// This is a nested step, add it to the current step
			currentStep.Steps = append(currentStep.Steps, *step)
		} else {
			// This is a top-level step, add it to the test result
			currentResult.AddStep(*step)
		}
	}
}

// AddMessage adds a message to the current test
func AddMessage(msg string) {
	fmt.Println("Message:", msg)

	// Try to capture error details from the current test result
	if currentResult := getCurrentTestResult(); currentResult != nil {
		// Check if this looks like an assertion error message
		if isAssertionError(msg) {
			// Extract error type and set error details
			errorType := extractErrorType(msg)
			currentResult.SetErrorDetailsFromString(errorType, msg)

			// Try to capture file and line information
			if file, line := getCallerInfo(); file != "" && line > 0 {
				currentResult.SetErrorLocation(file, line)
			}
		}
	}
}

// AddAttachments adds attachments to the current test
func AddAttachments(path string) {
	// Get current test result
	result := getCurrentTestResult()
	if result == nil {
		logging.Warn("Warning: No current test result available for attachment %s", path)
		return
	}

	// Create attachment from file path
	attachment := domain.NewAttachment("", path, "", nil)
	attachment.SetFilePath(path)

	// Add to result attachments
	result.Attachments = append(result.Attachments, *attachment)

	// Log for debugging
	if getActiveReporter() != nil && getActiveReporter().GetConfig().Debug {
		logging.Info("Added attachment: %s to test: %s", path, result.Title)
	}
}

// AttachFile attaches a file with specific content type
func AttachFile(name, filePath, contentType string) {
	// Get current test result
	result := getCurrentTestResult()
	if result == nil {
		logging.Warn("Warning: No current test result available for attachment %s (%s)", name, filePath)
		return
	}

	// Create attachment from file path
	attachment := domain.NewAttachment("", name, contentType, nil)
	attachment.SetFilePath(filePath)

	// Add to result attachments
	result.Attachments = append(result.Attachments, *attachment)

	// Log for debugging
	if getActiveReporter() != nil && getActiveReporter().GetConfig().Debug {
		logging.Info("Added file attachment: %s (%s) - %s to test: %s", name, contentType, filePath, result.Title)
	}
}

// AttachContent attaches content with specific content type
func AttachContent(name, content, contentType string) {
	// Get current test result
	result := getCurrentTestResult()
	if result == nil {
		logging.Warn("Warning: No current test result available for content attachment %s", name)
		return
	}

	// Create attachment with content
	attachment := domain.NewAttachment("", name, contentType, []byte(content))

	// Add to result attachments
	result.Attachments = append(result.Attachments, *attachment)

	// Log for debugging
	if getActiveReporter() != nil && getActiveReporter().GetConfig().Debug {
		contentPreview := content
		if len(content) > 50 {
			contentPreview = content[:50] + "..."
		}
		logging.Info("Added content attachment: %s (%s) - %s to test: %s", name, contentType, contentPreview, result.Title)
	}
}

// CompleteTestRun completes the test run only once
// Note: Test run will be automatically completed after all tests finish
// This function can still be called manually if immediate completion is needed
func CompleteTestRun() error {
	var err error
	completeTestRunOnce.Do(func() {
		if getActiveReporter() != nil {
			logging.Info("Completing test run (will only happen once)")
			err = getActiveReporter().CompleteTestRun(context.Background())
			if err != nil {
				logging.Error("Error completing test run: %v", err)
			} else {
				logging.Info("Test run completed successfully")
			}
		}
	})
	return err
}

// TestMain should be called in the main test file for automatic test run completion
// Example usage in your test file:
//
//	func TestMain(m *testing.M) {
//	    qase.TestMain(m)
//	}
func TestMain(m *testing.M) {
	// Run tests
	exitCode := m.Run()

	// Complete test run automatically after all tests finish
	if reporter != nil {
		logging.Info("All tests completed, auto-completing test run")
		if err := CompleteTestRun(); err != nil {
			logging.Warn("Warning: Failed to auto-complete test run: %v", err)
		}
	}

	os.Exit(exitCode)
}

// Global initialization functions

// InitializeGlobal initializes qase globally with configuration that searches parent directories
// This should be called once at the beginning of your test suite, typically in a TestMain function
// or in the root package's init function
func InitializeGlobal() error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	if IsGlobalInitialized() {
		return nil // Already initialized
	}

	// Use parent search to find config in root directory
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
		return err
	}

	reporter = r
	if err := reporter.StartTestRun(context.Background()); err != nil {
		logging.Error("Failed to start test run: %v", err)
		return err
	}

		SetGlobalInitialized(true)
	SetGlobalReporter(r)
	
	// Set environment variable to indicate global initialization
	os.Setenv("QASE_GLOBAL_INITIALIZED", "true")
	
	logging.Info("Qase test run started successfully (global initialization)")
	return nil
}

// InitializeGlobalWithConfig initializes qase globally with a specific configuration
// This allows you to provide your own configuration instead of loading from file
func InitializeGlobalWithConfig(cfg *config.Config) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	if IsGlobalInitialized() {
		return nil // Already initialized
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
		return err
	}

	reporter = r
	if err := reporter.StartTestRun(context.Background()); err != nil {
		logging.Error("Failed to start test run: %v", err)
		return err
	}

		SetGlobalInitialized(true)
	SetGlobalReporter(r)
	
	// Set environment variable to indicate global initialization
	os.Setenv("QASE_GLOBAL_INITIALIZED", "true")
	
	logging.Info("Qase test run started successfully (global initialization with custom config)")
	return nil
}

// ResetGlobal resets the global initialization state (useful for testing)
func ResetGlobal() {
	ResetGlobalState()
	reporter = nil
	once = sync.Once{}
	// Clear environment variable
	os.Unsetenv("QASE_GLOBAL_INITIALIZED")
}

// Helper functions

// getActiveReporter returns the active reporter (global or local)
func getActiveReporter() *reporters.CoreReporter {
	// First check if global reporter is available
	globalReporter := GetGlobalReporter()
	if globalReporter != nil {
		return globalReporter
	}
	// Fallback to local reporter
	return reporter
}

func getTestFile() string {
	_, file, _, ok := runtime.Caller(2)
	if ok {
		return file
	}
	return "unknown"
}

func getTestLine() string {
	_, _, line, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("%d", line)
	}
	return "0"
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Helper functions for managing current test result

// setCurrentTestResult sets the current test result for error details capture
func setCurrentTestResult(result *domain.TestResult) {
	currentTestMutex.Lock()
	defer currentTestMutex.Unlock()
	currentTestResult = result
}

// clearCurrentTestResult clears the current test result
func clearCurrentTestResult() {
	currentTestMutex.Lock()
	defer currentTestMutex.Unlock()
	currentTestResult = nil
}

// getCurrentTestResult gets the current test result
func getCurrentTestResult() *domain.TestResult {
	currentTestMutex.RLock()
	defer currentTestMutex.RUnlock()
	return currentTestResult
}

// Helper functions for managing step stack

// pushStep pushes a step onto the stack
func pushStep(step *domain.TestStep) {
	stepStackMutex.Lock()
	defer stepStackMutex.Unlock()
	currentStepStack = append(currentStepStack, step)
}

// popStep pops a step from the stack
func popStep() *domain.TestStep {
	stepStackMutex.Lock()
	defer stepStackMutex.Unlock()

	if len(currentStepStack) == 0 {
		return nil
	}

	lastIndex := len(currentStepStack) - 1
	step := currentStepStack[lastIndex]
	currentStepStack = currentStepStack[:lastIndex]
	return step
}

// getCurrentStep gets the current step from the stack
func getCurrentStep() *domain.TestStep {
	stepStackMutex.RLock()
	defer stepStackMutex.RUnlock()

	if len(currentStepStack) == 0 {
		return nil
	}

	return currentStepStack[len(currentStepStack)-1]
}

// isAssertionError checks if the message looks like an assertion error
func isAssertionError(msg string) bool {
	// Common assertion error patterns
	errorPatterns := []string{
		"Should be true",
		"Should be false",
		"Not equal",
		"Should not be",
		"An error is expected but got nil",
		"Received unexpected error",
		"Error message not equal",
		"Should contain",
		"Should not contain",
		"Should be empty",
		"Should not be empty",
	}

	for _, pattern := range errorPatterns {
		if strings.Contains(msg, pattern) {
			return true
		}
	}
	return false
}

// extractErrorType extracts the error type from the error message
func extractErrorType(msg string) string {
	// Try to extract the main error type
	if strings.Contains(msg, "Should be true") {
		return "Should be true"
	}
	if strings.Contains(msg, "Should be false") {
		return "Should be false"
	}
	if strings.Contains(msg, "Not equal") {
		return "Not equal"
	}
	if strings.Contains(msg, "Should not be") {
		return "Should not be"
	}
	if strings.Contains(msg, "An error is expected but got nil") {
		return "Expected error but got nil"
	}
	if strings.Contains(msg, "Received unexpected error") {
		return "Unexpected error"
	}
	if strings.Contains(msg, "Error message not equal") {
		return "Error message not equal"
	}
	if strings.Contains(msg, "Should contain") {
		return "Should contain"
	}
	if strings.Contains(msg, "Should not contain") {
		return "Should not contain"
	}
	if strings.Contains(msg, "Should be empty") {
		return "Should be empty"
	}
	if strings.Contains(msg, "Should not be empty") {
		return "Should not be empty"
	}

	// Default to first line or first 50 characters
	if idx := strings.Index(msg, "\n"); idx > 0 {
		return strings.TrimSpace(msg[:idx])
	}
	if len(msg) > 50 {
		return strings.TrimSpace(msg[:50]) + "..."
	}
	return strings.TrimSpace(msg)
}

// getCallerInfo gets the file and line information of the caller
func getCallerInfo() (string, int) {
	// Skip the current function and AddMessage, go to the actual caller
	_, file, line, ok := runtime.Caller(2)
	if ok {
		return file, line
	}
	return "", 0
}
