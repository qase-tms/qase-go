package _go

import (
	"fmt"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/logging"
)

func main() {
	fmt.Println("Qase Go Logging Example")
	fmt.Println("========================")

	// Example 1: Basic logger configuration
	fmt.Println("\n1. Basic Logger Configuration")
	basicExample()

	// Example 2: Global logger
	fmt.Println("\n2. Global Logger")
	globalLoggerExample()

	// Example 3: Custom configuration
	fmt.Println("\n3. Custom Configuration")
	customConfigExample()

	// Example 4: Log levels
	fmt.Println("\n4. Log Levels")
	logLevelsExample()

	// Example 5: File logging
	fmt.Println("\n5. File Logging")
	fileLoggingExample()

	fmt.Println("\nExamples completed. Check the logs directory for log files.")
}

func basicExample() {
	// Create a basic logger with default configuration
	config := logging.DefaultLoggerConfig()
	logger, err := logging.NewLogger(config)
	if err != nil {
		fmt.Printf("Error creating logger: %v\n", err)
		return
	}
	defer logger.Close()

	logger.Info("Basic logger created successfully")
	logger.Debug("This is a debug message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}

func globalLoggerExample() {
	// Initialize global logger
	config := logging.DefaultLoggerConfig()
	if err := logging.InitializeGlobalLogger(config); err != nil {
		fmt.Printf("Error initializing global logger: %v\n", err)
		return
	}

	// Use global logging functions
	logging.Info("Global logger initialized")
	logging.Debug("Debug message from global logger")
	logging.Warn("Warning from global logger")
	logging.Error("Error from global logger")

	// Close global logger
	logging.CloseGlobalLogger()
}

func customConfigExample() {
	// Create custom logger configuration
	config := logging.LoggerConfig{
		LogToConsole: true,
		LogToFile:    true,
		LogDir:       "./logs",
		LogFileName:  "custom_example.log",
		LogLevel:     logging.DEBUG,
	}

	logger, err := logging.NewLogger(config)
	if err != nil {
		fmt.Printf("Error creating custom logger: %v\n", err)
		return
	}
	defer logger.Close()

	logger.Info("Custom logger created with DEBUG level")
	logger.Debug("This debug message should be visible")
	logger.Info("Custom configuration working")
}

func logLevelsExample() {
	// Create logger with INFO level (default)
	config := logging.DefaultLoggerConfig()
	config.LogLevel = logging.INFO

	logger, err := logging.NewLogger(config)
	if err != nil {
		fmt.Printf("Error creating logger: %v\n", err)
		return
	}
	defer logger.Close()

	logger.Debug("This DEBUG message should NOT be visible (level too low)")
	logger.Info("This INFO message should be visible")
	logger.Warn("This WARN message should be visible")
	logger.Error("This ERROR message should be visible")

	// Change log level to DEBUG
	logger.SetLogLevel(logging.DEBUG)
	logger.Info("Log level changed to DEBUG")
	logger.Debug("Now this DEBUG message should be visible")
}

func fileLoggingExample() {
	// Create logger that only writes to file
	config := logging.LoggerConfig{
		LogToConsole: false, // Disable console output
		LogToFile:    true,  // Enable file output
		LogDir:       "./logs",
		LogFileName:  "file_only.log",
		LogLevel:     logging.INFO,
	}

	logger, err := logging.NewLogger(config)
	if err != nil {
		fmt.Printf("Error creating file-only logger: %v\n", err)
		return
	}
	defer logger.Close()

	logger.Info("This message goes only to file")
	logger.Warn("File-only logging working correctly")
	logger.Error("Error message in file")

	// Get log file path
	logPath := logger.GetLogFilePath()
	fmt.Printf("Log file created at: %s\n", logPath)

	// Demonstrate log rotation
	time.Sleep(1 * time.Second) // Small delay to ensure file is written
	if err := logger.RotateLogFile(); err != nil {
		fmt.Printf("Error rotating log file: %v\n", err)
	} else {
		logger.Info("Log file rotated successfully")
		newLogPath := logger.GetLogFilePath()
		fmt.Printf("New log file created at: %s\n", newLogPath)
	}
}
