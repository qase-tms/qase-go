package suite02

import (
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

func TestMain(m *testing.M) {
	qase.TestMainForPackage(m)
}

func TestPassInSuite02(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:   "Pass in suite02",
			QaseIDs: []int64{201},
		},
		func() {
			qase.True(t, true)
		})
}

func TestWithAllFeatures(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:       "Comprehensive test with all features",
			Description: "Covers fields, params, suite, and steps",
			QaseIDs:     []int64{202},
			Suite:       []string{"Integration", "AllFeatures"},
			Fields: map[string]string{
				"severity": "normal",
			},
			Parameters: map[string]string{
				"env": "staging",
			},
		},
		func() {
			qase.Step(t, qase.StepMetadata{
				Name:           "Setup phase",
				ExpectedResult: "Ready for test",
			}, func() {
				qase.True(t, true)
			})

			qase.Step(t, qase.StepMetadata{
				Name:           "Execution phase",
				ExpectedResult: "Test logic runs",
			}, func() {
				qase.Equal(t, 2+2, 4)
				qase.AttachContent("result.json", `{"status":"ok"}`, "application/json")
			})

			qase.Step(t, qase.StepMetadata{
				Name:           "Verification phase",
				ExpectedResult: "Assertions pass",
			}, func() {
				qase.True(t, true)
			})
		})
}
