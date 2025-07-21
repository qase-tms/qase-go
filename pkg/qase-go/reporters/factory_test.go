package reporters

import (
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
)

func TestNewReporterFactory(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	factory := NewReporterFactory(cfg)
	if factory == nil {
		t.Fatal("Factory should not be nil")
	}

	if factory.GetConfig() != cfg {
		t.Error("Factory should return the same config")
	}

	if factory.IsInitialized() {
		t.Error("Factory should not be initialized initially")
	}
}

func TestReporterFactory_GetReporter(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	factory := NewReporterFactory(cfg)

	// Get reporter first time
	reporter1, err := factory.GetReporter()
	if err != nil {
		t.Fatalf("Failed to get reporter: %v", err)
	}

	if reporter1 == nil {
		t.Fatal("Reporter should not be nil")
	}

	if !factory.IsInitialized() {
		t.Error("Factory should be initialized after getting reporter")
	}

	// Get reporter second time - should be the same instance
	reporter2, err := factory.GetReporter()
	if err != nil {
		t.Fatalf("Failed to get reporter second time: %v", err)
	}

	if reporter1 != reporter2 {
		t.Error("Should return the same reporter instance (singleton)")
	}
}

func TestReporterFactory_Reset(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	factory := NewReporterFactory(cfg)

	// Get reporter
	reporter1, err := factory.GetReporter()
	if err != nil {
		t.Fatalf("Failed to get reporter: %v", err)
	}

	if !factory.IsInitialized() {
		t.Error("Factory should be initialized")
	}

	// Reset factory
	factory.Reset()

	if factory.IsInitialized() {
		t.Error("Factory should not be initialized after reset")
	}

	// Get reporter again - should be a new instance
	reporter2, err := factory.GetReporter()
	if err != nil {
		t.Fatalf("Failed to get reporter after reset: %v", err)
	}

	if reporter1 == reporter2 {
		t.Error("Should return different reporter instance after reset")
	}
}

func TestReporterFactory_UpdateConfig(t *testing.T) {
	cfg1 := config.NewConfig()
	cfg1.Mode = "report"
	cfg1.Report.Connection.Local.Path = "/tmp/test-reports"

	factory := NewReporterFactory(cfg1)

	// Get reporter with first config
	reporter1, err := factory.GetReporter()
	if err != nil {
		t.Fatalf("Failed to get reporter: %v", err)
	}

	// Update config
	cfg2 := config.NewConfig()
	cfg2.Mode = "report"
	cfg2.Report.Connection.Local.Path = "/tmp/other-reports"
	factory.UpdateConfig(cfg2)

	if factory.IsInitialized() {
		t.Error("Factory should not be initialized after config update")
	}

	// Get reporter with new config
	reporter2, err := factory.GetReporter()
	if err != nil {
		t.Fatalf("Failed to get reporter with new config: %v", err)
	}

	if reporter1 == reporter2 {
		t.Error("Should return different reporter instance after config update")
	}

	if factory.GetConfig() != cfg2 {
		t.Error("Factory should return the updated config")
	}
}

func TestReporterFactory_GetReporterError(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "testops"
	cfg.Fallback = "off"
	// No token provided, should fail

	factory := NewReporterFactory(cfg)

	reporter, err := factory.GetReporter()
	if err == nil {
		t.Error("Expected error for invalid configuration")
	}
	if reporter != nil {
		t.Error("Expected nil reporter for invalid configuration")
	}
}

func TestGlobalFactory(t *testing.T) {
	// Reset global factory before test
	ResetGlobalFactory()

	if IsGlobalFactoryInitialized() {
		t.Error("Global factory should not be initialized initially")
	}

	// Try to get global reporter without initialization
	reporter, err := GetGlobalReporter()
	if err == nil {
		t.Error("Expected error when global factory not initialized")
	}
	if reporter != nil {
		t.Error("Expected nil reporter when global factory not initialized")
	}

	// Initialize global factory
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"
	InitializeGlobalFactory(cfg)

	if !IsGlobalFactoryInitialized() {
		t.Error("Global factory should be initialized")
	}

	// Get global reporter
	reporter1, err := GetGlobalReporter()
	if err != nil {
		t.Fatalf("Failed to get global reporter: %v", err)
	}

	if reporter1 == nil {
		t.Fatal("Global reporter should not be nil")
	}

	// Get global reporter again - should be the same instance
	reporter2, err := GetGlobalReporter()
	if err != nil {
		t.Fatalf("Failed to get global reporter second time: %v", err)
	}

	if reporter1 != reporter2 {
		t.Error("Should return the same global reporter instance")
	}
}

func TestGlobalFactory_UpdateConfig(t *testing.T) {
	// Reset global factory before test
	ResetGlobalFactory()

	// Initialize with first config
	cfg1 := config.NewConfig()
	cfg1.Mode = "report"
	cfg1.Report.Connection.Local.Path = "/tmp/test-reports"
	InitializeGlobalFactory(cfg1)

	// Get reporter
	reporter1, err := GetGlobalReporter()
	if err != nil {
		t.Fatalf("Failed to get global reporter: %v", err)
	}

	// Update config
	cfg2 := config.NewConfig()
	cfg2.Mode = "report"
	cfg2.Report.Connection.Local.Path = "/tmp/other-reports"
	UpdateGlobalConfig(cfg2)

	// Get reporter again - should be a new instance
	reporter2, err := GetGlobalReporter()
	if err != nil {
		t.Fatalf("Failed to get global reporter after config update: %v", err)
	}

	if reporter1 == reporter2 {
		t.Error("Should return different global reporter instance after config update")
	}
}

func TestGlobalFactory_Reset(t *testing.T) {
	// Initialize global factory
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"
	InitializeGlobalFactory(cfg)

	if !IsGlobalFactoryInitialized() {
		t.Error("Global factory should be initialized")
	}

	// Reset global factory
	ResetGlobalFactory()

	if IsGlobalFactoryInitialized() {
		t.Error("Global factory should not be initialized after reset")
	}

	// Try to get global reporter after reset
	reporter, err := GetGlobalReporter()
	if err == nil {
		t.Error("Expected error when global factory not initialized")
	}
	if reporter != nil {
		t.Error("Expected nil reporter when global factory not initialized")
	}
}
