package logging

import (
	"sync"
)

var (
	globalLogger *Logger
	globalMutex  sync.RWMutex
)

// InitializeGlobalLogger initializes the global logger with the given configuration
func InitializeGlobalLogger(config LoggerConfig) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	// Close existing logger if any
	if globalLogger != nil {
		globalLogger.Close()
	}

	// Create new logger
	logger, err := NewLogger(config)
	if err != nil {
		return err
	}

	globalLogger = logger
	return nil
}

// GetGlobalLogger returns the global logger instance
func GetGlobalLogger() *Logger {
	globalMutex.RLock()
	defer globalMutex.RUnlock()
	return globalLogger
}

// SetGlobalLogger sets the global logger instance
func SetGlobalLogger(logger *Logger) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	// Close existing logger if any
	if globalLogger != nil {
		globalLogger.Close()
	}

	globalLogger = logger
}

// CloseGlobalLogger closes the global logger and releases resources
func CloseGlobalLogger() error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	if globalLogger != nil {
		err := globalLogger.Close()
		globalLogger = nil
		return err
	}
	return nil
}

// Global logging functions for convenience

// Debug logs a debug message using the global logger
func Debug(format string, args ...interface{}) {
	if logger := GetGlobalLogger(); logger != nil {
		logger.Debug(format, args...)
	}
}

// Info logs an info message using the global logger
func Info(format string, args ...interface{}) {
	if logger := GetGlobalLogger(); logger != nil {
		logger.Info(format, args...)
	}
}

// Warn logs a warning message using the global logger
func Warn(format string, args ...interface{}) {
	if logger := GetGlobalLogger(); logger != nil {
		logger.Warn(format, args...)
	}
}

// Error logs an error message using the global logger
func Error(format string, args ...interface{}) {
	if logger := GetGlobalLogger(); logger != nil {
		logger.Error(format, args...)
	}
}

// Printf logs a message without level prefix using the global logger
func Printf(format string, args ...interface{}) {
	if logger := GetGlobalLogger(); logger != nil {
		logger.Printf(format, args...)
	}
}

// Print logs a message without level prefix using the global logger
func Print(v ...interface{}) {
	if logger := GetGlobalLogger(); logger != nil {
		logger.Print(v...)
	}
}
