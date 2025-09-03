package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// Logger provides logging functionality with both console and file output
type Logger struct {
	consoleLogger *log.Logger
	fileLogger    *log.Logger
	file          *os.File
	mutex         sync.Mutex
	config        LoggerConfig
}

// LoggerConfig holds configuration for the logger
type LoggerConfig struct {
	LogToConsole bool
	LogToFile    bool
	LogDir       string
	LogFileName  string
	LogLevel     LogLevel
}

// LogLevel represents the logging level
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// String returns the string representation of the log level
func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// NewLogger creates a new logger instance
func NewLogger(config LoggerConfig) (*Logger, error) {
	logger := &Logger{
		consoleLogger: log.New(os.Stdout, "", log.LstdFlags),
		config:        config,
	}

	// Initialize console logging
	if config.LogToConsole {
		logger.consoleLogger = log.New(os.Stdout, "", log.LstdFlags)
	}

	// Initialize file logging
	if config.LogToFile {
		if err := logger.initializeFileLogging(); err != nil {
			return nil, fmt.Errorf("failed to initialize file logging: %w", err)
		}
	}

	return logger, nil
}

// initializeFileLogging sets up file logging
func (l *Logger) initializeFileLogging() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// Create logs directory if it doesn't exist
	if err := os.MkdirAll(l.config.LogDir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %w", err)
	}

	// Use fixed filename
	filename := l.config.LogFileName

	// Create or open log file
	logPath := filepath.Join(l.config.LogDir, filename)
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	l.file = file
	l.fileLogger = log.New(file, "", log.LstdFlags)

	return nil
}

// log writes a log message to all configured outputs
func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if level < l.config.LogLevel {
		return
	}

	// Format the message with level only (timestamp is added by log.Logger)
	message := fmt.Sprintf(format, args...)
	formattedMessage := fmt.Sprintf("[%s] %s", level.String(), message)

	// Write to console if enabled
	if l.config.LogToConsole {
		l.consoleLogger.Print(formattedMessage)
	}

	// Write to file if enabled
	if l.config.LogToFile && l.fileLogger != nil {
		l.mutex.Lock()
		defer l.mutex.Unlock()
		l.fileLogger.Print(formattedMessage)
	}
}

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(WARN, format, args...)
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

// Printf logs a message without level prefix (for backward compatibility)
func (l *Logger) Printf(format string, args ...interface{}) {
	l.Info(format, args...)
}

// Print logs a message without level prefix (for backward compatibility)
func (l *Logger) Print(v ...interface{}) {
	l.Info(fmt.Sprint(v...))
}

// Close closes the logger and releases resources
func (l *Logger) Close() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

// RotateLogFile rotates the log file by closing the current one and opening a new one
func (l *Logger) RotateLogFile() error {
	if !l.config.LogToFile {
		return nil
	}

	// Close current file first
	l.mutex.Lock()
	if l.file != nil {
		l.file.Close()
		l.file = nil
		l.fileLogger = nil
	}
	l.mutex.Unlock()

	// Initialize new file logging (this will acquire its own lock)
	return l.initializeFileLogging()
}

// GetLogFilePath returns the current log file path
func (l *Logger) GetLogFilePath() string {
	if l.file != nil {
		return l.file.Name()
	}
	return ""
}

// SetLogLevel sets the logging level
func (l *Logger) SetLogLevel(level LogLevel) {
	l.config.LogLevel = level
}

// IsConsoleEnabled returns true if console logging is enabled
func (l *Logger) IsConsoleEnabled() bool {
	return l.config.LogToConsole
}

// IsFileEnabled returns true if file logging is enabled
func (l *Logger) IsFileEnabled() bool {
	return l.config.LogToFile
}
