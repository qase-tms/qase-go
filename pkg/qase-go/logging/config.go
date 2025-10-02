package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
)

// DefaultLoggerConfig returns a default logger configuration
func DefaultLoggerConfig() LoggerConfig {
	// Generate filename with current timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("logs_%s.log", timestamp)

	// Find project root directory (where qase.config.json is located)
	projectRoot := config.FindProjectRootWithParentSearch()
	logDir := filepath.Join(projectRoot, "logs")

	return LoggerConfig{
		LogToConsole: true, // Default: console enabled
		LogToFile:    true, // Default: file enabled
		LogDir:       logDir,
		LogFileName:  filename,
		LogLevel:     INFO,
	}
}

// LoadLoggerConfigFromEnvironment loads logger configuration from environment variables
func LoadLoggerConfigFromEnvironment(config *LoggerConfig) {
	// Load console setting from environment
	if console := os.Getenv("QASE_LOGGING_CONSOLE"); console != "" {
		config.LogToConsole = strings.ToLower(console) == "true"
	}

	// Load file setting from environment
	if file := os.Getenv("QASE_LOGGING_FILE"); file != "" {
		config.LogToFile = strings.ToLower(file) == "true"
	}
}

// CreateLogsDirectory creates the logs directory if it doesn't exist
func CreateLogsDirectory(logDir string) error {
	if logDir == "" {
		// Find project root directory (where qase.config.json is located)
		projectRoot := config.FindProjectRootWithParentSearch()
		logDir = filepath.Join(projectRoot, "logs")
	}

	// Create logs directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	return nil
}
