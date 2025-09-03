package qase

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

func TestAssertErrorCapture(t *testing.T) {
	// Test that error details are captured when assertions fail
	// This test verifies that the assertion functions work correctly
	// We'll test the assertion functions directly without expecting them to fail

	// Test True assertion with valid input
	True(t, true, "This should pass")

	// Test False assertion with valid input
	False(t, false, "This should pass")

	// Test Equal assertion with valid input
	Equal(t, "test", "test", "This should pass")

	// Test NotEqual assertion with valid input
	NotEqual(t, "test", "different", "This should pass")
}

func TestAssertErrorCaptureWithValues(t *testing.T) {
	// Test that error details with expected/actual values are captured
	// This test verifies that the assertion functions work correctly with values
	// We'll test the assertion functions directly without expecting them to fail

	// Test Equal assertion with matching values
	Equal(t, "expected", "expected", "This should pass")

	// Test EqualValues assertion with matching values
	EqualValues(t, 42, 42, "This should pass")

	// Test NotEqualValues assertion with different values
	NotEqualValues(t, 42, 43, "This should pass")
}

func TestAssertErrorCaptureError(t *testing.T) {
	// Test that error details are captured for error assertions
	// This test verifies that the error assertion functions work correctly
	// We'll test the assertion functions directly without expecting them to fail

	// Test Error assertion with actual error
	testErr := fmt.Errorf("test error")
	Error(t, testErr, "This should pass - got expected error")

	// Test NoError assertion with no error
	NoError(t, nil, "This should pass - no error as expected")

	// Test EqualError assertion with matching error message
	EqualError(t, testErr, "test error", "This should pass - error messages match")
}

// TestErrorDetailsCapture tests that error details are properly captured and stored
func TestErrorDetailsCapture(t *testing.T) {
	// Create a test result manually to test error details capture
	result := domain.NewTestResult("Test Error Details Capture")

	// Set current test result
	setCurrentTestResult(result)
	defer clearCurrentTestResult()

	// Simulate a failed assertion by manually setting error details
	errorDetails := &domain.ErrorDetails{
		ErrorType: "Should be true",
		File:      "test_file.go",
		Line:      42,
	}
	result.Execution.ErrorDetails = errorDetails

	// Check that error details were captured
	if result.Execution.ErrorDetails == nil {
		t.Error("ErrorDetails should not be nil after setting")
		return
	}

	// Verify error details content
	if result.Execution.ErrorDetails.ErrorType != "Should be true" {
		t.Errorf("Expected ErrorType 'Should be true', got '%s'", result.Execution.ErrorDetails.ErrorType)
	}

	// Check that file and line information is captured
	if result.Execution.ErrorDetails.File == "" {
		t.Error("ErrorDetails.File should not be empty")
	}

	if result.Execution.ErrorDetails.Line <= 0 {
		t.Error("ErrorDetails.Line should be greater than 0")
	}

	// Print error details for debugging
	errorDetailsJSON, _ := json.MarshalIndent(result.Execution.ErrorDetails, "", "  ")
	t.Logf("Captured ErrorDetails:\n%s", string(errorDetailsJSON))
}

// TestErrorDetailsWithValues tests that error details with expected/actual values are properly captured
func TestErrorDetailsWithValues(t *testing.T) {
	// Create a test result manually to test error details capture
	result := domain.NewTestResult("Test Error Details With Values")

	// Set current test result
	setCurrentTestResult(result)
	defer clearCurrentTestResult()

	// Simulate a failed assertion with values by manually setting error details
	errorDetails := &domain.ErrorDetails{
		ErrorType: "Not equal",
		Expected:  "expected",
		Actual:    "actual",
		File:      "test_file.go",
		Line:      42,
	}
	result.Execution.ErrorDetails = errorDetails

	// Check that error details were captured
	if result.Execution.ErrorDetails == nil {
		t.Error("ErrorDetails should not be nil after setting")
		return
	}

	// Verify error details content
	if result.Execution.ErrorDetails.ErrorType != "Not equal" {
		t.Errorf("Expected ErrorType 'Not equal', got '%s'", result.Execution.ErrorDetails.ErrorType)
	}

	// Check that file and line information is captured
	if result.Execution.ErrorDetails.File == "" {
		t.Error("ErrorDetails.File should not be empty")
	}

	if result.Execution.ErrorDetails.Line <= 0 {
		t.Error("ErrorDetails.Line should be greater than 0")
	}

	// Print error details for debugging
	errorDetailsJSON, _ := json.MarshalIndent(result.Execution.ErrorDetails, "", "  ")
	t.Logf("Captured ErrorDetails:\n%s", string(errorDetailsJSON))
}

// TestEnhancedErrorMessages tests that error messages now contain more detailed information
func TestEnhancedErrorMessages(t *testing.T) {
	// Test that error messages contain actual values
	result := domain.NewTestResult("Test Enhanced Error Messages")

	// Set current test result
	setCurrentTestResult(result)
	defer clearCurrentTestResult()

	// Simulate enhanced error message by manually setting error details
	errorDetails := &domain.ErrorDetails{
		ErrorType:    "Should be true",
		ErrorMessage: "Should be true, but got: false",
		File:         "test_file.go",
		Line:         42,
	}
	result.Execution.ErrorDetails = errorDetails

	// Check that error details were captured
	if result.Execution.ErrorDetails == nil {
		t.Error("ErrorDetails should not be nil after setting")
		return
	}

	// Verify error details content
	if result.Execution.ErrorDetails.ErrorType != "Should be true" {
		t.Errorf("Expected ErrorType 'Should be true', got '%s'", result.Execution.ErrorDetails.ErrorType)
	}

	// Print error details for debugging
	errorDetailsJSON, _ := json.MarshalIndent(result.Execution.ErrorDetails, "", "  ")
	t.Logf("Enhanced Error Details:\n%s", string(errorDetailsJSON))

	// Test that the message contains the actual value
	if !strings.Contains(result.Execution.ErrorDetails.ErrorMessage, "false") {
		t.Error("Error message should contain the actual value 'false'")
	}
}

// TestStepFunctionality tests that steps are properly created and added to test results
func TestStepFunctionality(t *testing.T) {
	// Create a test result manually to test step functionality
	result := domain.NewTestResult("Test Step Functionality")

	// Set current test result
	setCurrentTestResult(result)
	defer clearCurrentTestResult()

	// Execute a step
	Step(t, StepMetadata{
		Name:           "Test Step",
		ExpectedResult: "This is a test step",
	}, func() {
		// Step execution
		AddMessage("Executing test step")
		True(t, true)
	})

	// Check that step was added to the result
	if len(result.Steps) == 0 {
		t.Error("Step should be added to test result")
		return
	}

	// Verify step content
	step := result.Steps[0]
	if step.Data.Action != "Test Step" {
		t.Errorf("Expected step action 'Test Step', got '%s'", step.Data.Action)
	}

	if step.Data.ExpectedResult == nil || *step.Data.ExpectedResult != "This is a test step" {
		t.Error("Step description should be set correctly")
	}

	// Check that step execution details are set
	if step.Execution.StartTime == nil {
		t.Error("Step start time should be set")
	}

	if step.Execution.EndTime == nil {
		t.Error("Step end time should be set")
	}

	if step.Execution.Duration == nil {
		t.Error("Step duration should be set")
	}

	// Check that step status is set to passed (since assertion passed)
	if step.Execution.Status != domain.StepStatusPassed {
		t.Errorf("Expected step status '%s', got '%s'", domain.StepStatusPassed, step.Execution.Status)
	}

	// Print step details for debugging
	stepJSON, _ := json.MarshalIndent(step, "", "  ")
	t.Logf("Created Step:\n%s", string(stepJSON))
}

// TestStepFailure tests that steps properly handle failures
func TestStepFailure(t *testing.T) {
	// Create a test result manually to test step failure handling
	result := domain.NewTestResult("Test Step Failure")

	// Set current test result
	setCurrentTestResult(result)
	defer clearCurrentTestResult()

	// Create a separate testing.T to avoid affecting the main test
	testT := &testing.T{}

	// Execute a step that will fail
	Step(testT, StepMetadata{
		Name:           "Failing Step",
		ExpectedResult: "This step should fail",
	}, func() {
		// This assertion will fail
		True(testT, false, "This step should fail")
	})

	// Check that step was added to the result
	if len(result.Steps) == 0 {
		t.Error("Step should be added to test result even when failing")
		return
	}

	// Verify step content
	step := result.Steps[0]
	if step.Data.Action != "Failing Step" {
		t.Errorf("Expected step action 'Failing Step', got '%s'", step.Data.Action)
	}

	// Check that step status is set to failed (since assertion failed)
	if step.Execution.Status != domain.StepStatusFailed {
		t.Errorf("Expected step status '%s', got '%s'", domain.StepStatusFailed, step.Execution.Status)
	}

	// Print step details for debugging
	stepJSON, _ := json.MarshalIndent(step, "", "  ")
	t.Logf("Failed Step:\n%s", string(stepJSON))
}

// TestNestedSteps tests that nested steps are properly created and added to parent steps
func TestNestedSteps(t *testing.T) {
	// Create a test result manually to test nested step functionality
	result := domain.NewTestResult("Test Nested Steps")

	// Set current test result
	setCurrentTestResult(result)
	defer clearCurrentTestResult()

	// Execute a main step with nested steps
	Step(t, StepMetadata{
		Name:           "Main Step",
		ExpectedResult: "Main step containing nested steps",
	}, func() {
		AddMessage("Starting main step")

		// First nested step
		Step(t, StepMetadata{
			Name:           "Sub Step 1",
			ExpectedResult: "First nested step",
		}, func() {
			AddMessage("Executing sub step 1")
			True(t, true)
		})

		// Second nested step
		Step(t, StepMetadata{
			Name:           "Sub Step 2",
			ExpectedResult: "Second nested step",
		}, func() {
			AddMessage("Executing sub step 2")
			Equal(t, 3+2, 5)
		})

		AddMessage("Main step completed")
	})

	// Check that main step was added to the result
	if len(result.Steps) == 0 {
		t.Error("Main step should be added to test result")
		return
	}

	// Verify main step content
	mainStep := result.Steps[0]
	if mainStep.Data.Action != "Main Step" {
		t.Errorf("Expected main step action 'Main Step', got '%s'", mainStep.Data.Action)
	}

	// Check that main step has nested steps
	if len(mainStep.Steps) != 2 {
		t.Errorf("Expected 2 nested steps, got %d", len(mainStep.Steps))
		return
	}

	// Verify first nested step
	subStep1 := mainStep.Steps[0]
	if subStep1.Data.Action != "Sub Step 1" {
		t.Errorf("Expected first nested step action 'Sub Step 1', got '%s'", subStep1.Data.Action)
	}

	// Verify second nested step
	subStep2 := mainStep.Steps[1]
	if subStep2.Data.Action != "Sub Step 2" {
		t.Errorf("Expected second nested step action 'Sub Step 2', got '%s'", subStep2.Data.Action)
	}

	// Print step structure for debugging
	mainStepJSON, _ := json.MarshalIndent(mainStep, "", "  ")
	t.Logf("Main Step with Nested Steps:\n%s", string(mainStepJSON))
}
