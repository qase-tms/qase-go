package reporters

import (
	"context"
	"fmt"
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

// MockReporter is a mock reporter for testing fallback logic
type MockReporter struct {
	shouldFail bool
	results    []*domain.TestResult
}

func NewMockReporter(shouldFail bool) *MockReporter {
	return &MockReporter{
		shouldFail: shouldFail,
		results:    make([]*domain.TestResult, 0),
	}
}

// StartTestRun method removed - no longer exists in Reporter interface

func (m *MockReporter) AddResult(result *domain.TestResult) error {
	if m.shouldFail {
		return fmt.Errorf("mock reporter add result failed")
	}
	m.results = append(m.results, result)
	return nil
}

func (m *MockReporter) CompleteTestRun(ctx context.Context) error {
	if m.shouldFail {
		return fmt.Errorf("mock reporter complete failed")
	}
	return nil
}

func TestCoreReporter_FallbackLogic(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "testops"
	cfg.Fallback = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"
	// Don't set run ID - this should fail with error, not use fallback

	// Create core reporter - this should fail because run ID is required
	reporter, err := NewCoreReporter(cfg)
	if err == nil {
		t.Error("Expected error when run ID is not provided for TestOps mode")
	}
	if reporter != nil {
		t.Error("Expected nil reporter when run ID is not provided for TestOps mode")
	}
}

func TestCoreReporter_FallbackOnAddResult(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	reporter, err := NewCoreReporter(cfg)
	if err != nil {
		t.Fatalf("Failed to create core reporter: %v", err)
	}

	ctx := context.Background()

	// Add a result
	result := domain.NewTestResult("Test 1")
	result.Execution.Status = domain.StatusPassed

	err = reporter.AddResult(result)
	if err != nil {
		t.Errorf("AddResult should not fail: %v", err)
	}

	// Complete test run
	err = reporter.CompleteTestRun(ctx)
	if err != nil {
		t.Errorf("CompleteTestRun should not fail: %v", err)
	}
}

func TestCoreReporter_NoFallbackAvailable(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "testops"
	cfg.Fallback = "off" // No fallback

	// This should fail because TestOps can't be created and no fallback is available
	reporter, err := NewCoreReporter(cfg)
	if err == nil {
		t.Error("Expected error when TestOps fails and no fallback is available")
	}
	if reporter != nil {
		t.Error("Expected nil reporter when TestOps fails and no fallback is available")
	}
}

func TestCoreReporter_FallbackBecomesMain(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "testops"
	cfg.Fallback = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"
	// Don't set run ID - this should fail with error, not use fallback

	reporter, err := NewCoreReporter(cfg)
	if err == nil {
		t.Error("Expected error when run ID is not provided for TestOps mode")
	}
	if reporter != nil {
		t.Error("Expected nil reporter when run ID is not provided for TestOps mode")
	}
}

func TestCoreReporter_ReportModeNoFallback(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Fallback = "off"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	reporter, err := NewCoreReporter(cfg)
	if err != nil {
		t.Fatalf("Failed to create core reporter: %v", err)
	}

	// Should have 1 reporter (no fallback needed for report mode)
	if reporter.GetReporterCount() != 1 {
		t.Errorf("Expected 1 reporter for report mode, got %d", reporter.GetReporterCount())
	}

	if !reporter.IsReportMode() {
		t.Error("Should be in report mode")
	}
}

func TestCoreReporter_TestOpsModeWithValidConfig(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "testops"
	cfg.Fallback = "report"
	cfg.TestOps.API.Token = "valid-token"
	cfg.TestOps.Project = "TEST"
	cfg.TestOps.Run.ID = func() *int64 { id := int64(123); return &id }() // Set run ID
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	// This should still fail because we don't have a real TestOps client implementation
	// but it should fallback to file reporter
	reporter, err := NewCoreReporter(cfg)
	if err != nil {
		t.Fatalf("Failed to create core reporter: %v", err)
	}

	// Should have 2 reporters (TestOps main + file fallback)
	if reporter.GetReporterCount() != 2 {
		t.Errorf("Expected 2 reporters (TestOps main + file fallback), got %d", reporter.GetReporterCount())
	}

	// Mode should be testops since we provided valid config
	currentMode := reporter.GetCurrentMode()
	if currentMode != "testops" {
		t.Errorf("Should be in testops mode, got %s", currentMode)
	}
}

func TestCoreReporter_AllModeFallbackCombinations(t *testing.T) {
	combinations := []struct {
		name        string
		mode        string
		fallback    string
		expectError bool
		description string
	}{
		{"Report Mode, No Fallback", "report", "off", false, "Should work normally"},
		{"Report Mode, Report Fallback", "report", "report", false, "Should work normally"},
		{"Report Mode, TestOps Fallback", "report", "testops", true, "Should fail without TestOps token"},
		{"TestOps Mode, No Fallback", "testops", "off", true, "Should fail without token"},
		{"TestOps Mode, Report Fallback", "testops", "report", false, "Should fallback to report"},
		{"TestOps Mode, TestOps Fallback", "testops", "testops", true, "Should fail without token"},
		{"Off Mode, No Fallback", "off", "off", true, "Should fail - reporting disabled"},
		{"Off Mode, Report Fallback", "off", "report", true, "Should fail - reporting disabled"},
		{"Off Mode, TestOps Fallback", "off", "testops", true, "Should fail - reporting disabled"},
	}

	for _, test := range combinations {
		t.Run(test.name, func(t *testing.T) {
			cfg := config.NewConfig()
			cfg.Mode = test.mode
			cfg.Fallback = test.fallback
			cfg.Report.Connection.Local.Path = "/tmp/test-reports"

			// For TestOps mode tests that expect success, we need to set run ID
			if test.mode == "testops" && !test.expectError {
				cfg.TestOps.Run.ID = func() *int64 { id := int64(123); return &id }()
			}

			reporter, err := NewCoreReporter(cfg)
			if test.expectError {
				if err == nil {
					t.Errorf("Expected error for %s, but got success", test.description)
				}
				if reporter != nil {
					t.Errorf("Expected nil reporter for %s", test.description)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for %s: %v", test.description, err)
				}
				if reporter == nil {
					t.Errorf("Expected reporter for %s", test.description)
				}
			}
		})
	}
}

func TestCoreReporter_ReportModeWithTestOpsFallback(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Fallback = "testops"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"
	cfg.TestOps.API.Token = "" // Empty token will cause TestOps fallback to fail

	// This should fail because TestOps fallback can't be created without token
	reporter, err := NewCoreReporter(cfg)
	if err == nil {
		t.Error("Expected error when TestOps fallback can't be created without token")
	}
	if reporter != nil {
		t.Error("Expected nil reporter when TestOps fallback can't be created without token")
	}
}

func TestCoreReporter_TestOpsModeWithTestOpsFallback(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "testops"
	cfg.Fallback = "testops"
	cfg.TestOps.API.Token = "" // Empty token will cause both to fail

	// This should fail because both main and fallback are TestOps and both fail
	reporter, err := NewCoreReporter(cfg)
	if err == nil {
		t.Error("Expected error when both main and fallback are TestOps and both fail")
	}
	if reporter != nil {
		t.Error("Expected nil reporter when both main and fallback are TestOps and both fail")
	}
}

func TestCoreReporter_InvalidMode(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "invalid"
	cfg.Fallback = "off"

	reporter, err := NewCoreReporter(cfg)
	if err == nil {
		t.Error("Expected error for invalid mode")
	}
	if reporter != nil {
		t.Error("Expected nil reporter for invalid mode")
	}
}

func TestCoreReporter_InvalidFallback(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Fallback = "invalid"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	reporter, err := NewCoreReporter(cfg)
	if err == nil {
		t.Error("Expected error for invalid fallback")
	}
	if reporter != nil {
		t.Error("Expected nil reporter for invalid fallback")
	}
}
