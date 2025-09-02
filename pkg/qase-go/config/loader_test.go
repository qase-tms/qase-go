package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfigLoaderWithParentSearch(t *testing.T) {
	// Create temporary directory structure
	tempDir := t.TempDir()
	rootDir := filepath.Join(tempDir, "project")
	subDir := filepath.Join(rootDir, "suite01")

	// Create directories
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("Failed to create directories: %v", err)
	}

	// Create config file in root directory
	configFile := filepath.Join(rootDir, DefaultConfigFileName)
	testConfig := NewConfig()
	testConfig.Mode = "testops"
	testConfig.TestOps.API.Token = "parent-search-token"
	testConfig.TestOps.Project = "PARENT"

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

	// Test parent search loader
	loader := NewConfigLoaderWithParentSearch()
	config, err := loader.Load()
	if err != nil {
		t.Fatalf("Failed to load config with parent search: %v", err)
	}

	// Verify config was loaded from parent directory
	if config.TestOps.API.Token != "parent-search-token" {
		t.Errorf("Expected token 'parent-search-token', got '%s'", config.TestOps.API.Token)
	}
	if config.TestOps.Project != "PARENT" {
		t.Errorf("Expected project 'PARENT', got '%s'", config.TestOps.Project)
	}
}

func TestLoadWithParentSearch(t *testing.T) {
	// Create temporary directory structure
	tempDir := t.TempDir()
	rootDir := filepath.Join(tempDir, "project")
	subDir := filepath.Join(rootDir, "suite01")

	// Create directories
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("Failed to create directories: %v", err)
	}

	// Create config file in root directory
	configFile := filepath.Join(rootDir, DefaultConfigFileName)
	testConfig := NewConfig()
	testConfig.Mode = "testops"
	testConfig.TestOps.API.Token = "global-parent-token"
	testConfig.TestOps.Project = "GLOBAL"

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

	// Test global function
	config, err := LoadWithParentSearch()
	if err != nil {
		t.Fatalf("Failed to load config with parent search: %v", err)
	}

	// Verify config was loaded from parent directory
	if config.TestOps.API.Token != "global-parent-token" {
		t.Errorf("Expected token 'global-parent-token', got '%s'", config.TestOps.API.Token)
	}
	if config.TestOps.Project != "GLOBAL" {
		t.Errorf("Expected project 'GLOBAL', got '%s'", config.TestOps.Project)
	}
}

func TestLoadUnsafeWithParentSearch(t *testing.T) {
	// Create temporary directory structure
	tempDir := t.TempDir()
	rootDir := filepath.Join(tempDir, "project")
	subDir := filepath.Join(rootDir, "suite01")

	// Create directories
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("Failed to create directories: %v", err)
	}

	// Create config file in root directory
	configFile := filepath.Join(rootDir, DefaultConfigFileName)
	testConfig := NewConfig()
	testConfig.Mode = "testops"
	testConfig.TestOps.API.Token = "unsafe-parent-token"
	testConfig.TestOps.Project = "UNSAFE"

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

	// Test global function
	config := LoadUnsafeWithParentSearch()

	// Verify config was loaded from parent directory
	if config.TestOps.API.Token != "unsafe-parent-token" {
		t.Errorf("Expected token 'unsafe-parent-token', got '%s'", config.TestOps.API.Token)
	}
	if config.TestOps.Project != "UNSAFE" {
		t.Errorf("Expected project 'UNSAFE', got '%s'", config.TestOps.Project)
	}
}
