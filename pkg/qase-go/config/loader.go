package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	// DefaultConfigFileName is the default configuration file name
	DefaultConfigFileName = "qase.config.json"
)

// ConfigLoader provides methods to load configuration from multiple sources
type ConfigLoader struct {
	searchPaths []string
	fileName    string
}

// NewConfigLoader creates a new configuration loader
func NewConfigLoader() *ConfigLoader {
	return &ConfigLoader{
		searchPaths: []string{
			".", // current directory
			"./config",
			"./configs",
			"./build",
			"./test",
			"./tests",
		},
		fileName: DefaultConfigFileName,
	}
}

// WithSearchPaths sets custom search paths for configuration file
func (l *ConfigLoader) WithSearchPaths(paths []string) *ConfigLoader {
	l.searchPaths = paths
	return l
}

// WithFileName sets custom configuration file name
func (l *ConfigLoader) WithFileName(fileName string) *ConfigLoader {
	l.fileName = fileName
	return l
}

// AddSearchPath adds a search path for configuration file
func (l *ConfigLoader) AddSearchPath(path string) *ConfigLoader {
	l.searchPaths = append(l.searchPaths, path)
	return l
}

// Load loads configuration from multiple sources with the following priority:
// 1. Environment variables (highest priority)
// 2. Configuration file
// 3. Default values (lowest priority)
func (l *ConfigLoader) Load() (*Config, error) {
	// Start with default configuration
	config := NewConfig()

	// Try to load from configuration file
	configPath := l.findConfigFile()
	if configPath != "" {
		fileConfig, err := LoadFromFile(configPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load config from file %s: %w", configPath, err)
		}
		// Merge file configuration over default
		config = fileConfig
	}

	// Load environment variables (highest priority)
	config.LoadFromEnvironment()

	// Validate final configuration
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return config, nil
}

// LoadUnsafe loads configuration without validation
func (l *ConfigLoader) LoadUnsafe() *Config {
	config := NewConfig()

	// Try to load from configuration file
	configPath := l.findConfigFile()
	if configPath != "" {
		if fileConfig, err := LoadFromFile(configPath); err == nil {
			config = fileConfig
		}
	}

	// Load environment variables
	config.LoadFromEnvironment()

	return config
}

// LoadWithDefaults loads configuration with custom default values
func (l *ConfigLoader) LoadWithDefaults(defaults *Config) (*Config, error) {
	// Start with provided defaults
	config := defaults

	// Try to load from configuration file
	configPath := l.findConfigFile()
	if configPath != "" {
		fileConfig, err := LoadFromFile(configPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load config from file %s: %w", configPath, err)
		}
		// Merge file configuration over defaults
		config = fileConfig
	}

	// Load environment variables (highest priority)
	config.LoadFromEnvironment()

	// Validate final configuration
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return config, nil
}

// FindConfigFile returns the path to configuration file if found
func (l *ConfigLoader) FindConfigFile() string {
	return l.findConfigFile()
}

// findConfigFile searches for configuration file in search paths
func (l *ConfigLoader) findConfigFile() string {
	for _, path := range l.searchPaths {
		configPath := filepath.Join(path, l.fileName)
		if _, err := os.Stat(configPath); err == nil {
			return configPath
		}
	}
	return ""
}

// CreateDefaultConfigFile creates a default configuration file
func (l *ConfigLoader) CreateDefaultConfigFile(path string) error {
	config := NewConfig()
	
	// Set some example values
	config.Mode = "testops"
	config.Fallback = "report"
	config.Debug = false
	config.Environment = "local"
	config.CaptureLogs = false
	
	// Example TestOps configuration
	config.TestOps.API.Token = "<token>"
	config.TestOps.API.Host = "qase.io"
	config.TestOps.Project = "<project_code>"
	config.TestOps.Run.Title = "Regress run"
	config.TestOps.Run.Description = "Regress run description"
	config.TestOps.Run.Complete = true
	config.TestOps.Run.Tags = []string{"tag1", "tag2"}
	config.TestOps.Run.Configurations.Values = []ConfigurationValue{
		{Name: "browser", Value: "chrome"},
		{Name: "environment", Value: "staging"},
	}
	config.TestOps.Run.Configurations.CreateIfNotExists = true
	config.TestOps.Defect = false
	config.TestOps.Batch.Size = 100

	return config.SaveToFile(path)
}

// Global convenience functions

// Load loads configuration using default loader
func Load() (*Config, error) {
	return NewConfigLoader().Load()
}

// LoadUnsafe loads configuration using default loader without validation
func LoadUnsafe() *Config {
	return NewConfigLoader().LoadUnsafe()
}

// LoadFrom loads configuration from specific file
func LoadFrom(filePath string) (*Config, error) {
	return NewConfigLoader().WithSearchPaths([]string{filepath.Dir(filePath)}).
		WithFileName(filepath.Base(filePath)).Load()
}

// CreateDefaultConfig creates a default configuration file in current directory
func CreateDefaultConfig() error {
	return NewConfigLoader().CreateDefaultConfigFile(DefaultConfigFileName)
}

// CreateDefaultConfigAt creates a default configuration file at specified path
func CreateDefaultConfigAt(path string) error {
	return NewConfigLoader().CreateDefaultConfigFile(path)
}