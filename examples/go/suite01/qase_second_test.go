package suite01

import (
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

// TestMain is not needed here since global initialization is handled in the root package
// The qase package will automatically use the globally initialized reporter

func TestQaseSecond_Success(t *testing.T) {
	// This test will use the globally initialized qase reporter
	// No need for defer qase.CompleteTestRun() - it's handled globally

	qase.Test(t,
		qase.TestMetadata{
			DisplayName: "Second Suite Success Test",
			Description: "A test from the second suite that uses global initialization",
		},
		func() {
			qase.AddMessage("Starting second suite test")
			qase.True(t, true)
			qase.Equal(t, 2+2, 4)
			qase.AddMessage("Second suite test completed successfully")
		})
}

func TestQaseSecond_WithSteps(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			DisplayName: "Second Suite Test with Steps",
			Description: "Demonstrates steps in the second suite",
		},
		func() {
			qase.Step(t, qase.StepMetadata{
				Name:        "Second Suite Step 1",
				Description: "First step in second suite",
			}, func() {
				qase.AddMessage("Executing second suite step 1")
				qase.True(t, true)
			})

			qase.Step(t, qase.StepMetadata{
				Name:        "Second Suite Step 2",
				Description: "Second step in second suite",
			}, func() {
				qase.AddMessage("Executing second suite step 2")
				qase.Equal(t, 3*3, 9)
			})
		})
}

func TestQaseSecond_Parallel(t *testing.T) {
	t.Parallel()

	qase.Test(t,
		qase.TestMetadata{
			DisplayName: "Second Suite Parallel Test",
			Description: "A parallel test in the second suite",
		},
		func() {
			qase.AddMessage("Starting parallel test in second suite")
			qase.True(t, true)
			qase.AddMessage("Parallel test in second suite completed")
		})
}
