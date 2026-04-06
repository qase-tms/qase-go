package suite01

import (
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

func TestWithSteps(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:   "Test with steps",
			QaseIDs: []int64{107},
		},
		func() {
			qase.Step(t, qase.StepMetadata{
				Name:           "Step 1: Setup",
				ExpectedResult: "Setup completes",
			}, func() {
				qase.True(t, true)
			})

			qase.Step(t, qase.StepMetadata{
				Name:           "Step 2: Verify",
				ExpectedResult: "Verification passes",
			}, func() {
				qase.Equal(t, 2+2, 4)
			})
		})
}

func TestWithNestedSteps(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:   "Test with nested steps",
			QaseIDs: []int64{108},
		},
		func() {
			qase.Step(t, qase.StepMetadata{
				Name:           "Parent step",
				ExpectedResult: "Parent completes",
			}, func() {
				qase.Step(t, qase.StepMetadata{
					Name:           "Child step 1",
					ExpectedResult: "Child 1 passes",
				}, func() {
					qase.True(t, true)
				})

				qase.Step(t, qase.StepMetadata{
					Name:           "Child step 2",
					ExpectedResult: "Child 2 passes",
				}, func() {
					qase.Equal(t, 3+2, 5)
				})
			})
		})
}

func TestFailWithSteps(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:   "Failing test with steps",
			QaseIDs: []int64{110},
		},
		func() {
			qase.Step(t, qase.StepMetadata{
				Name:           "Setup step",
				ExpectedResult: "Setup succeeds",
			}, func() {
				qase.True(t, true)
			})

			qase.Step(t, qase.StepMetadata{
				Name:           "Failing step",
				ExpectedResult: "This step fails",
			}, func() {
				qase.Equal(t, "expected", "actual")
			})

			qase.Step(t, qase.StepMetadata{
				Name:           "Cleanup step",
				ExpectedResult: "Cleanup runs anyway",
			}, func() {
				qase.True(t, true)
			})
		})
}

func TestWithStepData(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:   "Test with step input data",
			QaseIDs: []int64{113},
		},
		func() {
			qase.Step(t, qase.StepMetadata{
				Name:           "Step with input data",
				ExpectedResult: "Data is processed",
				Data:           `{"key": "value"}`,
			}, func() {
				qase.True(t, true)
			})
		})
}
