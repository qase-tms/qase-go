package qase

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
)

func TestInitializeGlobal(t *testing.T) {
	// Reset global state before test
	ResetGlobal()

	// Create temporary directory structure
	tempDir := t.TempDir()
	rootDir := filepath.Join(tempDir, "project")
	subDir := filepath.Join(rootDir, "suite01")

	// Create directories
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("Failed to create directories: %v", err)
	}

	// Create config file in root directory
	configFile := filepath.Join(rootDir, "qase.config.json")
	testConfig := config.NewConfig()
	testConfig.Mode = "report" // Use report mode to avoid API calls
	testConfig.Fallback = "off"
	testConfig.Debug = false

	err := testConfig.SaveToFile(configFile)
	if err != nil {
		t.Fatalf("Failed to save test config: %v", err)
	}

	// Change to subdirectory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(subDir); err != nil {
		t.Fatalf("Failed to change to subdirectory: %v", err)
	}

	// Test global initialization
	err = InitializeGlobal()
	if err != nil {
		t.Fatalf("Failed to initialize globally: %v", err)
	}

	// Verify global initialization
	if !IsGlobalInitialized() {
		t.Error("Expected global initialization to be true")
	}

	// Test that second call doesn't fail
	err = InitializeGlobal()
	if err != nil {
		t.Fatalf("Second call to InitializeGlobal should not fail: %v", err)
	}
}

func TestInitializeGlobalWithConfig(t *testing.T) {
	// Reset global state before test
	ResetGlobal()

	// Create custom config
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Fallback = "off"
	cfg.Debug = false

	// Test global initialization with custom config
	err := InitializeGlobalWithConfig(cfg)
	if err != nil {
		t.Fatalf("Failed to initialize globally with config: %v", err)
	}

	// Verify global initialization
	if !IsGlobalInitialized() {
		t.Error("Expected global initialization to be true")
	}

	// Test that second call doesn't fail
	err = InitializeGlobalWithConfig(cfg)
	if err != nil {
		t.Fatalf("Second call to InitializeGlobalWithConfig should not fail: %v", err)
	}
}

func TestIsGlobalInitialized(t *testing.T) {
	// Reset global state before test
	ResetGlobal()

	// Should be false initially
	if IsGlobalInitialized() {
		t.Error("Expected global initialization to be false initially")
	}

	// Initialize
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Fallback = "off"
	cfg.Debug = false

	err := InitializeGlobalWithConfig(cfg)
	if err != nil {
		t.Fatalf("Failed to initialize: %v", err)
	}

	// Should be true after initialization
	if !IsGlobalInitialized() {
		t.Error("Expected global initialization to be true after initialization")
	}
}

func TestResetGlobal(t *testing.T) {
	// Initialize first
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Fallback = "off"
	cfg.Debug = false

	err := InitializeGlobalWithConfig(cfg)
	if err != nil {
		t.Fatalf("Failed to initialize: %v", err)
	}

	// Verify it's initialized
	if !IsGlobalInitialized() {
		t.Error("Expected global initialization to be true")
	}

	// Reset
	ResetGlobal()

	// Should be false after reset
	if IsGlobalInitialized() {
		t.Error("Expected global initialization to be false after reset")
	}
}
