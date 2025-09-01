package logging

import (
	"fmt"
	"time"
)

// Config represents a simplified configuration structure for logging
// Only needs the Debug flag from main config
type Config struct {
	Debug bool
}

// InitFromConfig initializes the global logger from the main configuration
func InitFromConfig(cfg *Config) error {
	// Create logger config with debug level based on main config
	logLevel := INFO
	if cfg.Debug {
		logLevel = DEBUG
	}

	// Generate filename with current timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("logs_%s.log", timestamp)

	loggerConfig := LoggerConfig{
		LogToConsole: true,     // Always enabled
		LogToFile:    true,     // Always enabled
		LogDir:       "./logs", // Fixed directory
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

	// Generate filename with current timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("logs_%s.log", timestamp)

	loggerConfig := LoggerConfig{
		LogToConsole: true,     // Always enabled
		LogToFile:    true,     // Always enabled
		LogDir:       "./logs", // Fixed directory
		LogFileName:  filename, // Generated filename with timestamp
		LogLevel:     logLevel,
	}

	// Create and return new logger
	return NewLogger(loggerConfig)
}
