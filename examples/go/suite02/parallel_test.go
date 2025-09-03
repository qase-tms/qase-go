package suite02

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
