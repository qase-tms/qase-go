package config

import "testing"

func TestConfig_GetBatchSize(t *testing.T) {
	tests := []struct {
		name       string
		configured int
		expected   int
	}{
		{"zero falls back to default", 0, DefaultBatchSize},
		{"negative falls back to default", -10, DefaultBatchSize},
		{"small value is kept as is", 1, 1},
		{"value below the cap is kept as is", 150, 150},
		{"value equal to the cap is kept as is", MaxBatchSize, MaxBatchSize},
		{"value above the cap is clamped", 2000, MaxBatchSize},
		{"one above the cap is clamped", MaxBatchSize + 1, MaxBatchSize},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := NewConfig()
			cfg.TestOps.Batch.Size = tt.configured

			if got := cfg.GetBatchSize(); got != tt.expected {
				t.Errorf("GetBatchSize() with configured size %d = %d, want %d", tt.configured, got, tt.expected)
			}
		})
	}
}

func TestMaxBatchSizeMatchesAPILimit(t *testing.T) {
	if MaxBatchSize != 200 {
		t.Errorf("MaxBatchSize = %d, want 200 (the bulk results API limit)", MaxBatchSize)
	}
}

func TestDefaultConfigBatchSizeIsWithinLimit(t *testing.T) {
	cfg := NewConfig()

	if cfg.TestOps.Batch.Size > MaxBatchSize {
		t.Errorf("default batch size %d exceeds MaxBatchSize %d", cfg.TestOps.Batch.Size, MaxBatchSize)
	}
}
