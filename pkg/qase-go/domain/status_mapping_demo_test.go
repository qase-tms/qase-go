package domain

import (
	"testing"
)

// TestStatusMappingDemo demonstrates the status mapping functionality
func TestStatusMappingDemo(t *testing.T) {
	t.Log("=== Qase Go Status Mapping Demo ===")

	// Example 1: Create status mapping from map
	t.Log("\n1. Creating status mapping from map:")
	mapping, err := NewStatusMapping(map[string]string{
		"invalid": "failed",
		"skipped": "passed",
	})
	if err != nil {
		t.Fatalf("Failed to create mapping: %v", err)
	}
	t.Logf("Created mapping: %s", mapping.String())

	// Example 2: Apply mapping to different statuses
	t.Log("\n2. Applying mapping to different statuses:")
	testCases := []struct {
		original TestResultStatus
		expected TestResultStatus
	}{
		{StatusInvalid, StatusFailed},
		{StatusSkipped, StatusPassed},
		{StatusPassed, StatusPassed}, // Should remain unchanged
		{StatusFailed, StatusFailed}, // Should remain unchanged
	}

	for _, tc := range testCases {
		result := mapping.ApplyMapping(tc.original)
		t.Logf("  %s -> %s (expected: %s)", tc.original, result, tc.expected)
		if result != tc.expected {
			t.Errorf("Expected %s, got %s", tc.expected, result)
		}
	}

	// Example 3: Apply mapping to test result
	t.Log("\n3. Applying mapping to test result:")
	testResult := &TestResult{
		Title: "Demo Test",
		Execution: TestExecution{
			Status: StatusInvalid,
		},
	}
	t.Logf("Original status: %s", testResult.Execution.Status)
	changed := mapping.ApplyMappingToResult(testResult)
	t.Logf("Status changed: %v", changed)
	t.Logf("New status: %s", testResult.Execution.Status)

	// Example 4: Environment variable parsing
	t.Log("\n4. Environment variable parsing:")
	envMapping, err := NewStatusMappingFromEnv("blocked=failed,skipped=passed")
	if err != nil {
		t.Fatalf("Failed to parse env mapping: %v", err)
	}
	t.Logf("Env mapping: %s", envMapping.String())

	// Example 5: Validation
	t.Log("\n5. Validation examples:")
	validMapping := map[string]string{
		"invalid": "failed",
		"skipped": "passed",
	}
	if err := ValidateMapping(validMapping); err != nil {
		t.Errorf("Valid mapping failed validation: %v", err)
	} else {
		t.Log("Valid mapping passed validation")
	}

	invalidMapping := map[string]string{
		"unknown": "failed", // Invalid source status
	}
	if err := ValidateMapping(invalidMapping); err != nil {
		t.Logf("Invalid mapping correctly rejected: %v", err)
	} else {
		t.Error("Invalid mapping should have failed validation")
	}

	t.Log("\n=== Demo completed successfully ===")
}
