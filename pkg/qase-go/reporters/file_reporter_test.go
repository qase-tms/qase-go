package reporters

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

func TestNewFileReporter(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"
	cfg.Report.Connection.Local.Format = "json"

	config := FileReporterConfig{
		Config:         cfg,
		CustomFilename: "test-report",
	}

	reporter := NewFileReporter(config)

	if reporter == nil {
		t.Fatal("reporter should not be nil")
	}
	if len(reporter.results) != 0 {
		t.Error("results slice should be empty initially")
	}
}

func TestFileReporter_StartTestRun(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	config := FileReporterConfig{
		Config:         cfg,
		CustomFilename: "test-report",
	}

	reporter := NewFileReporter(config)
	ctx := context.Background()

	err := reporter.StartTestRun(ctx)
	if err != nil {
		t.Errorf("StartTestRun should not return error: %v", err)
	}

	if reporter.startTime == 0 {
		t.Error("start time should be set")
	}
}

func TestFileReporter_AddResult(t *testing.T) {
	cfg := config.NewConfig()
	cfg.Report.Connection.Local.Path = "/tmp/test-reports"

	config := FileReporterConfig{
		Config:         cfg,
		CustomFilename: "test-report",
	}

	reporter := NewFileReporter(config)

	tests := []struct {
		name        string
		result      *domain.TestResult
		expectError bool
	}{
		{
			name:        "valid result",
			result:      domain.NewTestResult("Test 1"),
			expectError: false,
		},
		{
			name:        "nil result",
			result:      nil,
			expectError: true,
		},
		{
			name:        "empty title",
			result:      domain.NewTestResult(""),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initialCount := len(reporter.results)
			err := reporter.AddResult(tt.result)

			if tt.expectError && err == nil {
				t.Error("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tt.expectError {
				if len(reporter.results) != initialCount+1 {
					t.Errorf("expected %d results, got %d", initialCount+1, len(reporter.results))
				}
			}
		})
	}
}

func TestFileReporter_CompleteTestRun(t *testing.T) {
	// Create temporary directory for test
	tmpDir := t.TempDir()

	cfg := config.NewConfig()
	cfg.Report.Connection.Local.Path = tmpDir
	cfg.Report.Connection.Local.Format = "json"

	config := FileReporterConfig{
		Config:         cfg,
		CustomFilename: "test-report",
	}

	reporter := NewFileReporter(config)
	ctx := context.Background()

	// Start test run
	err := reporter.StartTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to start test run: %v", err)
	}

	// Add some test results
	result1 := domain.NewTestResult("Test 1")
	result1.Execution.Status = domain.StatusPassed
	now := time.Now().Unix()
	result1.Execution.StartTime = &now
	result1.Execution.EndTime = &now

	result2 := domain.NewTestResult("Test 2")
	result2.Execution.Status = domain.StatusFailed
	result2.Execution.StartTime = &now
	result2.Execution.EndTime = &now

	err = reporter.AddResult(result1)
	if err != nil {
		t.Fatalf("failed to add result 1: %v", err)
	}

	err = reporter.AddResult(result2)
	if err != nil {
		t.Fatalf("failed to add result 2: %v", err)
	}

	// Complete test run
	err = reporter.CompleteTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to complete test run: %v", err)
	}

	// Check that report file was created
	reportPath := filepath.Join(tmpDir, "test-report.json")
	if _, err := os.Stat(reportPath); os.IsNotExist(err) {
		t.Fatalf("report file was not created: %s", reportPath)
	}

	// Read and validate report content
	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report file: %v", err)
	}

	var report Report
	err = json.Unmarshal(data, &report)
	if err != nil {
		t.Fatalf("failed to unmarshal report: %v", err)
	}

	// Validate report structure
	if report.QaseReport.Version != "1.0.0" {
		t.Errorf("expected version 1.0.0, got %s", report.QaseReport.Version)
	}

	summary := report.QaseReport.Summary
	if summary.TotalTests != 2 {
		t.Errorf("expected 2 total tests, got %d", summary.TotalTests)
	}
	if summary.PassedTests != 1 {
		t.Errorf("expected 1 passed test, got %d", summary.PassedTests)
	}
	if summary.FailedTests != 1 {
		t.Errorf("expected 1 failed test, got %d", summary.FailedTests)
	}
	if summary.Status != "failed" {
		t.Errorf("expected status 'failed', got %s", summary.Status)
	}

	if len(report.QaseReport.Results) != 2 {
		t.Errorf("expected 2 results, got %d", len(report.QaseReport.Results))
	}

	// Validate first result
	result1Report := report.QaseReport.Results[0]
	if result1Report.Title != "Test 1" {
		t.Errorf("expected title 'Test 1', got %s", result1Report.Title)
	}
	if result1Report.Execution.Status != "passed" {
		t.Errorf("expected status 'passed', got %s", result1Report.Execution.Status)
	}

	// Validate second result
	result2Report := report.QaseReport.Results[1]
	if result2Report.Title != "Test 2" {
		t.Errorf("expected title 'Test 2', got %s", result2Report.Title)
	}
	if result2Report.Execution.Status != "failed" {
		t.Errorf("expected status 'failed', got %s", result2Report.Execution.Status)
	}
}

func TestFileReporter_CompleteTestRun_EmptyResults(t *testing.T) {
	tmpDir := t.TempDir()
	
	cfg := config.NewConfig()
	cfg.Report.Connection.Local.Path = tmpDir
	
	config := FileReporterConfig{
		Config:         cfg,
		CustomFilename: "empty-report",
	}
	
	reporter := NewFileReporter(config)
	ctx := context.Background()
	
	// Start test run
	err := reporter.StartTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to start test run: %v", err)
	}
	
	// Complete test run without adding results
	err = reporter.CompleteTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to complete test run: %v", err)
	}
	
	// Check that report file was created
	reportPath := filepath.Join(tmpDir, "empty-report.json")
	if _, err := os.Stat(reportPath); os.IsNotExist(err) {
		t.Fatalf("report file was not created: %s", reportPath)
	}
	
	// Read and validate report content
	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report file: %v", err)
	}
	
	var report Report
	err = json.Unmarshal(data, &report)
	if err != nil {
		t.Fatalf("failed to unmarshal report: %v", err)
	}
	
	summary := report.QaseReport.Summary
	if summary.TotalTests != 0 {
		t.Errorf("expected 0 total tests, got %d", summary.TotalTests)
	}
	if summary.Status != "skipped" {
		t.Errorf("expected status 'skipped', got %s", summary.Status)
	}
}

func TestFileReporter_CompleteTestRun_DefaultFilename(t *testing.T) {
	tmpDir := t.TempDir()
	
	cfg := config.NewConfig()
	cfg.Report.Connection.Local.Path = tmpDir
	
	config := FileReporterConfig{
		Config: cfg,
		// CustomFilename not set, should use default
	}
	
	reporter := NewFileReporter(config)
	ctx := context.Background()
	
	err := reporter.StartTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to start test run: %v", err)
	}
	
	err = reporter.CompleteTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to complete test run: %v", err)
	}
	
	// Check that default filename was used
	reportPath := filepath.Join(tmpDir, "qase-report.json")
	if _, err := os.Stat(reportPath); os.IsNotExist(err) {
		t.Fatalf("report file was not created with default name: %s", reportPath)
	}
}

func TestFileReporter_CompleteTestRun_CreateDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "reports", "subdir")
	
	cfg := config.NewConfig()
	cfg.Report.Connection.Local.Path = newDir
	
	config := FileReporterConfig{
		Config:         cfg,
		CustomFilename: "test-report",
	}
	
	reporter := NewFileReporter(config)
	ctx := context.Background()
	
	err := reporter.StartTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to start test run: %v", err)
	}
	
	err = reporter.CompleteTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to complete test run: %v", err)
	}
	
	// Check that directory was created
	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		t.Fatalf("directory was not created: %s", newDir)
	}
	
	// Check that report file was created
	reportPath := filepath.Join(newDir, "test-report.json")
	if _, err := os.Stat(reportPath); os.IsNotExist(err) {
		t.Fatalf("report file was not created: %s", reportPath)
	}
}

func TestFileReporter_WorkflowIntegration(t *testing.T) {
	tmpDir := t.TempDir()
	
	cfg := config.NewConfig()
	cfg.Report.Connection.Local.Path = tmpDir
	
	config := FileReporterConfig{
		Config:         cfg,
		CustomFilename: "workflow-test",
	}
	
	reporter := NewFileReporter(config)
	ctx := context.Background()
	
	// Simulate complete workflow
	err := reporter.StartTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to start test run: %v", err)
	}
	
	// Add multiple results with different statuses
	results := []*domain.TestResult{
		CreateSimpleResult("Passing Test", domain.StatusPassed),
		CreateSimpleResult("Failing Test", domain.StatusFailed),
		CreateSimpleResult("Skipped Test", domain.StatusSkipped),
	}
	
	for _, result := range results {
		err = reporter.AddResult(result)
		if err != nil {
			t.Fatalf("failed to add result: %v", err)
		}
	}
	
	err = reporter.CompleteTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to complete test run: %v", err)
	}
	
	// Validate report
	reportPath := filepath.Join(tmpDir, "workflow-test.json")
	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report file: %v", err)
	}
	
	var report Report
	err = json.Unmarshal(data, &report)
	if err != nil {
		t.Fatalf("failed to unmarshal report: %v", err)
	}
	
	summary := report.QaseReport.Summary
	if summary.TotalTests != 3 {
		t.Errorf("expected 3 total tests, got %d", summary.TotalTests)
	}
	if summary.PassedTests != 1 {
		t.Errorf("expected 1 passed test, got %d", summary.PassedTests)
	}
	if summary.FailedTests != 1 {
		t.Errorf("expected 1 failed test, got %d", summary.FailedTests)
	}
	if summary.SkippedTests != 1 {
		t.Errorf("expected 1 skipped test, got %d", summary.SkippedTests)
	}
	if summary.Status != "failed" {
		t.Errorf("expected status 'failed', got %s", summary.Status)
	}
}
