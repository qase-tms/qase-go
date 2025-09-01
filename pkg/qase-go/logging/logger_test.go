package logging

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	// Test with console only
	config := LoggerConfig{
		LogToConsole: true,
		LogToFile:    false,
		LogLevel:     INFO,
	}

	logger, err := NewLogger(config)
	if err != nil {
		t.Fatalf("Failed to create console-only logger: %v", err)
	}
	defer logger.Close()

	if !logger.IsConsoleEnabled() {
		t.Error("Console logging should be enabled")
	}
	if logger.IsFileEnabled() {
		t.Error("File logging should be disabled")
	}
}

func TestNewLoggerWithFile(t *testing.T) {
	// Test with file logging
	tempDir := t.TempDir()
	config := LoggerConfig{
		LogToConsole: true,
		LogToFile:    true,
		LogDir:       tempDir,
		LogFileName:  "test.log",
		LogLevel:     INFO,
	}

	logger, err := NewLogger(config)
	if err != nil {
		t.Fatalf("Failed to create file logger: %v", err)
	}
	defer logger.Close()

	if !logger.IsConsoleEnabled() {
		t.Error("Console logging should be enabled")
	}
	if !logger.IsFileEnabled() {
		t.Error("File logging should be enabled")
	}

	// Check if log file was created
	logPath := logger.GetLogFilePath()
	if logPath == "" {
		t.Error("Log file path should not be empty")
	}

	// Check if file exists
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		t.Errorf("Log file should exist at: %s", logPath)
	}
}

func TestLogLevels(t *testing.T) {
	config := LoggerConfig{
		LogToConsole: true,
		LogToFile:    false,
		LogLevel:     INFO,
	}

	logger, err := NewLogger(config)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()

	// Test that DEBUG messages are filtered out at INFO level
	logger.Debug("This should not appear")
	logger.Info("This should appear")
	logger.Warn("This should appear")
	logger.Error("This should appear")

	// Change to DEBUG level
	logger.SetLogLevel(DEBUG)
	logger.Debug("Now this should appear")
}

func TestLogFileCreation(t *testing.T) {
	tempDir := t.TempDir()
	config := LoggerConfig{
		LogToConsole: false,
		LogToFile:    true,
		LogDir:       tempDir,
		LogFileName:  "logs_data.log",
		LogLevel:     INFO,
	}

	logger, err := NewLogger(config)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()

	// Check if logs directory was created
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		t.Errorf("Logs directory should exist at: %s", tempDir)
	}

	// Check if log file was created
	logPath := logger.GetLogFilePath()
	if logPath == "" {
		t.Error("Log file path should not be empty")
	}

	// Verify filename is correct
	filename := filepath.Base(logPath)
	if filename != "logs_data.log" {
		t.Errorf("Log filename should be 'logs_data.log', got: %s", filename)
	}
}

func TestLogRotation(t *testing.T) {
	tempDir := t.TempDir()
	config := LoggerConfig{
		LogToConsole: false,
		LogToFile:    true,
		LogDir:       tempDir,
		LogFileName:  "rotation_test.log",
		LogLevel:     INFO,
	}

	logger, err := NewLogger(config)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()

	// Write some logs
	logger.Info("First log entry")
	originalPath := logger.GetLogFilePath()

	// Rotate log file
	if err := logger.RotateLogFile(); err != nil {
		t.Fatalf("Failed to rotate log file: %v", err)
	}

	// After rotation, the path should remain the same since we use fixed filename
	newPath := logger.GetLogFilePath()
	if newPath != originalPath {
		t.Errorf("Log file path should remain the same after rotation, got: %s, expected: %s", newPath, originalPath)
	}

	// Write more logs to the same file
	logger.Info("Second log entry in same file")
}

func TestGlobalLogger(t *testing.T) {
	// Test global logger initialization
	config := DefaultLoggerConfig()
	config.LogToFile = false // Disable file logging for test

	if err := InitializeGlobalLogger(config); err != nil {
		t.Fatalf("Failed to initialize global logger: %v", err)
	}

	// Test global logging functions
	Info("Global info message")
	Debug("Global debug message")
	Warn("Global warning message")
	Error("Global error message")

	// Test Printf and Print
	Printf("Global printf message: %s", "test")
	Print("Global print message")

	// Clean up
	CloseGlobalLogger()
}

func TestDefaultLoggerConfig(t *testing.T) {
	config := DefaultLoggerConfig()

	if !config.LogToConsole {
		t.Error("Default config should enable console logging")
	}
	if !config.LogToFile {
		t.Error("Default config should enable file logging")
	}
	if config.LogDir != "./logs" {
		t.Errorf("Default log directory should be './logs', got: %s", config.LogDir)
	}
	if config.LogLevel != INFO {
		t.Errorf("Default log level should be INFO, got: %v", config.LogLevel)
	}
}

func TestLogLevelString(t *testing.T) {
	testCases := []struct {
		level    LogLevel
		expected string
	}{
		{DEBUG, "DEBUG"},
		{INFO, "INFO"},
		{WARN, "WARN"},
		{ERROR, "ERROR"},
	}

	for _, tc := range testCases {
		if tc.level.String() != tc.expected {
			t.Errorf("LogLevel.String() for %v should return %s, got %s", tc.level, tc.expected, tc.level.String())
		}
	}
}

func TestConcurrentLogging(t *testing.T) {
	tempDir := t.TempDir()
	config := LoggerConfig{
		LogToConsole: true,
		LogToFile:    true,
		LogDir:       tempDir,
		LogFileName:  "concurrent_test.log",
		LogLevel:     INFO,
	}

	logger, err := NewLogger(config)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()

	// Test concurrent logging
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 10; j++ {
				logger.Info("Concurrent log message %d from goroutine %d", j, id)
				time.Sleep(1 * time.Millisecond)
			}
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}
}
