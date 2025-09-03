package _go

import (
	"sync"
	"testing"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

// TestMain automatically completes the test run after all tests finish
func TestMain(m *testing.M) {
	qase.TestMain(m)
}

func TestQaseSimple_Success(t *testing.T) {
	// Note: No defer qase.CompleteTestRun() needed!
	// Test run will be automatically completed after all tests finish

	qase.Test(t,
		qase.TestMetadata{
			Title:       "Simple Success Test",
			Description: "A basic test that demonstrates simple Qase reporting",
		},
		func() {
			qase.AddMessage("Starting simple test")
			qase.True(t, true)
			qase.Equal(t, 1+1, 2)
			qase.AddMessage("Simple test completed successfully")
		})
}

func TestQaseWithSteps_Success(t *testing.T) {
	// Note: No defer qase.CompleteTestRun() needed!

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

func TestQaseWithParameters(t *testing.T) {
	// Note: No defer qase.CompleteTestRun() needed!

	qase.Test(t,
		qase.TestMetadata{
			Title:       "Test with Parameters",
			Description: "Shows how to add parameters to test results",
			Parameters: map[string]string{
				"browser": "chrome",
				"version": "120",
				"os":      "macos",
				"screen":  "1920x1080",
			},
		},
		func() {
			qase.AddMessage("Testing with browser parameters")
			qase.True(t, true)
		})
}

func TestQaseWithLabels(t *testing.T) {
	// Note: No defer qase.CompleteTestRun() needed!

	qase.Test(t,
		qase.TestMetadata{
			Title:       "Test with Labels",
			Description: "Demonstrates adding labels to test results",
		},
		func() {
			qase.AddMessage("Testing with labels")
			qase.True(t, true)
		})
}

func TestQaseWithLinks(t *testing.T) {
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			Title:       "Test with Links",
			Description: "Shows how to add links to test results",
		},
		func() {
			qase.AddMessage("Testing with links")
			qase.AddMessage("Added link during test execution")
			qase.True(t, true)
		})
}

func TestQaseNestedSteps(t *testing.T) {
	defer qase.CompleteTestRun()

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

func TestQaseWithAttachments(t *testing.T) {
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			Title:       "Test with Attachments",
			Description: "Demonstrates adding attachments to test results",
		},
		func() {
			qase.AddMessage("Testing with attachments")
			qase.AddAttachments("test-data.json")
			qase.AddAttachments("screenshot.png")
			qase.True(t, true)
		})
}

func TestQaseWithExternalId(t *testing.T) {
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			Title:       "Test with External ID",
			Description: "Shows how to set external test ID",
		},
		func() {
			qase.AddMessage("Testing with external ID")
			qase.True(t, true)
		})
}

func TestQaseComplexExample(t *testing.T) {
	defer qase.CompleteTestRun()

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

func TestQaseFailing_Simple(t *testing.T) {
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			Title:       "Simple Failing Test",
			Description: "A simple test that should fail",
		},
		func() {
			qase.AddMessage("This test should fail")
			qase.True(t, false) // This will fail
			qase.AddMessage("This should not appear")
		})
}

func TestQaseFailing_PanicSimple(t *testing.T) {
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			Title:       "Panic Test",
			Description: "Test that demonstrates panic handling",
		},
		func() {
			qase.AddMessage("About to panic")
			panic("This is a test panic") // This will cause panic
			qase.AddMessage("This should not appear")
		})
}

func TestQaseFailing_EqualSimple(t *testing.T) {
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			Title:       "Failing Equal Test",
			Description: "Test that demonstrates failing equality assertion",
		},
		func() {
			qase.AddMessage("Testing equality assertion")
			qase.Equal(t, 1+1, 3) // This will fail
			qase.AddMessage("This should not appear")
		})
}

func TestQaseFailing_WithSteps(t *testing.T) {
	defer qase.CompleteTestRun()

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

func TestQaseFailing_WithAttachmentsSimple(t *testing.T) {
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			Title:       "Failing Test with Attachments",
			Description: "Test that fails but includes attachments",
		},
		func() {
			qase.AddMessage("Adding attachments before failure")
			qase.AddAttachments("error-log.txt")
			qase.AddAttachments("screenshot-error.png")

			qase.AddMessage("About to fail")
			qase.True(t, false) // This will fail

			qase.AddMessage("This should not appear")
		})
}

func TestQaseParallel_SimpleInline(t *testing.T) {
	t.Parallel()
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			Title:       "Parallel Simple Test",
			Description: "A simple parallel test",
		},
		func() {
			qase.AddMessage("Starting parallel test")
			time.Sleep(100 * time.Millisecond) // Simulate work
			qase.True(t, true)
			qase.AddMessage("Parallel test completed")
		})
}

func TestQaseParallel_WithStepsInline(t *testing.T) {
	t.Parallel()
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			Title:       "Parallel Test with Steps",
			Description: "Parallel test with multiple steps",
		},
		func() {
			qase.AddMessage("Starting parallel test with steps")

			qase.Step(t, qase.StepMetadata{
				Name:           "Parallel Step 1",
				ExpectedResult: "First parallel step",
			}, func() {
				time.Sleep(50 * time.Millisecond)
				qase.AddMessage("Step 1 completed")
				qase.True(t, true)
			})

			qase.Step(t, qase.StepMetadata{
				Name:           "Parallel Step 2",
				ExpectedResult: "Second parallel step",
			}, func() {
				time.Sleep(50 * time.Millisecond)
				qase.AddMessage("Step 2 completed")
				qase.Equal(t, 2+2, 4)
			})
		})
}

func TestQaseParallel_FailingInline(t *testing.T) {
	t.Parallel()
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			Title:       "Parallel Failing Test",
			Description: "A parallel test that should fail",
		},
		func() {
			qase.AddMessage("Starting parallel failing test")
			time.Sleep(100 * time.Millisecond)
			qase.True(t, false) // This will fail
			qase.AddMessage("This should not appear")
		})
}

func TestQaseParallel_ConcurrentInline(t *testing.T) {
	t.Parallel()
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			Title:       "Parallel Concurrent Test",
			Description: "Test with internal concurrency",
		},
		func() {
			qase.AddMessage("Starting concurrent test")

			var wg sync.WaitGroup
			results := make([]int, 3)

			// Launch 3 concurrent goroutines
			for i := 0; i < 3; i++ {
				wg.Add(1)
				go func(index int) {
					defer wg.Done()
					time.Sleep(50 * time.Millisecond)
					results[index] = index * 2
				}(i)
			}

			wg.Wait()

			qase.AddMessage("Concurrent operations completed")
			qase.Equal(t, results[0], 0)
			qase.Equal(t, results[1], 2)
			qase.Equal(t, results[2], 4)
		})
}
