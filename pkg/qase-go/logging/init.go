package logging

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
)

// Config represents a simplified configuration structure for logging
type Config struct {
	Debug   bool
	Logging LoggingConfig
}

// LoggingConfig represents logging configuration
type LoggingConfig struct {
	Console bool
	File    bool
}

// InitFromConfig initializes the global logger from the main configuration
func InitFromConfig(cfg *Config) error {
	// Create logger config with debug level based on main config
	logLevel := INFO
	if cfg.Debug {
		logLevel = DEBUG
	}

	// Determine console and file logging settings
	// Default values: console=true, file=true
	logToConsole := true
	logToFile := true

	// If logging config is provided, use it
	if cfg.Logging.Console || cfg.Logging.File {
		// At least one logging option is explicitly set
		logToConsole = cfg.Logging.Console
		logToFile = cfg.Logging.File
	}

	// Generate filename with current timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("logs_%s.log", timestamp)

	// Find project root directory (where qase.config.json is located)
	projectRoot := config.FindProjectRootWithParentSearch()
	logDir := filepath.Join(projectRoot, "logs")

	loggerConfig := LoggerConfig{
		LogToConsole: logToConsole,
		LogToFile:    logToFile,
		LogDir:       logDir,   // Use project root logs directory
		LogFileName:  filename, // Generated filename with timestamp
		LogLevel:     logLevel,
	}

	// Initialize global logger
	return InitializeGlobalLogger(loggerConfig)
}

// parseLogLevel is no longer needed - using boolean debug flag instead

// GetLoggerFromConfig creates a new logger instance from the main configuration
func GetLoggerFromConfig(cfg *Config) (*Logger, error) {
	// Create logger config with debug level based on main config
	logLevel := INFO
	if cfg.Debug {
		logLevel = DEBUG
	}

	// Determine console and file logging settings
	// Default values: console=true, file=true
	logToConsole := true
	logToFile := true

	// If logging config is provided, use it
	if cfg.Logging.Console || cfg.Logging.File {
		// At least one logging option is explicitly set
		logToConsole = cfg.Logging.Console
		logToFile = cfg.Logging.File
	}

	// Generate filename with current timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("logs_%s.log", timestamp)

	// Find project root directory (where qase.config.json is located)
	projectRoot := config.FindProjectRootWithParentSearch()
	logDir := filepath.Join(projectRoot, "logs")

	loggerConfig := LoggerConfig{
		LogToConsole: logToConsole,
		LogToFile:    logToFile,
		LogDir:       logDir,   // Use project root logs directory
		LogFileName:  filename, // Generated filename with timestamp
		LogLevel:     logLevel,
	}

	// Create and return new logger
	return NewLogger(loggerConfig)
}
