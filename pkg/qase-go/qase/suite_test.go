package qase

import (
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

func TestApplyTestMetadata_WithSuiteSlice(t *testing.T) {
	// Test with empty suite slice
	meta := TestMetadata{
		Title: "Test with Empty Suite",
		Suite: []string{},
	}

	result := domain.NewTestResult("Test")
	applyTestMetadata(result, meta)

	// Should auto-detect suite from file path when Suite is empty
	if result.Relations == nil || result.Relations.Suite == nil {
		t.Error("Expected suite to be auto-detected from file path")
	}
}

func TestApplyTestMetadata_WithSingleSuite(t *testing.T) {
	// Test with single suite
	meta := TestMetadata{
		Title: "Test with Single Suite",
		Suite: []string{"My Suite"},
	}

	result := domain.NewTestResult("Test")
	applyTestMetadata(result, meta)

	// Verify suite data
	if result.Relations == nil || result.Relations.Suite == nil {
		t.Fatal("Expected Relations.Suite to be set")
	}

	if len(result.Relations.Suite.Data) != 1 {
		t.Errorf("Expected 1 suite data item, got %d", len(result.Relations.Suite.Data))
	}

	if result.Relations.Suite.Data[0].Title != "My Suite" {
		t.Errorf("Expected suite title 'My Suite', got %q", result.Relations.Suite.Data[0].Title)
	}
}

func TestApplyTestMetadata_WithNestedSuites(t *testing.T) {
	// Test with nested suites
	meta := TestMetadata{
		Title: "Test with Nested Suites",
		Suite: []string{"Parent Suite", "Child Suite", "Grandchild Suite"},
	}

	result := domain.NewTestResult("Test")
	applyTestMetadata(result, meta)

	// Verify suite data
	if result.Relations == nil || result.Relations.Suite == nil {
		t.Fatal("Expected Relations.Suite to be set")
	}

	if len(result.Relations.Suite.Data) != 3 {
		t.Errorf("Expected 3 suite data items, got %d", len(result.Relations.Suite.Data))
	}

	expectedTitles := []string{"Parent Suite", "Child Suite", "Grandchild Suite"}
	for i, expected := range expectedTitles {
		if result.Relations.Suite.Data[i].Title != expected {
			t.Errorf("Expected suite[%d] title %q, got %q", i, expected, result.Relations.Suite.Data[i].Title)
		}
	}
}

func TestApplyTestMetadata_WithSuiteWhitespace(t *testing.T) {
	// Test with suite titles that have whitespace
	meta := TestMetadata{
		Title: "Test with Whitespace",
		Suite: []string{"  Suite with spaces  ", "\tTabbed Suite\t", "\nNewline Suite\n"},
	}

	result := domain.NewTestResult("Test")
	applyTestMetadata(result, meta)

	// Verify suite data with trimmed whitespace
	if result.Relations == nil || result.Relations.Suite == nil {
		t.Fatal("Expected Relations.Suite to be set")
	}

	if len(result.Relations.Suite.Data) != 3 {
		t.Errorf("Expected 3 suite data items, got %d", len(result.Relations.Suite.Data))
	}

	expectedTitles := []string{"Suite with spaces", "Tabbed Suite", "Newline Suite"}
	for i, expected := range expectedTitles {
		if result.Relations.Suite.Data[i].Title != expected {
			t.Errorf("Expected suite[%d] title %q, got %q", i, expected, result.Relations.Suite.Data[i].Title)
		}
	}
}

func TestApplyTestMetadata_WithNilSuite(t *testing.T) {
	// Test with nil suite slice
	meta := TestMetadata{
		Title: "Test with Nil Suite",
		Suite: nil,
	}

	result := domain.NewTestResult("Test")
	applyTestMetadata(result, meta)

	// Should auto-detect suite from file path when Suite is nil
	if result.Relations == nil || result.Relations.Suite == nil {
		t.Error("Expected suite to be auto-detected from file path")
	}
}
