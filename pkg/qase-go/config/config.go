package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config represents the main configuration structure
type Config struct {
	Mode        string        `json:"mode"`
	Fallback    string        `json:"fallback"`
	Debug       bool          `json:"debug"`
	Environment string        `json:"environment"`
	CaptureLogs bool          `json:"captureLogs"`
	Report      ReportConfig  `json:"report"`
	TestOps     TestOpsConfig `json:"testops"`
}

// ReportConfig represents report configuration
type ReportConfig struct {
	Driver     string           `json:"driver"`
	Connection ConnectionConfig `json:"connection"`
}

// ConnectionConfig represents connection configuration for report
type ConnectionConfig struct {
	Local LocalConfig `json:"local"`
}

// LocalConfig represents local file system configuration
type LocalConfig struct {
	Path   string `json:"path"`
	Format string `json:"format"`
}

// TestOpsConfig represents TestOps configuration
type TestOpsConfig struct {
	API     APIConfig   `json:"api"`
	Run     RunConfig   `json:"run"`
	Defect  bool        `json:"defect"`
	Project string      `json:"project"`
	Batch   BatchConfig `json:"batch"`
}

// APIConfig represents API configuration
type APIConfig struct {
	Token string `json:"token"`
	Host  string `json:"host"`
}

// RunConfig represents test run configuration
type RunConfig struct {
	ID *int64 `json:"id,omitempty"` // Test run ID
}

// BatchConfig represents batch configuration
type BatchConfig struct {
	Size int `json:"size"`
}

// LoggingConfig is no longer needed - logging is always enabled

// NewConfig creates a new configuration with default values
func NewConfig() *Config {
	return &Config{
		Mode:        "off",
		Fallback:    "off",
		Debug:       false,
		Environment: "local",
		CaptureLogs: false,
		// Logging is always enabled - no configuration needed
		Report: ReportConfig{
			Driver: "local",
			Connection: ConnectionConfig{
				Local: LocalConfig{
					Path:   "./build/qase-report",
					Format: "json",
				},
			},
		},
		TestOps: TestOpsConfig{
			API: APIConfig{
				Token: "",
				Host:  "qase.io",
			},
			Run:     RunConfig{},
			Defect:  false,
			Project: "",
			Batch: BatchConfig{
				Size: 100,
			},
		},
	}
}

// LoadFromEnvironment loads configuration from environment variables
// Using the same environment variable names as qase-java-commons
func (c *Config) LoadFromEnvironment() {
	// Main configuration
	if mode := os.Getenv("QASE_MODE"); mode != "" {
		c.Mode = mode
	}
	if fallback := os.Getenv("QASE_FALLBACK"); fallback != "" {
		c.Fallback = fallback
	}
	if debug := os.Getenv("QASE_DEBUG"); debug != "" {
		c.Debug = strings.ToLower(debug) == "true"
	}
	if env := os.Getenv("QASE_ENVIRONMENT"); env != "" {
		c.Environment = env
	}
	if captureLogs := os.Getenv("QASE_CAPTURE_LOGS"); captureLogs != "" {
		c.CaptureLogs = strings.ToLower(captureLogs) == "true"
	}

	// Logging configuration - always enabled to both console and file

	// Report configuration
	if driver := os.Getenv("QASE_REPORT_DRIVER"); driver != "" {
		c.Report.Driver = driver
	}
	if path := os.Getenv("QASE_REPORT_CONNECTION_PATH"); path != "" {
		c.Report.Connection.Local.Path = path
	}
	if format := os.Getenv("QASE_REPORT_CONNECTION_FORMAT"); format != "" {
		c.Report.Connection.Local.Format = format
	}

	// TestOps API configuration
	if token := os.Getenv("QASE_TESTOPS_API_TOKEN"); token != "" {
		c.TestOps.API.Token = token
	}
	if host := os.Getenv("QASE_TESTOPS_API_HOST"); host != "" {
		c.TestOps.API.Host = host
	}

	// TestOps Run configuration
	if runID := os.Getenv("QASE_TESTOPS_RUN_ID"); runID != "" {
		if id, err := strconv.ParseInt(runID, 10, 64); err == nil {
			c.TestOps.Run.ID = &id
		}
	}

	// TestOps other configuration
	if defect := os.Getenv("QASE_TESTOPS_DEFECT"); defect != "" {
		c.TestOps.Defect = strings.ToLower(defect) == "true"
	}
	if project := os.Getenv("QASE_TESTOPS_PROJECT"); project != "" {
		c.TestOps.Project = project
	}
	if batchSize := os.Getenv("QASE_TESTOPS_BATCH_SIZE"); batchSize != "" {
		if size, err := strconv.Atoi(batchSize); err == nil {
			c.TestOps.Batch.Size = size
		}
	}
}

// LoadFromFile loads configuration from JSON file
func LoadFromFile(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	config := NewConfig()
	if err := json.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return config, nil
}

// SaveToFile saves configuration to JSON file
func (c *Config) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// Validate validates the configuration
func (c *Config) Validate() error {
	// Validate mode
	validModes := []string{"testops", "report", "off"}
	if !contains(validModes, c.Mode) {
		return fmt.Errorf("invalid mode '%s', must be one of: %s", c.Mode, strings.Join(validModes, ", "))
	}

	// Validate fallback
	validFallbacks := []string{"report", "off"}
	if !contains(validFallbacks, c.Fallback) {
		return fmt.Errorf("invalid fallback '%s', must be one of: %s", c.Fallback, strings.Join(validFallbacks, ", "))
	}

	// Validate report format
	validFormats := []string{"json", "yaml"}
	if !contains(validFormats, c.Report.Connection.Local.Format) {
		return fmt.Errorf("invalid report format '%s', must be one of: %s", c.Report.Connection.Local.Format, strings.Join(validFormats, ", "))
	}

	// Validate TestOps configuration if mode is testops
	if c.Mode == "testops" {
		if c.TestOps.API.Token == "" {
			return fmt.Errorf("API token is required when mode is 'testops'")
		}
		if c.TestOps.Project == "" {
			return fmt.Errorf("project code is required when mode is 'testops'")
		}
		if c.TestOps.Run.ID == nil || *c.TestOps.Run.ID == 0 {
			return fmt.Errorf("run ID is required when mode is 'testops'")
		}
	}

	return nil
}

// helper function to check if slice contains a value
func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
