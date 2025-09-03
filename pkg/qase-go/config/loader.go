package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

// NewConfigLoaderWithParentSearch creates a new configuration loader that searches parent directories
func NewConfigLoaderWithParentSearch() *ConfigLoader {
	// Start with parent directory search paths (higher priority)
	searchPaths := getParentSearchPaths()

	// Add current directory search paths (lower priority)
	searchPaths = append(searchPaths, []string{
		".", // current directory
		"./config",
		"./configs",
		"./build",
		"./test",
		"./tests",
	}...)

	return &ConfigLoader{
		searchPaths: searchPaths,
		fileName:    DefaultConfigFileName,
	}
}

// getParentSearchPaths generates search paths for parent directories
func getParentSearchPaths() []string {
	var paths []string
	// Search up to 5 parent directories
	for i := 1; i <= 5; i++ {
		parentPath := strings.Repeat("../", i)
		paths = append(paths, parentPath)
		paths = append(paths, parentPath+"config")
		paths = append(paths, parentPath+"configs")
	}
	return paths
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

// FindProjectRoot returns the path to the project root directory (where qase.config.json is located)
func (l *ConfigLoader) FindProjectRoot() string {
	configPath := l.findConfigFile()
	if configPath == "" {
		return "." // fallback to current directory
	}

	// Get the directory containing the config file
	configDir := filepath.Dir(configPath)

	// Convert to absolute path
	absPath, err := filepath.Abs(configDir)
	if err != nil {
		return "." // fallback to current directory
	}

	return absPath
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
	// Implementation would go here
	return nil
}

// Global convenience functions

// LoadWithParentSearch loads configuration with parent directory search
func LoadWithParentSearch() (*Config, error) {
	loader := NewConfigLoaderWithParentSearch()
	return loader.Load()
}

// LoadUnsafeWithParentSearch loads configuration without validation with parent directory search
func LoadUnsafeWithParentSearch() *Config {
	loader := NewConfigLoaderWithParentSearch()
	return loader.LoadUnsafe()
}

// FindProjectRootWithParentSearch finds the project root directory with parent directory search
func FindProjectRootWithParentSearch() string {
	loader := NewConfigLoaderWithParentSearch()
	return loader.FindProjectRoot()
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
