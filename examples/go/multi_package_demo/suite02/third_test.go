package suite02

import (
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

func TestThirdSuiteTest(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			DisplayName: "Third Suite Test",
			Description: "A test in the third suite that uses global initialization",
		},
		func() {
			qase.AddMessage("Starting third suite test")
			qase.True(t, true)
			qase.Equal(t, 3+3, 6)
			qase.AddMessage("Third suite test completed successfully")
		})
}

func TestThirdSuiteParallel(t *testing.T) {
	t.Parallel()

	qase.Test(t,
		qase.TestMetadata{
			DisplayName: "Third Suite Parallel Test",
			Description: "A parallel test in the third suite",
		},
		func() {
			qase.AddMessage("Starting parallel test in third suite")
			qase.True(t, true)
			qase.AddMessage("Parallel test in third suite completed")
		})
}
