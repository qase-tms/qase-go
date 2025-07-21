package reporters

import (
	"context"
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

func TestNewCoreReporter(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	reporter, err := NewCoreReporter(cfg)
	if err != nil {
		t.Fatalf("Failed to create core reporter: %v", err)
	}

	if reporter == nil {
		t.Fatal("Core reporter should not be nil")
	}

	if reporter.GetReporterCount() != 1 {
		t.Errorf("Expected 1 reporter, got %d", reporter.GetReporterCount())
	}

	if !reporter.IsReportMode() {
		t.Error("Should be in report mode")
	}
}

func TestNewCoreReporter_NilConfig(t *testing.T) {
	reporter, err := NewCoreReporter(nil)
	if err == nil {
		t.Error("Expected error for nil config")
	}
	if reporter != nil {
		t.Error("Expected nil reporter for nil config")
	}
}

func TestNewCoreReporter_OffMode(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "off"

	reporter, err := NewCoreReporter(cfg)
	if err == nil {
		t.Error("Expected error for off mode")
	}
	if reporter != nil {
		t.Error("Expected nil reporter for off mode")
	}
}

func TestNewCoreReporter_TestOpsMode_NoToken(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "testops"
	cfg.Fallback = "off"
	
	reporter, err := NewCoreReporter(cfg)
	if err == nil {
		t.Error("Expected error for TestOps mode without token")
	}
	if reporter != nil {
		t.Error("Expected nil reporter for TestOps mode without token")
	}
}

func TestNewCoreReporter_TestOpsMode_TestOpsFallback(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "testops"
	cfg.Fallback = "testops"
	
	// This should fail because TestOps client creation fails for both main and fallback
	reporter, err := NewCoreReporter(cfg)
	if err == nil {
		t.Error("Expected error for TestOps mode with TestOps fallback without token")
	}
	if reporter != nil {
		t.Error("Expected nil reporter for TestOps mode with TestOps fallback without token")
	}
}

func TestNewCoreReporter_TestOpsMode_WithFallback(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "testops"
	cfg.Fallback = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	reporter, err := NewCoreReporter(cfg)
	if err != nil {
		t.Fatalf("Failed to create core reporter: %v", err)
	}

	if reporter == nil {
		t.Fatal("Core reporter should not be nil")
	}

	// Should have file reporter as main reporter (fallback became main)
	if reporter.GetReporterCount() != 1 {
		t.Errorf("Expected 1 reporter (fallback became main), got %d", reporter.GetReporterCount())
	}
}

func TestCoreReporter_StartTestRun(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	reporter, err := NewCoreReporter(cfg)
	if err != nil {
		t.Fatalf("Failed to create core reporter: %v", err)
	}

	ctx := context.Background()
	err = reporter.StartTestRun(ctx)
	if err != nil {
		t.Errorf("StartTestRun should not return error: %v", err)
	}

	// Results are not stored in core reporter anymore
	// This test is now a placeholder
}

func TestCoreReporter_AddResult(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	reporter, err := NewCoreReporter(cfg)
	if err != nil {
		t.Fatalf("Failed to create core reporter: %v", err)
	}

	ctx := context.Background()
	err = reporter.StartTestRun(ctx)
	if err != nil {
		t.Fatalf("Failed to start test run: %v", err)
	}

	result := domain.NewTestResult("Test 1")
	result.Execution.Status = domain.StatusPassed

	err = reporter.AddResult(result)
	if err != nil {
		t.Errorf("AddResult should not return error: %v", err)
	}

	// Results are not stored in core reporter anymore
	// This test is now a placeholder - we just verify that AddResult doesn't return error
}

func TestCoreReporter_AddResult_NilResult(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	reporter, err := NewCoreReporter(cfg)
	if err != nil {
		t.Fatalf("Failed to create core reporter: %v", err)
	}

	err = reporter.AddResult(nil)
	if err == nil {
		t.Error("Expected error for nil result")
	}
}

func TestCoreReporter_AddResult_EmptyTitle(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	reporter, err := NewCoreReporter(cfg)
	if err != nil {
		t.Fatalf("Failed to create core reporter: %v", err)
	}

	result := domain.NewTestResult("")
	err = reporter.AddResult(result)
	if err == nil {
		t.Error("Expected error for empty title")
	}
}

func TestCoreReporter_CompleteTestRun(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	reporter, err := NewCoreReporter(cfg)
	if err != nil {
		t.Fatalf("Failed to create core reporter: %v", err)
	}

	ctx := context.Background()
	err = reporter.StartTestRun(ctx)
	if err != nil {
		t.Fatalf("Failed to start test run: %v", err)
	}

	result := domain.NewTestResult("Test 1")
	result.Execution.Status = domain.StatusPassed

	err = reporter.AddResult(result)
	if err != nil {
		t.Fatalf("Failed to add result: %v", err)
	}

	err = reporter.CompleteTestRun(ctx)
	if err != nil {
		t.Errorf("CompleteTestRun should not return error: %v", err)
	}
}

func TestCoreReporter_GetRunID(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	reporter, err := NewCoreReporter(cfg)
	if err != nil {
		t.Fatalf("Failed to create core reporter: %v", err)
	}

	// Initially should be nil
	if reporter.GetRunID() != nil {
		t.Error("RunID should be nil initially")
	}

	// Set run ID
	runID := int64(123)
	reporter.SetRunID(runID)

	// Check if set correctly
	resultRunID := reporter.GetRunID()
	if resultRunID == nil {
		t.Error("RunID should not be nil after setting")
	}
	if *resultRunID != runID {
		t.Errorf("Expected run ID %d, got %d", runID, *resultRunID)
	}
}

func TestCoreReporter_ModeChecks(t *testing.T) {
	tests := []struct {
		name     string
		mode     string
		expected bool
	}{
		{"testops mode", "testops", true},
		{"report mode", "report", true},
		{"off mode", "off", true},
		{"invalid mode", "invalid", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.NewConfig()
			cfg.Mode = tt.mode

			// Only test valid modes that can create a reporter
			if tt.mode == "report" {
				cfg.Report.Connection.Local.Path = "/tmp/test-reports"
				reporter, err := NewCoreReporter(cfg)
				if err != nil {
					t.Fatalf("Failed to create reporter: %v", err)
				}

				switch tt.mode {
				case "testops":
					if reporter.IsTestOpsMode() != tt.expected {
						t.Errorf("IsTestOpsMode() returned %v, expected %v", reporter.IsTestOpsMode(), tt.expected)
					}
				case "report":
					if reporter.IsReportMode() != tt.expected {
						t.Errorf("IsReportMode() returned %v, expected %v", reporter.IsReportMode(), tt.expected)
					}
				case "off":
					if reporter.IsOffMode() != tt.expected {
						t.Errorf("IsOffMode() returned %v, expected %v", reporter.IsOffMode(), tt.expected)
					}
				}
			}
		})
	}
}

func TestCoreReporter_HelperMethods(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Mode = "report"
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	reporter, err := NewCoreReporter(cfg)
	if err != nil {
		t.Fatalf("Failed to create core reporter: %v", err)
	}

	// Test CreateSimpleResult
	result := reporter.CreateSimpleResult("Test", domain.StatusPassed)
	if result.Title != "Test" {
		t.Errorf("Expected title 'Test', got %s", result.Title)
	}
	if result.Execution.Status != domain.StatusPassed {
		t.Errorf("Expected status passed, got %s", result.Execution.Status)
	}

	// Test CreateStep
	step := reporter.CreateStep("Action", domain.StepStatusPassed)
	if step.Data.Action != "Action" {
		t.Errorf("Expected action 'Action', got %s", step.Data.Action)
	}
	if step.Execution.Status != domain.StepStatusPassed {
		t.Errorf("Expected step status passed, got %s", step.Execution.Status)
	}

	// Test CreateResultWithSteps
	steps := []domain.TestStep{step}
	resultWithSteps := reporter.CreateResultWithSteps("Test with Steps", domain.StatusPassed, steps)
	if resultWithSteps.Title != "Test with Steps" {
		t.Errorf("Expected title 'Test with Steps', got %s", resultWithSteps.Title)
	}
	if len(resultWithSteps.Steps) != 1 {
		t.Errorf("Expected 1 step, got %d", len(resultWithSteps.Steps))
	}
}
