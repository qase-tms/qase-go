package _go

import (
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

// TestMain is already defined in qase_simple_test.go

func TestSuiteSliceExample_SingleSuite(t *testing.T) {
	// Example with single suite
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Single Suite Test",
			Description: "Test with a single suite",
			Suite:       []string{"Authentication"},
		},
		func() {
			qase.AddMessage("Testing single suite functionality")
			qase.True(t, true)
		})
}

func TestSuiteSliceExample_NestedSuites(t *testing.T) {
	// Example with nested suites using slice
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Nested Suites Test",
			Description: "Test with nested suites using slice",
			Suite:       []string{"Authentication", "Login", "Valid Credentials"},
		},
		func() {
			qase.AddMessage("Testing nested suites functionality")
			qase.True(t, true)
		})
}

func TestSuiteSliceExample_DeepNesting(t *testing.T) {
	// Example with deep nesting
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Deep Nesting Test",
			Description: "Test with deep suite nesting",
			Suite:       []string{"E-commerce", "User Management", "Profile", "Settings", "Security"},
		},
		func() {
			qase.AddMessage("Testing deep nesting functionality")
			qase.True(t, true)
		})
}

func TestSuiteSliceExample_EmptySuite(t *testing.T) {
	// Example with empty suite - will auto-detect from file path
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Auto-detect Suite Test",
			Description: "Test with empty suite - will auto-detect from file path",
			Suite:       []string{}, // Empty slice
		},
		func() {
			qase.AddMessage("Testing auto-detection functionality")
			qase.True(t, true)
		})
}

func TestSuiteSliceExample_NilSuite(t *testing.T) {
	// Example with nil suite - will auto-detect from file path
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Nil Suite Test",
			Description: "Test with nil suite - will auto-detect from file path",
			Suite:       nil, // Nil slice
		},
		func() {
			qase.AddMessage("Testing nil suite functionality")
			qase.True(t, true)
		})
}

func TestSuiteSliceExample_WithSteps(t *testing.T) {
	// Example with steps and nested suites
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Suite with Steps Test",
			Description: "Test with steps and nested suites",
			Suite:       []string{"API Testing", "User API", "Create User"},
		},
		func() {
			qase.Step(t, qase.StepMetadata{
				Name:           "Prepare test data",
				ExpectedResult: "Test data is ready",
			}, func() {
				qase.AddMessage("Preparing test data")
				qase.True(t, true)
			})

			qase.Step(t, qase.StepMetadata{
				Name:           "Send API request",
				ExpectedResult: "API request is sent successfully",
			}, func() {
				qase.AddMessage("Sending API request")
				qase.True(t, true)
			})

			qase.Step(t, qase.StepMetadata{
				Name:           "Validate response",
				ExpectedResult: "Response is valid",
			}, func() {
				qase.AddMessage("Validating response")
				qase.True(t, true)
			})
		})
}
