package suite01

import (
	"sync"
	"testing"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

// Sum function for testing purposes
func Sum(a, b int) int {
	return a + b
}

// TestMain ensures that results are sent for this package
func TestMain(m *testing.M) {
	qase.TestMainForPackage(m)
}

func TestQaseSimple_Success(t *testing.T) {
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

func TestQase_WithSuites_Success(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Second Success Test with suite",
			Description: "A basic test that demonstrates simple Qase reporting",
			Suite:       []string{"suite01FormMetadada", "child_suite"},
		},
		func() {
			qase.AddMessage("Starting simple test")
			qase.True(t, true)
			qase.Equal(t, 1+1, 2)
			qase.AddMessage("Simple test completed successfully")
		})
}

func TestQaseWithParameters(t *testing.T) {
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

func TestQaseWithAttachments(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Test with Attachments",
			Description: "Demonstrates adding attachments to test results",
		},
		func() {
			qase.AddMessage("Testing with attachments")
			qase.AddAttachments("../go.mod")
			qase.AttachContent("log.txt", "My logs", "plain/text")
			qase.True(t, true)
		})
}

func TestQaseWithExternalId(t *testing.T) {
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

func TestQaseFailing_Simple(t *testing.T) {
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
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Panic Test",
			Description: "Test that demonstrates panic handling",
		},
		func() {
			qase.AddMessage("About to panic")
			panic("This is a test panic") // This will cause panic
		})
}

func TestQaseFailing_EqualSimple(t *testing.T) {
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

func TestQaseFailing_WithAttachmentsSimple(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Failing Test with Attachments",
			Description: "Test that fails but includes attachments",
		},
		func() {
			qase.AddMessage("Adding attachments before failure")
			qase.AddAttachments("../go.mod")
			qase.AttachContent("log.txt", "My logs", "plain/text")

			qase.AddMessage("About to fail")
			qase.True(t, false) // This will fail

			qase.AddMessage("This should not appear")
		})
}

func TestQaseParallel_SimpleInline(t *testing.T) {
	t.Parallel()
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
