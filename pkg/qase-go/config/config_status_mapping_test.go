package config

import (
	"os"
	"testing"
)

// createValidConfig creates a valid configuration with the given status mapping
func createValidConfig(statusMapping map[string]string) *Config {
	return &Config{
		Mode:         "report",
		Fallback:     "off",
		StatusMapping: statusMapping,
		Report: ReportConfig{
			Driver: "local",
			Connection: ConnectionConfig{
				Local: LocalConfig{
					Path:   "./test-reports",
					Format: "json",
				},
			},
		},
	}
}

func TestConfig_StatusMapping(t *testing.T) {
	tests := []struct {
		name           string
		config         *Config
		expectedMapping map[string]string
	}{
		{
			name: "empty status mapping",
			config: &Config{
				StatusMapping: nil,
			},
			expectedMapping: map[string]string{},
		},
		{
			name: "valid status mapping",
			config: &Config{
				StatusMapping: map[string]string{
					"invalid": "failed",
					"skipped": "passed",
				},
			},
			expectedMapping: map[string]string{
				"invalid": "failed",
				"skipped": "passed",
			},
		},
		{
			name: "single status mapping",
			config: &Config{
				StatusMapping: map[string]string{
					"invalid": "failed",
				},
			},
			expectedMapping: map[string]string{
				"invalid": "failed",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapping := tt.config.GetStatusMapping()
			
			if len(mapping) != len(tt.expectedMapping) {
				t.Errorf("Expected mapping length %d, got %d", len(tt.expectedMapping), len(mapping))
			}
			
			for from, to := range tt.expectedMapping {
				if mapping[from] != to {
					t.Errorf("Expected mapping %s->%s, got %s->%s", from, to, from, mapping[from])
				}
			}
		})
	}
}

func TestConfig_LoadFromEnvironment_StatusMapping(t *testing.T) {
	tests := []struct {
		name           string
		envValue       string
		expectedMapping map[string]string
		expectError    bool
	}{
		{
			name:           "valid env format",
			envValue:       "invalid=failed,skipped=passed",
			expectedMapping: map[string]string{
				"invalid": "failed",
				"skipped": "passed",
			},
			expectError: false,
		},
		{
			name:           "single mapping",
			envValue:       "invalid=failed",
			expectedMapping: map[string]string{
				"invalid": "failed",
			},
			expectError: false,
		},
		{
			name:           "empty env value",
			envValue:       "",
			expectedMapping: map[string]string{},
			expectError:    false,
		},
		{
			name:           "whitespace handling",
			envValue:       " invalid = failed , skipped = passed ",
			expectedMapping: map[string]string{
				"invalid": "failed",
				"skipped": "passed",
			},
			expectError: false,
		},
		{
			name:           "empty pairs",
			envValue:       "invalid=failed,,skipped=passed,",
			expectedMapping: map[string]string{
				"invalid": "failed",
				"skipped": "passed",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variable
			if tt.envValue != "" {
				os.Setenv("QASE_STATUS_MAPPING", tt.envValue)
				defer os.Unsetenv("QASE_STATUS_MAPPING")
			}

			config := NewConfig()
			config.LoadFromEnvironment()

			mapping := config.GetStatusMapping()

			if len(mapping) != len(tt.expectedMapping) {
				t.Errorf("Expected mapping length %d, got %d", len(tt.expectedMapping), len(mapping))
			}

			for from, to := range tt.expectedMapping {
				if mapping[from] != to {
					t.Errorf("Expected mapping %s->%s, got %s->%s", from, to, from, mapping[from])
				}
			}
		})
	}
}

func TestConfig_Validate_StatusMapping(t *testing.T) {
	tests := []struct {
		name        string
		config      *Config
		expectError bool
		errorMsg    string
	}{
		{
			name:        "valid status mapping",
			config:      createValidConfig(map[string]string{"invalid": "failed", "skipped": "passed"}),
			expectError: false,
		},
		{
			name:        "empty status mapping",
			config:      createValidConfig(map[string]string{}),
			expectError: false,
		},
		{
			name:        "nil status mapping",
			config:      createValidConfig(nil),
			expectError: false,
		},
		{
			name:        "invalid source status",
			config:      createValidConfig(map[string]string{"unknown": "failed"}),
			expectError: true,
			errorMsg:    "invalid source status",
		},
		{
			name:        "invalid target status",
			config:      createValidConfig(map[string]string{"failed": "unknown"}),
			expectError: true,
			errorMsg:    "invalid target status",
		},
		{
			name:        "mapping to same status",
			config:      createValidConfig(map[string]string{"failed": "failed"}),
			expectError: true,
			errorMsg:    "cannot map status",
		},
		{
			name:        "case insensitive validation",
			config:      createValidConfig(map[string]string{"INVALID": "FAILED"}),
			expectError: false,
		},
		{
			name:        "whitespace handling",
			config:      createValidConfig(map[string]string{" invalid ": " failed "}),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if tt.errorMsg != "" && !containsString(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error message to contain '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

// helper function to check if string contains substring
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || 
		(len(s) > len(substr) && (s[:len(substr)] == substr || 
		s[len(s)-len(substr):] == substr || 
		containsSubstring(s, substr))))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
