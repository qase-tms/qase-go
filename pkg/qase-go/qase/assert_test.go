package qase

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

func TestAssertErrorCapture(t *testing.T) {
	// Test that error details are captured when assertions fail
	Test(t, TestMetadata{
		DisplayName: "Test Assert Error Capture",
		Title:       "Test Assert Error Capture",
		Description: "Test that error details are captured when assertions fail",
	}, func() {
		// This should fail and capture error details
		True(t, false, "This should fail")
	})
}

func TestAssertErrorCaptureWithValues(t *testing.T) {
	// Test that error details with expected/actual values are captured
	Test(t, TestMetadata{
		DisplayName: "Test Assert Error Capture With Values",
		Title:       "Test Assert Error Capture With Values",
		Description: "Test that error details with expected/actual values are captured",
	}, func() {
		// This should fail and capture error details with values
		Equal(t, "expected", "actual", "This should fail")
	})
}

func TestAssertErrorCaptureError(t *testing.T) {
	// Test that error details are captured for error assertions
	Test(t, TestMetadata{
		DisplayName: "Test Assert Error Capture Error",
		Title:       "Test Assert Error Capture Error",
		Description: "Test that error details are captured for error assertions",
	}, func() {
		// This should fail and capture error details
		Error(t, nil, "This should fail - expecting error but got nil")
	})
}

// TestErrorDetailsCapture tests that error details are properly captured and stored
func TestErrorDetailsCapture(t *testing.T) {
	// Create a test result manually to test error details capture
	result := domain.NewTestResult("Test Error Details Capture")

	// Set current test result
	setCurrentTestResult(result)
	defer clearCurrentTestResult()

	// Simulate a failed assertion by calling AddMessage directly
	AddMessage("Should be true")

	// Check that error details were captured
	if result.Execution.ErrorDetails == nil {
		t.Error("ErrorDetails should not be nil after failed assertion")
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

	// Simulate a failed assertion with values by calling AddMessage directly
	AddMessage("Not equal: \nexpected: expected\nactual  : actual")

	// Check that error details were captured
	if result.Execution.ErrorDetails == nil {
		t.Error("ErrorDetails should not be nil after failed assertion")
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

	// Test True function with enhanced message
	AddMessage("Should be true, but got: false")

	// Check that error details were captured
	if result.Execution.ErrorDetails == nil {
		t.Error("ErrorDetails should not be nil after failed assertion")
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
