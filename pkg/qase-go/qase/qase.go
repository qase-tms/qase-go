package qase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
	"github.com/qase-tms/qase-go/pkg/qase-go/reporters"
)

var (
	once     sync.Once
	reporter *reporters.CoreReporter
)

func init() {
	once.Do(func() {
		cfg, err := config.Load()
		if err != nil {
			// Fallback to unsafe loading if validation fails
			log.Printf("Warning: Configuration validation failed, using unsafe loading: %v", err)
			cfg = config.LoadUnsafe()
		}

		// Log configuration as JSON
		cfgJSON, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			log.Printf("Failed to marshal config to JSON: %v", err)
		} else {
			log.Printf("Qase configuration loaded:\n%s", string(cfgJSON))
		}

		r, err := reporters.NewCoreReporter(cfg)
		if err != nil {
			log.Printf("Error: Failed to create core reporter: %v", err)
			return
		}

		reporter = r
		if err := reporter.StartTestRun(context.Background()); err != nil {
			log.Printf("Error: Failed to start test run: %v", err)
			return
		}

		log.Printf("Qase test run started successfully")
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

	defer func() {
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
			msg := "Test failed due to assertion failure"
			result.Message = &msg

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
		if reporter != nil {
			_ = reporter.AddResult(result)
		}
	}()

	fn()
}

// Step creates a test step
func Step(meta StepMetadata, fn func()) {
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

	defer func() {
		if r := recover(); r != nil {
			// Panic occurred in step - set status to failed
			step.Execution.Status = domain.StepStatusFailed
		}

		// End step execution
		end := time.Now().UnixMilli()
		step.Execution.EndTime = &end
		duration := end - now
		step.Execution.Duration = &duration
	}()

	fmt.Printf("Step: %s - %s\n", meta.Name, meta.Description)
	fn()
}

// AddMessage adds a message to the current test
func AddMessage(msg string) {
	fmt.Println("Message:", msg)
}

// AddAttachments adds attachments to the current test
func AddAttachments(path string) {
	fmt.Println("Attachment:", path)
}

// AttachFile attaches a file with specific content type
func AttachFile(name, filePath, contentType string) {
	fmt.Printf("File Attachment: %s (%s) - %s\n", name, contentType, filePath)
}

// AttachContent attaches content with specific content type
func AttachContent(name, content, contentType string) {
	fmt.Printf("Content Attachment: %s (%s) - %s\n", name, contentType, content[:min(len(content), 50)])
}

// CompleteTestRun completes the test run
func CompleteTestRun() error {
	if reporter != nil {
		return reporter.CompleteTestRun(context.Background())
	}
	return nil
}

// Helper functions
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
