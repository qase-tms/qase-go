package _go

import (
	"sync"
	"testing"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

func TestQaseSimple_Success(t *testing.T) {
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			DisplayName: "Simple Success Test",
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
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			DisplayName: "Test with Steps",
			Description: "Demonstrates how to use steps in Qase reporting",
		},
		func() {
			qase.Step(t, qase.StepMetadata{
				Name:        "Step 1",
				Description: "First step of the test",
			}, func() {
				qase.AddMessage("Executing step 1")
				qase.True(t, true)
			})

			qase.Step(t, qase.StepMetadata{
				Name:        "Step 2",
				Description: "Second step of the test",
			}, func() {
				qase.AddMessage("Executing step 2")
				qase.Equal(t, 2*2, 4)
			})
		})
}

func TestQaseWithParameters(t *testing.T) {
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			DisplayName: "Test with Parameters",
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
	defer qase.CompleteTestRun()

	qase.Test(t,
		qase.TestMetadata{
			DisplayName: "Test with Labels",
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
			DisplayName: "Test with Links",
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
			DisplayName: "Test with Nested Steps",
			Description: "Demonstrates nested step structure",
		},
		func() {
			qase.Step(t, qase.StepMetadata{
				Name:        "Main Step",
				Description: "Main step containing nested steps",
			}, func() {
				qase.AddMessage("Starting main step")

				qase.Step(t, qase.StepMetadata{
					Name:        "Sub Step 1",
					Description: "First nested step",
				}, func() {
					qase.AddMessage("Executing sub step 1")
					qase.True(t, true)
				})

				qase.Step(t, qase.StepMetadata{
					Name:        "Sub Step 2",
					Description: "Second nested step",
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
			DisplayName: "Test with Attachments",
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
			DisplayName: "Test with External ID",
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
			DisplayName: "Complex Test Example",
			Description: "A comprehensive example showing all Qase features",
			Parameters: map[string]string{
				"environment": "staging",
				"user":        "testuser",
			},
		},
		func() {
			qase.AddMessage("Starting complex test")

			qase.Step(t, qase.StepMetadata{
				Name:        "Setup",
				Description: "Test setup phase",
			}, func() {
				qase.AddMessage("Setting up test environment")
				time.Sleep(100 * time.Millisecond) // Simulate setup
				qase.True(t, true)
			})

			qase.Step(t, qase.StepMetadata{
				Name:        "Execute",
				Description: "Main test execution",
			}, func() {
				qase.AddMessage("Executing main test logic")
				qase.Equal(t, Sum(2, 3), 5)

				qase.Step(t, qase.StepMetadata{
					Name:        "Validation",
					Description: "Validate results",
				}, func() {
					qase.AddMessage("Validating test results")
					qase.True(t, true)
				})
			})

			qase.Step(t, qase.StepMetadata{
				Name:        "Cleanup",
				Description: "Test cleanup phase",
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
			DisplayName: "Simple Failing Test",
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
			DisplayName: "Panic Test",
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
			DisplayName: "Failing Equal Test",
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
			DisplayName: "Failing Test with Steps",
			Description: "Test with steps where one fails",
		},
		func() {
			qase.AddMessage("Starting test with steps")

			qase.Step(t, qase.StepMetadata{
				Name:        "Setup Step",
				Description: "This step should pass",
			}, func() {
				qase.AddMessage("Setting up")
				qase.True(t, true)
			})

			qase.Step(t, qase.StepMetadata{
				Name:        "Failing Step",
				Description: "This step should fail",
			}, func() {
				qase.AddMessage("About to fail")
				qase.Equal(t, "expected", "actual") // This will fail
				qase.AddMessage("This should not appear")
			})

			qase.Step(t, qase.StepMetadata{
				Name:        "Cleanup Step",
				Description: "This step should not execute",
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
			DisplayName: "Failing Test with Attachments",
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
			DisplayName: "Parallel Simple Test",
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
			DisplayName: "Parallel Test with Steps",
			Description: "Parallel test with multiple steps",
		},
		func() {
			qase.AddMessage("Starting parallel test with steps")

			qase.Step(t, qase.StepMetadata{
				Name:        "Parallel Step 1",
				Description: "First parallel step",
			}, func() {
				time.Sleep(50 * time.Millisecond)
				qase.AddMessage("Step 1 completed")
				qase.True(t, true)
			})

			qase.Step(t, qase.StepMetadata{
				Name:        "Parallel Step 2",
				Description: "Second parallel step",
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
			DisplayName: "Parallel Failing Test",
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
			DisplayName: "Parallel Concurrent Test",
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
