package suite01

import (
	"testing"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

func TestQaseWithSteps_Success(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Test with Steps",
			Description: "Demonstrates how to use steps in Qase reporting",
		},
		func() {
			qase.Step(t, qase.StepMetadata{
				Name:           "Step 1",
				ExpectedResult: "First step of the test",
			}, func() {
				qase.AddMessage("Executing step 1")
				qase.True(t, true)
			})

			qase.Step(t, qase.StepMetadata{
				Name:           "Step 2",
				ExpectedResult: "Second step of the test",
			}, func() {
				qase.AddMessage("Executing step 2")
				qase.Equal(t, 2*2, 4)
			})
		})
}

func TestQaseNestedSteps(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Test with Nested Steps",
			Description: "Demonstrates nested step structure",
		},
		func() {
			qase.Step(t, qase.StepMetadata{
				Name:           "Main Step",
				ExpectedResult: "Main step containing nested steps",
			}, func() {
				qase.AddMessage("Starting main step")

				qase.Step(t, qase.StepMetadata{
					Name:           "Sub Step 1",
					ExpectedResult: "First nested step",
				}, func() {
					qase.AddMessage("Executing sub step 1")
					qase.True(t, true)
				})

				qase.Step(t, qase.StepMetadata{
					Name:           "Sub Step 2",
					ExpectedResult: "Second nested step",
				}, func() {
					qase.AddMessage("Executing sub step 2")
					qase.Equal(t, 3+2, 5)
				})
			})
		})
}

func TestQaseComplexExample(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Complex Test Example",
			Description: "A comprehensive example showing all Qase features",
			Parameters: map[string]string{
				"environment": "staging",
				"user":        "testuser",
			},
		},
		func() {
			qase.AddMessage("Starting complex test")

			qase.Step(t, qase.StepMetadata{
				Name:           "Setup",
				ExpectedResult: "Test setup phase",
			}, func() {
				qase.AddMessage("Setting up test environment")
				time.Sleep(100 * time.Millisecond) // Simulate setup
				qase.True(t, true)
			})

			qase.Step(t, qase.StepMetadata{
				Name:           "Execute",
				ExpectedResult: "Main test execution",
			}, func() {
				qase.AddMessage("Executing main test logic")
				qase.Equal(t, Sum(2, 3), 5)

				qase.Step(t, qase.StepMetadata{
					Name:           "Validation",
					ExpectedResult: "Validate results",
				}, func() {
					qase.AddMessage("Validating test results")
					qase.True(t, true)
				})
			})

			qase.Step(t, qase.StepMetadata{
				Name:           "Cleanup",
				ExpectedResult: "Test cleanup phase",
			}, func() {
				qase.AddMessage("Cleaning up test resources")
				qase.AddAttachments("test-report.json")
				qase.True(t, true)
			})

			qase.AddMessage("Complex test completed successfully")
		})
}

func TestQaseFailing_WithSteps(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Failing Test with Steps",
			Description: "Test with steps where one fails",
		},
		func() {
			qase.AddMessage("Starting test with steps")

			qase.Step(t, qase.StepMetadata{
				Name:           "Setup Step",
				ExpectedResult: "This step should pass",
			}, func() {
				qase.AddMessage("Setting up")
				qase.True(t, true)
			})

			qase.Step(t, qase.StepMetadata{
				Name:           "Failing Step",
				ExpectedResult: "This step should fail",
			}, func() {
				qase.AddMessage("About to fail")
				qase.Equal(t, "expected", "actual") // This will fail
				qase.AddMessage("This should not appear")
			})

			qase.Step(t, qase.StepMetadata{
				Name:           "Cleanup Step",
				ExpectedResult: "This step should not execute",
			}, func() {
				qase.AddMessage("Cleaning up")
				qase.True(t, true)
			})
		})
}
