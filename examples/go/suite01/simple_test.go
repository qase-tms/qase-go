package suite01

import (
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

func TestMain(m *testing.M) {
	qase.TestMainForPackage(m)
}

func TestPass(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:   "Simple passing test",
			QaseIDs: []int64{101},
		},
		func() {
			qase.True(t, true)
		})
}

func TestFail(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:   "Simple failing test",
			QaseIDs: []int64{102},
		},
		func() {
			qase.True(t, false)
		})
}

func TestPanic(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:   "Test with panic",
			QaseIDs: []int64{103},
		},
		func() {
			panic("intentional panic for testing")
		})
}

func TestWithParameters(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:   "Test with parameters",
			QaseIDs: []int64{104},
			Parameters: map[string]string{
				"browser": "chrome",
				"version": "120",
			},
		},
		func() {
			qase.True(t, true)
		})
}

func TestWithFields(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:   "Test with custom fields",
			QaseIDs: []int64{105},
			Fields: map[string]string{
				"severity": "critical",
				"priority": "high",
				"layer":    "unit",
			},
		},
		func() {
			qase.True(t, true)
		})
}

func TestWithCustomSuite(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:   "Test with explicit suite",
			QaseIDs: []int64{106},
			Suite:   []string{"Custom Suite", "Child Suite"},
		},
		func() {
			qase.True(t, true)
		})
}

func TestWithAttachments(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:   "Test with attachments",
			QaseIDs: []int64{109},
		},
		func() {
			qase.AttachContent("log.txt", "Test log content", "text/plain")
			qase.True(t, true)
		})
}

func TestWithMultipleQaseIDs(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:   "Test mapped to multiple Qase IDs",
			QaseIDs: []int64{111, 112},
		},
		func() {
			qase.True(t, true)
		})
}

func TestWithTags(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			Title:   "Test with tags",
			QaseIDs: []int64{114},
			Tags:    []string{"smoke", "regression"},
		},
		func() {
			qase.True(t, true)
		})
}
