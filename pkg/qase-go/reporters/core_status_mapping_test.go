package reporters

import (
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

func TestCoreReporter_StatusMapping(t *testing.T) {
	tests := []struct {
		name           string
		config         *config.Config
		inputStatus    domain.TestResultStatus
		expectedStatus domain.TestResultStatus
		expectMapping  bool
	}{
		{
			name: "no status mapping configured",
			config: &config.Config{
				Mode:     "report",
				Fallback: "off",
				Report: config.ReportConfig{
					Driver: "local",
					Connection: config.ConnectionConfig{
						Local: config.LocalConfig{
							Path:   "./test-reports",
							Format: "json",
						},
					},
				},
			},
			inputStatus:    domain.StatusInvalid,
			expectedStatus: domain.StatusInvalid,
			expectMapping:  false,
		},
		{
			name: "status mapping - invalid to failed",
			config: &config.Config{
				Mode:     "report",
				Fallback: "off",
				StatusMapping: map[string]string{
					"invalid": "failed",
				},
				Report: config.ReportConfig{
					Driver: "local",
					Connection: config.ConnectionConfig{
						Local: config.LocalConfig{
							Path:   "./test-reports",
							Format: "json",
						},
					},
				},
			},
			inputStatus:    domain.StatusInvalid,
			expectedStatus: domain.StatusFailed,
			expectMapping:  true,
		},
		{
			name: "status mapping - skipped to passed",
			config: &config.Config{
				Mode:     "report",
				Fallback: "off",
				StatusMapping: map[string]string{
					"skipped": "passed",
				},
				Report: config.ReportConfig{
					Driver: "local",
					Connection: config.ConnectionConfig{
						Local: config.LocalConfig{
							Path:   "./test-reports",
							Format: "json",
						},
					},
				},
			},
			inputStatus:    domain.StatusSkipped,
			expectedStatus: domain.StatusPassed,
			expectMapping:  true,
		},
		{
			name: "status mapping - unmapped status",
			config: &config.Config{
				Mode:     "report",
				Fallback: "off",
				StatusMapping: map[string]string{
					"invalid": "failed",
				},
				Report: config.ReportConfig{
					Driver: "local",
					Connection: config.ConnectionConfig{
						Local: config.LocalConfig{
							Path:   "./test-reports",
							Format: "json",
						},
					},
				},
			},
			inputStatus:    domain.StatusPassed,
			expectedStatus: domain.StatusPassed,
			expectMapping:  false,
		},
		{
			name: "multiple status mappings",
			config: &config.Config{
				Mode:     "report",
				Fallback: "off",
				StatusMapping: map[string]string{
					"invalid": "failed",
					"skipped": "passed",
				},
				Report: config.ReportConfig{
					Driver: "local",
					Connection: config.ConnectionConfig{
						Local: config.LocalConfig{
							Path:   "./test-reports",
							Format: "json",
						},
					},
				},
			},
			inputStatus:    domain.StatusSkipped,
			expectedStatus: domain.StatusPassed,
			expectMapping:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create core reporter
			reporter, err := NewCoreReporter(tt.config)
			if err != nil {
				t.Fatalf("Failed to create core reporter: %v", err)
			}

			// Create test result
			result := &domain.TestResult{
				Title: "Test Case",
				Execution: domain.TestExecution{
					Status: tt.inputStatus,
				},
			}

			// Add result to reporter
			err = reporter.AddResult(result)
			if err != nil {
				t.Fatalf("Failed to add result: %v", err)
			}

			// Verify status mapping
			if result.Execution.Status != tt.expectedStatus {
				t.Errorf("Expected status %s, got %s", tt.expectedStatus, result.Execution.Status)
			}

			// Verify mapping was applied
			statusMapping := reporter.GetStatusMapping()
			if tt.expectMapping && statusMapping.IsEmpty() {
				t.Error("Expected status mapping to be configured")
			}
		})
	}
}

func TestCoreReporter_StatusMapping_InvalidConfig(t *testing.T) {
	tests := []struct {
		name        string
		config      *config.Config
		expectError bool
		errorMsg    string
	}{
		{
			name: "invalid source status in mapping",
			config: &config.Config{
				Mode:     "report",
				Fallback: "off",
				StatusMapping: map[string]string{
					"unknown": "failed",
				},
				Report: config.ReportConfig{
					Driver: "local",
					Connection: config.ConnectionConfig{
						Local: config.LocalConfig{
							Path:   "./test-reports",
							Format: "json",
						},
					},
				},
			},
			expectError: true,
			errorMsg:    "invalid status mapping configuration",
		},
		{
			name: "invalid target status in mapping",
			config: &config.Config{
				Mode:     "report",
				Fallback: "off",
				StatusMapping: map[string]string{
					"failed": "unknown",
				},
				Report: config.ReportConfig{
					Driver: "local",
					Connection: config.ConnectionConfig{
						Local: config.LocalConfig{
							Path:   "./test-reports",
							Format: "json",
						},
					},
				},
			},
			expectError: true,
			errorMsg:    "invalid status mapping configuration",
		},
		{
			name: "mapping to same status",
			config: &config.Config{
				Mode:     "report",
				Fallback: "off",
				StatusMapping: map[string]string{
					"failed": "failed",
				},
				Report: config.ReportConfig{
					Driver: "local",
					Connection: config.ConnectionConfig{
						Local: config.LocalConfig{
							Path:   "./test-reports",
							Format: "json",
						},
					},
				},
			},
			expectError: true,
			errorMsg:    "invalid status mapping configuration",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create core reporter
			_, err := NewCoreReporter(tt.config)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if tt.errorMsg != "" && !contains(err.Error(), tt.errorMsg) {
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

func TestCoreReporter_GetStatusMapping(t *testing.T) {
	tests := []struct {
		name           string
		config         *config.Config
		expectedEmpty  bool
		expectedMappings map[domain.TestResultStatus]domain.TestResultStatus
	}{
		{
			name: "no status mapping",
			config: &config.Config{
				Mode:     "report",
				Fallback: "off",
				Report: config.ReportConfig{
					Driver: "local",
					Connection: config.ConnectionConfig{
						Local: config.LocalConfig{
							Path:   "./test-reports",
							Format: "json",
						},
					},
				},
			},
			expectedEmpty: true,
		},
		{
			name: "status mapping configured",
			config: &config.Config{
				Mode:     "report",
				Fallback: "off",
				StatusMapping: map[string]string{
					"invalid": "failed",
					"skipped": "passed",
				},
				Report: config.ReportConfig{
					Driver: "local",
					Connection: config.ConnectionConfig{
						Local: config.LocalConfig{
							Path:   "./test-reports",
							Format: "json",
						},
					},
				},
			},
			expectedEmpty: false,
			expectedMappings: map[domain.TestResultStatus]domain.TestResultStatus{
				domain.StatusInvalid: domain.StatusFailed,
				domain.StatusSkipped: domain.StatusPassed,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create core reporter
			reporter, err := NewCoreReporter(tt.config)
			if err != nil {
				t.Fatalf("Failed to create core reporter: %v", err)
			}

			// Get status mapping
			mapping := reporter.GetStatusMapping()

			// Verify empty state
			if mapping.IsEmpty() != tt.expectedEmpty {
				t.Errorf("Expected empty %v, got %v", tt.expectedEmpty, mapping.IsEmpty())
			}

			// Verify mappings if not empty
			if !tt.expectedEmpty && tt.expectedMappings != nil {
				actualMappings := mapping.GetMappings()
				if len(actualMappings) != len(tt.expectedMappings) {
					t.Errorf("Expected %d mappings, got %d", len(tt.expectedMappings), len(actualMappings))
				}

				for from, to := range tt.expectedMappings {
					if actualMappings[from] != to {
						t.Errorf("Expected mapping %s->%s, got %s->%s", from, to, from, actualMappings[from])
					}
				}
			}
		})
	}
}

// helper function to check if string contains substring
func contains(s, substr string) bool {
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
