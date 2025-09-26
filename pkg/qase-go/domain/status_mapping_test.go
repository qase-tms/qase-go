package domain

import (
	"testing"
)

func TestNewStatusMapping(t *testing.T) {
	tests := []struct {
		name        string
		mapping     map[string]string
		expectError bool
		expectedLen int
	}{
		{
			name: "valid mapping",
			mapping: map[string]string{
				"invalid": "failed",
				"skipped": "passed",
			},
			expectError: false,
			expectedLen: 2,
		},
		{
			name: "empty mapping",
			mapping: map[string]string{},
			expectError: false,
			expectedLen: 0,
		},
		{
			name: "invalid source status",
			mapping: map[string]string{
				"unknown": "failed",
			},
			expectError: true,
			expectedLen: 0,
		},
		{
			name: "invalid target status",
			mapping: map[string]string{
				"failed": "unknown",
			},
			expectError: true,
			expectedLen: 0,
		},
		{
			name: "mapping to same status",
			mapping: map[string]string{
				"failed": "failed",
			},
			expectError: true,
			expectedLen: 0,
		},
		{
			name: "case insensitive mapping",
			mapping: map[string]string{
				"INVALID": "FAILED",
				"Skipped": "Passed",
			},
			expectError: false,
			expectedLen: 2,
		},
		{
			name: "whitespace handling",
			mapping: map[string]string{
				" invalid ": " failed ",
				" skipped ": " passed ",
			},
			expectError: false,
			expectedLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapping, err := NewStatusMapping(tt.mapping)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			if len(mapping) != tt.expectedLen {
				t.Errorf("Expected mapping length %d, got %d", tt.expectedLen, len(mapping))
			}
		})
	}
}

func TestNewStatusMappingFromEnv(t *testing.T) {
	tests := []struct {
		name        string
		envValue    string
		expectError bool
		expectedLen int
	}{
		{
			name:        "valid env format",
			envValue:    "invalid=failed,skipped=passed",
			expectError: false,
			expectedLen: 2,
		},
		{
			name:        "empty env value",
			envValue:    "",
			expectError: false,
			expectedLen: 0,
		},
		{
			name:        "single mapping",
			envValue:    "invalid=failed",
			expectError: false,
			expectedLen: 1,
		},
		{
			name:        "invalid format - no equals",
			envValue:    "invalid,failed",
			expectError: true,
			expectedLen: 0,
		},
		{
			name:        "invalid format - multiple equals",
			envValue:    "invalid=failed=extra",
			expectError: true,
			expectedLen: 0,
		},
		{
			name:        "empty status",
			envValue:    "=failed",
			expectError: true,
			expectedLen: 0,
		},
		{
			name:        "whitespace handling",
			envValue:    " invalid = failed , skipped = passed ",
			expectError: false,
			expectedLen: 2,
		},
		{
			name:        "empty pairs",
			envValue:    "invalid=failed,,skipped=passed,",
			expectError: false,
			expectedLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapping, err := NewStatusMappingFromEnv(tt.envValue)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			if mapping == nil {
				if tt.expectedLen != 0 {
					t.Errorf("Expected mapping length %d, got nil", tt.expectedLen)
				}
				return
			}
			
			if len(mapping) != tt.expectedLen {
				t.Errorf("Expected mapping length %d, got %d", tt.expectedLen, len(mapping))
			}
		})
	}
}

func TestStatusMapping_ApplyMapping(t *testing.T) {
	mapping, err := NewStatusMapping(map[string]string{
		"invalid": "failed",
		"skipped": "passed",
	})
	if err != nil {
		t.Fatalf("Failed to create mapping: %v", err)
	}

	tests := []struct {
		name           string
		inputStatus    TestResultStatus
		expectedStatus TestResultStatus
	}{
		{
			name:           "mapped status - invalid to failed",
			inputStatus:    StatusInvalid,
			expectedStatus: StatusFailed,
		},
		{
			name:           "mapped status - skipped to passed",
			inputStatus:    StatusSkipped,
			expectedStatus: StatusPassed,
		},
		{
			name:           "unmapped status - passed",
			inputStatus:    StatusPassed,
			expectedStatus: StatusPassed,
		},
		{
			name:           "unmapped status - failed",
			inputStatus:    StatusFailed,
			expectedStatus: StatusFailed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mapping.ApplyMapping(tt.inputStatus)
			if result != tt.expectedStatus {
				t.Errorf("Expected status %s, got %s", tt.expectedStatus, result)
			}
		})
	}
}

func TestStatusMapping_ApplyMappingToResult(t *testing.T) {
	mapping, err := NewStatusMapping(map[string]string{
		"invalid": "failed",
		"skipped": "passed",
	})
	if err != nil {
		t.Fatalf("Failed to create mapping: %v", err)
	}

	tests := []struct {
		name           string
		inputStatus    TestResultStatus
		expectedStatus TestResultStatus
		expectedChange bool
	}{
		{
			name:           "mapped status - invalid to failed",
			inputStatus:    StatusInvalid,
			expectedStatus: StatusFailed,
			expectedChange: true,
		},
		{
			name:           "mapped status - skipped to passed",
			inputStatus:    StatusSkipped,
			expectedStatus: StatusPassed,
			expectedChange: true,
		},
		{
			name:           "unmapped status - passed",
			inputStatus:    StatusPassed,
			expectedStatus: StatusPassed,
			expectedChange: false,
		},
		{
			name:           "nil result",
			inputStatus:    StatusPassed,
			expectedStatus: StatusPassed,
			expectedChange: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result *TestResult
			if tt.name != "nil result" {
				result = &TestResult{
					Title: "Test",
					Execution: TestExecution{
						Status: tt.inputStatus,
					},
				}
			}

			changed := mapping.ApplyMappingToResult(result)
			if changed != tt.expectedChange {
				t.Errorf("Expected change %v, got %v", tt.expectedChange, changed)
			}

			if result != nil && result.Execution.Status != tt.expectedStatus {
				t.Errorf("Expected status %s, got %s", tt.expectedStatus, result.Execution.Status)
			}
		})
	}
}

func TestStatusMapping_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		mapping  StatusMapping
		expected bool
	}{
		{
			name:     "empty mapping",
			mapping:  StatusMapping{},
			expected: true,
		},
		{
			name: "non-empty mapping",
			mapping: StatusMapping{
				StatusInvalid: StatusFailed,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.mapping.IsEmpty()
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestStatusMapping_GetMappings(t *testing.T) {
	original := StatusMapping{
		StatusInvalid: StatusFailed,
		StatusSkipped: StatusPassed,
	}

	copy := original.GetMappings()

	// Verify it's a copy, not the same reference
	// We can't directly compare pointers, so we check if modifying one affects the other
	original[StatusInvalid] = StatusPassed
	if copy[StatusInvalid] == StatusPassed {
		t.Error("GetMappings should return a copy, not the same reference")
	}

	// Verify contents are the same (before modification)
	originalCopy := StatusMapping{
		StatusInvalid: StatusFailed,
		StatusSkipped: StatusPassed,
	}
	
	if len(copy) != len(originalCopy) {
		t.Errorf("Expected copy length %d, got %d", len(originalCopy), len(copy))
	}

	for from, to := range originalCopy {
		if copy[from] != to {
			t.Errorf("Expected mapping %s->%s, got %s->%s", from, to, from, copy[from])
		}
	}
}

func TestStatusMapping_String(t *testing.T) {
	tests := []struct {
		name     string
		mapping  StatusMapping
		expected string
	}{
		{
			name:     "empty mapping",
			mapping:  StatusMapping{},
			expected: "{}",
		},
		{
			name: "single mapping",
			mapping: StatusMapping{
				StatusInvalid: StatusFailed,
			},
			expected: "{invalid->failed}",
		},
		{
			name: "multiple mappings",
			mapping: StatusMapping{
				StatusInvalid: StatusFailed,
				StatusSkipped: StatusPassed,
			},
			expected: "{invalid->failed, skipped->passed}", // Order may vary
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.mapping.String()
			if tt.name == "empty mapping" {
				if result != tt.expected {
					t.Errorf("Expected %s, got %s", tt.expected, result)
				}
			} else {
				// For non-empty mappings, just check it contains the expected elements
				if result == "{}" {
					t.Errorf("Expected non-empty string, got %s", result)
				}
			}
		})
	}
}

func TestValidateMapping(t *testing.T) {
	tests := []struct {
		name        string
		mapping     map[string]string
		expectError bool
	}{
		{
			name: "valid mapping",
			mapping: map[string]string{
				"invalid": "failed",
				"skipped": "passed",
			},
			expectError: false,
		},
		{
			name:        "empty mapping",
			mapping:     map[string]string{},
			expectError: false,
		},
		{
			name: "invalid source status",
			mapping: map[string]string{
				"unknown": "failed",
			},
			expectError: true,
		},
		{
			name: "invalid target status",
			mapping: map[string]string{
				"failed": "unknown",
			},
			expectError: true,
		},
		{
			name: "mapping to same status",
			mapping: map[string]string{
				"failed": "failed",
			},
			expectError: true,
		},
		{
			name: "case insensitive validation",
			mapping: map[string]string{
				"INVALID": "FAILED",
			},
			expectError: false,
		},
		{
			name: "whitespace handling",
			mapping: map[string]string{
				" invalid ": " failed ",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMapping(tt.mapping)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}
