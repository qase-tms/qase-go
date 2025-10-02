package qase

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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
	reporter          *reporters.CoreReporter
	once              sync.Once
	currentTestResult *domain.TestResult
	currentTestMutex  sync.RWMutex
	currentStepStack  []*domain.TestStep
	stepStackMutex    sync.RWMutex
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
		loggingConfig := &logging.Config{
			Debug: cfg.Debug,
			Logging: logging.LoggingConfig{
				Console: cfg.Logging.Console,
				File:    cfg.Logging.File,
			},
		}
		if err := logging.InitFromConfig(loggingConfig); err != nil {
			logging.Warn("Warning: Failed to initialize logging system: %v", err)
		} else {
			logging.Debug("Logging system initialized successfully")
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
		logging.Debug("Qase reporter initialized successfully")
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

	// Use Title as title, fallback to test name if Title is empty
	title := meta.Title
	if title == "" {
		title = t.Name() // Use test name as fallback
	}

	// Create test result
	result := domain.NewTestResult(title)
	result.Execution.Status = domain.StatusPassed

	// Set test start time
	startTime := time.Now().UnixMilli()
	result.Execution.StartTime = &startTime

	// Apply metadata from TestMetadata to TestResult
	applyTestMetadata(result, meta)

	// Set current test result for step collection
	setCurrentTestResult(result)
	defer func() {
		clearCurrentTestResult()
		clearStepStack()
	}()

	defer func() {
		// Set test end time and duration
		endTime := time.Now().UnixMilli()
		result.Execution.EndTime = &endTime
		if result.Execution.StartTime != nil {
			duration := endTime - *result.Execution.StartTime
			result.Execution.Duration = &duration
		}

		if r := recover(); r != nil {
			// Panic occurred - set status to invalid and capture stacktrace
			result.Execution.Status = domain.StatusInvalid
			msg := fmt.Sprintf("panic: %v", r)
			result.AddSystemMessage(msg)

			// Capture stacktrace for panic
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			stacktrace := string(buf[:n])
			result.Execution.Stacktrace = &stacktrace

		} else if t.Failed() {
			// Test failed due to assertion failure
			result.Execution.Status = domain.StatusFailed
			msg := "Test failed due to assertion failure"
			result.AddSystemMessage(msg)

			// Capture stacktrace for failed test
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			stacktrace := string(buf[:n])
			result.Execution.Stacktrace = &stacktrace

		} else {
			result.Execution.Status = domain.StatusPassed
		}

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

	// Use Title as title, fallback to test name if Title is empty
	title := meta.Title
	if title == "" {
		title = t.Name() // Use test name as fallback
	}

	// Create test result
	result := domain.NewTestResult(title)
	result.Execution.Status = domain.StatusPassed

	// Set test start time
	startTime := time.Now().UnixMilli()
	result.Execution.StartTime = &startTime

	// Apply metadata from TestMetadata to TestResult
	applyTestMetadata(result, meta)

	// Set current test result for step collection
	setCurrentTestResult(result)
	defer func() {
		clearCurrentTestResult()
		clearStepStack()
	}()

	// Create step builder
	builder := &TestStepBuilder{
		result: result,
	}

	defer func() {
		// Set test end time and duration
		endTime := time.Now().UnixMilli()
		result.Execution.EndTime = &endTime
		if result.Execution.StartTime != nil {
			duration := endTime - *result.Execution.StartTime
			result.Execution.Duration = &duration
		}
		if r := recover(); r != nil {
			// Panic occurred - set status to invalid and capture stacktrace
			result.Execution.Status = domain.StatusInvalid
			msg := fmt.Sprintf("panic: %v", r)
			result.AddSystemMessage(msg)

			// Capture stacktrace for panic
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			stacktrace := string(buf[:n])
			result.Execution.Stacktrace = &stacktrace

		} else if t.Failed() {
			// Test failed due to assertion failure
			result.Execution.Status = domain.StatusFailed
			msg := "Test failed due to assertion failure"
			result.AddSystemMessage(msg)

			// Capture stacktrace for failed test
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			stacktrace := string(buf[:n])
			result.Execution.Stacktrace = &stacktrace

		} else {
			result.Execution.Status = domain.StatusPassed
		}

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

	// Get current test result
	currentResult := getCurrentTestResult()
	if currentResult == nil {
		// No current test result, just log the message
		logging.Debug("Test message: %s (no current test result)", message)
		return
	}

	// Add message to current test result
	currentResult.AddMessage(message)
	logging.Debug("Test message: %s", message)
}

// Step executes a test step with metadata
func Step(t *testing.T, meta StepMetadata, fn func()) {
	// Get current test result
	currentResult := getCurrentTestResult()
	if currentResult == nil {
		// No current test result, just log and execute
		logging.Debug("Executing step: %s (no current test result)", meta.Name)
		if meta.ExpectedResult != "" {
			logging.Debug("Step expected result: %s", meta.ExpectedResult)
		}
		fn()
		return
	}

	// Create step
	step := domain.NewTestStep(meta.Name)
	if meta.ExpectedResult != "" {
		step.SetExpectedResult(meta.ExpectedResult)
	}
	if meta.Data != "" {
		step.SetData(meta.Data)
	}

	// Set timing
	startTime := time.Now().UnixMilli()
	step.Execution.StartTime = &startTime

	// Push step to stack for nested step support
	pushCurrentStep(step)

	// Execute the step function
	defer func() {
		// Pop step from stack
		poppedStep := popCurrentStep()
		if poppedStep != step {
			logging.Warn("Step stack mismatch: expected %p, got %p", step, poppedStep)
		}

		endTime := time.Now().UnixMilli()
		step.Execution.EndTime = &endTime
		duration := endTime - startTime
		step.Execution.Duration = &duration

		// Determine step status based on test failure
		if t.Failed() {
			step.Execution.Status = domain.StepStatusFailed
		} else {
			step.Execution.Status = domain.StepStatusPassed
		}

		// Add step to parent or test result
		parentStep := getCurrentStep()
		if parentStep != nil {
			// Add to parent step
			parentStep.AddStep(*step)
		} else {
			// Add to test result
			currentResult.AddStep(*step)
		}
	}()

	// Log step execution
	logging.Debug("Executing step: %s", meta.Name)
	if meta.ExpectedResult != "" {
		logging.Debug("Step expected result: %s", meta.ExpectedResult)
	}

	// Execute the step function
	fn()
}

// AddAttachments adds file attachments to the current test result
func AddAttachments(filePaths ...string) {
	if reporter == nil {
		return
	}

	// Get current test result
	currentResult := getCurrentTestResult()
	if currentResult == nil {
		// No current test result, just log the attachments
		for _, path := range filePaths {
			logging.Debug("Adding attachment: %s (no current test result)", path)
		}
		return
	}

	// Add attachments to current test result
	for _, path := range filePaths {
		attachment := domain.Attachment{
			FileName: path,
		}
		attachment.SetFilePath(path)
		currentResult.AddAttachment(attachment)
		logging.Debug("Adding attachment: %s", path)
	}
}

// AttachContent adds content as an attachment to the current test result
func AttachContent(name, content, mimeType string) {
	if reporter == nil {
		return
	}

	// Get current test result
	currentResult := getCurrentTestResult()
	if currentResult == nil {
		// No current test result, just log the content attachment
		logging.Debug("Adding content attachment: %s (type: %s) (no current test result)", name, mimeType)
		return
	}

	// Add content attachment to current test result
	attachment := domain.Attachment{
		FileName: name,
		Content:  []byte(content),
		MimeType: mimeType,
	}
	currentResult.AddAttachment(attachment)
	logging.Debug("Adding content attachment: %s (type: %s)", name, mimeType)
}

// setCurrentTestResult sets the current test result for step collection
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

// pushCurrentStep pushes a step to the current step stack
func pushCurrentStep(step *domain.TestStep) {
	stepStackMutex.Lock()
	defer stepStackMutex.Unlock()
	currentStepStack = append(currentStepStack, step)
}

// popCurrentStep pops a step from the current step stack
func popCurrentStep() *domain.TestStep {
	stepStackMutex.Lock()
	defer stepStackMutex.Unlock()
	if len(currentStepStack) == 0 {
		return nil
	}
	step := currentStepStack[len(currentStepStack)-1]
	currentStepStack = currentStepStack[:len(currentStepStack)-1]
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

// clearStepStack clears the step stack
func clearStepStack() {
	stepStackMutex.Lock()
	defer stepStackMutex.Unlock()
	currentStepStack = nil
}

// InitializeGlobal initializes qase globally (for backward compatibility)
// This function is kept for backward compatibility but doesn't do anything special
// since initialization now happens automatically in init()
func InitializeGlobal() error {
	// Initialization already happened in init(), just return success
	// But we need to ensure reporter is set for testing
	if reporter == nil {
		// Load configuration from root directory
		cfg, err := config.LoadWithParentSearch()
		if err != nil {
			// Fallback to unsafe loading if validation fails
			logging.Warn("Warning: Configuration validation failed, using unsafe loading: %v", err)
			cfg = config.LoadUnsafeWithParentSearch()
		}

		r, err := reporters.NewCoreReporter(cfg)
		if err != nil {
			logging.Error("Failed to create core reporter: %v", err)
			return err
		}

		reporter = r
	}
	return nil
}

// InitializeGlobalWithConfig initializes qase globally with custom config
func InitializeGlobalWithConfig(cfg *config.Config) error {
	// Create reporter with the provided config
	r, err := reporters.NewCoreReporter(cfg)
	if err != nil {
		logging.Error("Failed to create core reporter with custom config: %v", err)
		return err
	}

	reporter = r
	return nil
}

// IsGlobalInitialized returns true if qase is globally initialized
func IsGlobalInitialized() bool {
	return reporter != nil
}

// ResetGlobal resets the global state (for testing purposes)
func ResetGlobal() {
	once = sync.Once{}
	reporter = nil
	currentTestMutex.Lock()
	currentTestResult = nil
	currentTestMutex.Unlock()
	clearStepStack()
}

// TestMain is a function that should be called in test main to properly complete the test run
func TestMain(m *testing.M) {
	// Run the tests
	code := m.Run()

	// Complete the test run and send all results
	if err := CompleteTestRun(); err != nil {
		logging.Error("Failed to complete test run: %v", err)
	}

	// Exit with the test result code
	os.Exit(code)
}

// TestMainForPackage is a function that should be called in each package's test main
// This ensures that results are sent for each package individually
func TestMainForPackage(m *testing.M) {
	// Run the tests
	code := m.Run()

	// Complete the test run and send all results for this package
	if err := CompleteTestRun(); err != nil {
		logging.Error("Failed to complete test run for package: %v", err)
	}

	// Exit with the test result code
	os.Exit(code)
}

// CompleteTestRun completes the test run and sends all results
func CompleteTestRun() error {
	if reporter != nil {
		return reporter.CompleteTestRun(context.Background())
	}
	return nil
}

// GetReporter returns the current reporter instance
func GetReporter() *reporters.CoreReporter {
	return reporter
}

// findProjectRoot finds the project root by looking for go.mod file
func findProjectRoot(filePath string) string {
	dir := filepath.Dir(filePath)

	for {
		// Check if go.mod exists in current directory
		goModPath := filepath.Join(dir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return dir
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root directory
			break
		}
		dir = parent
	}

	// Fallback to current working directory
	wd, err := os.Getwd()
	if err != nil {
		return "."
	}
	return wd
}

// GetTestFileInfo extracts file path information from the call stack
func GetTestFileInfo() (string, string) {
	// Get more frames to find the test file
	pc := make([]uintptr, 30)
	n := runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc[:n])

	for {
		frame, more := frames.Next()
		if !more {
			break
		}

		// Look for test files (files ending with _test.go)
		// Skip frames from qase package itself, but allow test files in qase package
		if strings.HasSuffix(frame.File, "_test.go") && !strings.Contains(frame.File, "/qase/qase.go") {
			// Try to find the project root by looking for go.mod
			projectRoot := findProjectRoot(frame.File)

			// Get relative path from project root
			relPath, err := filepath.Rel(projectRoot, frame.File)
			if err != nil {
				// Fallback to current working directory
				wd, err := os.Getwd()
				if err != nil {
					wd = "."
				}
				relPath, err = filepath.Rel(wd, frame.File)
				if err != nil {
					relPath = frame.File
				}
			}

			// Extract directory and filename
			dir := filepath.Dir(relPath)
			filename := filepath.Base(relPath)

			// Remove _test.go suffix
			suiteName := strings.TrimSuffix(filename, "_test.go")

			return dir, suiteName
		}
	}
	return ".", "unknown"
}

// applyTestMetadata applies metadata from TestMetadata to TestResult
func applyTestMetadata(result *domain.TestResult, meta TestMetadata) {
	// Set description
	if meta.Description != "" {
		result.SetField("description", meta.Description)
	}

	// Set comment
	if meta.Comment != "" {
		result.AddMessage(meta.Comment)
	}

	// Set fields
	if meta.Fields != nil {
		for key, value := range meta.Fields {
			result.SetField(key, value)
		}
	}

	// Set parameters
	if meta.Parameters != nil {
		for key, value := range meta.Parameters {
			result.SetParam(key, value)
		}
	}

	// Set group parameters
	if meta.GroupParameters != nil {
		for key, value := range meta.GroupParameters {
			result.SetGroupParam(key, value)
		}
	}

	// Set suite
	var suiteData []domain.SuiteData

	if len(meta.Suite) > 0 {
		// Use suite titles directly from the slice
		suiteData = make([]domain.SuiteData, len(meta.Suite))
		for i, title := range meta.Suite {
			suiteData[i] = domain.SuiteData{
				Title: strings.TrimSpace(title),
			}
		}
	} else {
		// Auto-detect suite from file path
		dir, filename := GetTestFileInfo()

		// Create suite hierarchy from directory path
		var suiteTitles []string

		// Add directory components as suite hierarchy
		if dir != "." && dir != "" {
			dirParts := strings.Split(dir, string(filepath.Separator))
			for _, part := range dirParts {
				if part != "" && part != "." {
					suiteTitles = append(suiteTitles, part)
				}
			}
		}

		// Add filename as the final suite
		suiteTitles = append(suiteTitles, filename)

		// Create suite data
		suiteData = make([]domain.SuiteData, len(suiteTitles))
		for i, title := range suiteTitles {
			suiteData[i] = domain.SuiteData{
				Title: title,
			}
		}
	}

	if len(suiteData) > 0 {
		result.SetSuite(suiteData)
	}
}
