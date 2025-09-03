package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig()

	// Test default values
	if config.Mode != "off" {
		t.Errorf("Expected default mode 'off', got '%s'", config.Mode)
	}
	if config.Fallback != "off" {
		t.Errorf("Expected default fallback 'off', got '%s'", config.Fallback)
	}
	if config.Debug != false {
		t.Errorf("Expected default debug false, got %t", config.Debug)
	}
	if config.Environment != "local" {
		t.Errorf("Expected default environment 'local', got '%s'", config.Environment)
	}
	if config.CaptureLogs != false {
		t.Errorf("Expected default captureLogs false, got %t", config.CaptureLogs)
	}
	if config.Report.Driver != "local" {
		t.Errorf("Expected default report driver 'local', got '%s'", config.Report.Driver)
	}
	if config.Report.Connection.Local.Path != "./build/qase-report" {
		t.Errorf("Expected default report path './build/qase-report', got '%s'", config.Report.Connection.Local.Path)
	}
	if config.Report.Connection.Local.Format != "json" {
		t.Errorf("Expected default report format 'json', got '%s'", config.Report.Connection.Local.Format)
	}
	if config.TestOps.API.Host != "qase.io" {
		t.Errorf("Expected default host 'qase.io', got '%s'", config.TestOps.API.Host)
	}
	if config.TestOps.Batch.Size != 100 {
		t.Errorf("Expected default batch size 100, got %d", config.TestOps.Batch.Size)
	}
}

func TestConfigValidation(t *testing.T) {
	tests := []struct {
		name        string
		config      *Config
		expectError bool
	}{
		{
			name:        "valid default config",
			config:      NewConfig(),
			expectError: false,
		},
		{
			name: "invalid mode",
			config: func() *Config {
				c := NewConfig()
				c.Mode = "invalid"
				return c
			}(),
			expectError: true,
		},
		{
			name: "invalid fallback",
			config: func() *Config {
				c := NewConfig()
				c.Fallback = "invalid"
				return c
			}(),
			expectError: true,
		},
		{
			name: "testops mode without token",
			config: func() *Config {
				c := NewConfig()
				c.Mode = "testops"
				return c
			}(),
			expectError: true,
		},
		{
			name: "testops mode without project",
			config: func() *Config {
				c := NewConfig()
				c.Mode = "testops"
				c.TestOps.API.Token = "token"
				return c
			}(),
			expectError: true,
		},
		{
			name: "testops mode without run ID",
			config: func() *Config {
				c := NewConfig()
				c.Mode = "testops"
				c.TestOps.API.Token = "token"
				c.TestOps.Project = "DEMO"
				return c
			}(),
			expectError: true,
		},
		{
			name: "valid testops config",
			config: func() *Config {
				c := NewConfig()
				c.Mode = "testops"
				c.TestOps.API.Token = "token"
				c.TestOps.Project = "DEMO"
				runID := int64(123)
				c.TestOps.Run.ID = &runID
				return c
			}(),
			expectError: false,
		},
		{
			name: "invalid report format",
			config: func() *Config {
				c := NewConfig()
				c.Report.Connection.Local.Format = "xml"
				return c
			}(),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.expectError && err == nil {
				t.Error("Expected validation error, but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Expected no validation error, but got: %v", err)
			}
		})
	}
}

func TestLoadFromEnvironment(t *testing.T) {
	// Set environment variables
	os.Setenv("QASE_MODE", "testops")
	os.Setenv("QASE_DEBUG", "true")
	os.Setenv("QASE_ENVIRONMENT", "staging")
	os.Setenv("QASE_CAPTURE_LOGS", "true")
	os.Setenv("QASE_REPORT_DRIVER", "local")
	os.Setenv("QASE_REPORT_CONNECTION_PATH", "/custom/path")
	os.Setenv("QASE_REPORT_CONNECTION_FORMAT", "yaml")
	os.Setenv("QASE_TESTOPS_API_TOKEN", "test-token")
	os.Setenv("QASE_TESTOPS_API_HOST", "custom.qase.io")
	os.Setenv("QASE_TESTOPS_PROJECT", "TEST")
	os.Setenv("QASE_TESTOPS_RUN_ID", "123")
	os.Setenv("QASE_TESTOPS_DEFECT", "true")
	os.Setenv("QASE_TESTOPS_BATCH_SIZE", "50")

	defer func() {
		// Clean up environment variables
		os.Unsetenv("QASE_MODE")
		os.Unsetenv("QASE_DEBUG")
		os.Unsetenv("QASE_ENVIRONMENT")
		os.Unsetenv("QASE_CAPTURE_LOGS")
		os.Unsetenv("QASE_REPORT_DRIVER")
		os.Unsetenv("QASE_REPORT_CONNECTION_PATH")
		os.Unsetenv("QASE_REPORT_CONNECTION_FORMAT")
		os.Unsetenv("QASE_TESTOPS_API_TOKEN")
		os.Unsetenv("QASE_TESTOPS_API_HOST")
		os.Unsetenv("QASE_TESTOPS_PROJECT")
		os.Unsetenv("QASE_TESTOPS_RUN_ID")
		os.Unsetenv("QASE_TESTOPS_DEFECT")
		os.Unsetenv("QASE_TESTOPS_BATCH_SIZE")
	}()

	config := NewConfig()
	config.LoadFromEnvironment()

	if config.Mode != "testops" {
		t.Errorf("Expected mode 'testops', got '%s'", config.Mode)
	}
	if !config.Debug {
		t.Error("Expected debug to be true")
	}
	if config.Environment != "staging" {
		t.Errorf("Expected environment 'staging', got '%s'", config.Environment)
	}
	if !config.CaptureLogs {
		t.Error("Expected captureLogs to be true")
	}
	if config.Report.Driver != "local" {
		t.Errorf("Expected report driver 'local', got '%s'", config.Report.Driver)
	}
	if config.Report.Connection.Local.Path != "/custom/path" {
		t.Errorf("Expected report path '/custom/path', got '%s'", config.Report.Connection.Local.Path)
	}
	if config.Report.Connection.Local.Format != "yaml" {
		t.Errorf("Expected report format 'yaml', got '%s'", config.Report.Connection.Local.Format)
	}
	if config.TestOps.API.Token != "test-token" {
		t.Errorf("Expected token 'test-token', got '%s'", config.TestOps.API.Token)
	}
	if config.TestOps.API.Host != "custom.qase.io" {
		t.Errorf("Expected host 'custom.qase.io', got '%s'", config.TestOps.API.Host)
	}
	if config.TestOps.Project != "TEST" {
		t.Errorf("Expected project 'TEST', got '%s'", config.TestOps.Project)
	}
	if config.TestOps.Run.ID == nil || *config.TestOps.Run.ID != 123 {
		t.Errorf("Expected run ID 123, got %v", config.TestOps.Run.ID)
	}
	if !config.TestOps.Defect {
		t.Error("Expected defect to be true")
	}
	if config.TestOps.Batch.Size != 50 {
		t.Errorf("Expected batch size 50, got %d", config.TestOps.Batch.Size)
	}
}

func TestLoadFromFile(t *testing.T) {
	// Create temporary config file
	tempDir := t.TempDir()
	configFile := filepath.Join(tempDir, "test-config.json")

	testConfig := NewConfig()
	testConfig.Mode = "testops"
	testConfig.Debug = true
	testConfig.TestOps.API.Token = "file-token"
	testConfig.TestOps.Project = "FILE"

	err := testConfig.SaveToFile(configFile)
	if err != nil {
		t.Fatalf("Failed to save test config: %v", err)
	}

	// Load config from file
	loadedConfig, err := LoadFromFile(configFile)
	if err != nil {
		t.Fatalf("Failed to load config from file: %v", err)
	}

	if loadedConfig.Mode != "testops" {
		t.Errorf("Expected mode 'testops', got '%s'", loadedConfig.Mode)
	}
	if !loadedConfig.Debug {
		t.Error("Expected debug to be true")
	}
	if loadedConfig.TestOps.API.Token != "file-token" {
		t.Errorf("Expected token 'file-token', got '%s'", loadedConfig.TestOps.API.Token)
	}
	if loadedConfig.TestOps.Project != "FILE" {
		t.Errorf("Expected project 'FILE', got '%s'", loadedConfig.TestOps.Project)
	}
}

func TestConfigBuilder(t *testing.T) {
	config, err := NewConfigBuilder().
		TestOpsMode().
		WithAPIToken("builder-token").
		WithProject("BUILD").
		WithDebug(true).
		WithEnvironment("test").
		WithCaptureLogs(true).
		WithReportDriver("local").
		WithReportPath("/custom/report").
		WithReportFormat("yaml").
		WithRunID(123).
		WithDefect(true).
		WithBatchSize(75).
		Build()

	if err != nil {
		t.Fatalf("Failed to build config: %v", err)
	}

	if config.Mode != "testops" {
		t.Errorf("Expected mode 'testops', got '%s'", config.Mode)
	}
	if config.TestOps.API.Token != "builder-token" {
		t.Errorf("Expected token 'builder-token', got '%s'", config.TestOps.API.Token)
	}
	if config.TestOps.Project != "BUILD" {
		t.Errorf("Expected project 'BUILD', got '%s'", config.TestOps.Project)
	}
	if !config.Debug {
		t.Error("Expected debug to be true")
	}
	if config.Environment != "test" {
		t.Errorf("Expected environment 'test', got '%s'", config.Environment)
	}
	if !config.CaptureLogs {
		t.Error("Expected captureLogs to be true")
	}
	if config.Report.Driver != "local" {
		t.Errorf("Expected report driver 'local', got '%s'", config.Report.Driver)
	}
	if config.Report.Connection.Local.Path != "/custom/report" {
		t.Errorf("Expected report path '/custom/report', got '%s'", config.Report.Connection.Local.Path)
	}
	if config.Report.Connection.Local.Format != "yaml" {
		t.Errorf("Expected report format 'yaml', got '%s'", config.Report.Connection.Local.Format)
	}
	if !config.TestOps.Defect {
		t.Error("Expected defect to be true")
	}
	if config.TestOps.Batch.Size != 75 {
		t.Errorf("Expected batch size 75, got %d", config.TestOps.Batch.Size)
	}
}

func TestConfigLoader(t *testing.T) {
	// Create temporary directory and config file
	tempDir := t.TempDir()
	configFile := filepath.Join(tempDir, DefaultConfigFileName)

	// Create test config file
	testConfig := NewConfig()
	testConfig.Mode = "testops"
	testConfig.TestOps.API.Token = "loader-token"
	testConfig.TestOps.Project = "LOAD"
	runID := int64(999)
	testConfig.TestOps.Run.ID = &runID

	err := testConfig.SaveToFile(configFile)
	if err != nil {
		t.Fatalf("Failed to save test config: %v", err)
	}

	// Set environment variable to override file value
	os.Setenv("QASE_TESTOPS_API_TOKEN", "env-token")
	defer os.Unsetenv("QASE_TESTOPS_API_TOKEN")

	// Load config with custom search path
	loader := NewConfigLoader().WithSearchPaths([]string{tempDir})
	config, err := loader.Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Environment variable should override file value
	if config.TestOps.API.Token != "env-token" {
		t.Errorf("Expected token 'env-token', got '%s'", config.TestOps.API.Token)
	}
	// File value should be used for project
	if config.TestOps.Project != "LOAD" {
		t.Errorf("Expected project 'LOAD', got '%s'", config.TestOps.Project)
	}
}

func TestJSONMarshaling(t *testing.T) {
	config := NewConfig()
	config.Mode = "testops"
	config.TestOps.API.Token = "test-token"
	config.TestOps.Project = "JSON"

	// Marshal to JSON
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal config to JSON: %v", err)
	}

	// Unmarshal from JSON
	var loadedConfig Config
	err = json.Unmarshal(data, &loadedConfig)
	if err != nil {
		t.Fatalf("Failed to unmarshal config from JSON: %v", err)
	}

	// Verify values
	if loadedConfig.Mode != config.Mode {
		t.Errorf("Mode mismatch: expected '%s', got '%s'", config.Mode, loadedConfig.Mode)
	}
	if loadedConfig.TestOps.API.Token != config.TestOps.API.Token {
		t.Errorf("Token mismatch: expected '%s', got '%s'", config.TestOps.API.Token, loadedConfig.TestOps.API.Token)
	}
}
