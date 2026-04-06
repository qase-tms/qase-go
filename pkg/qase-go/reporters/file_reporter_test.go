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
	tmpDir := t.TempDir()

	cfg := config.NewConfig()
	cfg.Report.Connection.Local.Path = tmpDir
	cfg.Report.Connection.Local.Format = "json"

	reporterConfig := FileReporterConfig{
		Config: cfg,
	}

	reporter := NewFileReporter(reporterConfig)
	ctx := context.Background()

	err := reporter.StartTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to start test run: %v", err)
	}

	// Add some test results
	result1 := domain.NewTestResult("Test 1")
	result1.Execution.Status = domain.StatusPassed
	now := time.Now().UnixMilli()
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

	err = reporter.CompleteTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to complete test run: %v", err)
	}

	// Check run.json was created
	runPath := filepath.Join(tmpDir, "run.json")
	if _, err := os.Stat(runPath); os.IsNotExist(err) {
		t.Fatalf("run.json was not created: %s", runPath)
	}

	// Read and validate run.json
	data, err := os.ReadFile(runPath)
	if err != nil {
		t.Fatalf("failed to read run.json: %v", err)
	}

	var runReport RunReport
	err = json.Unmarshal(data, &runReport)
	if err != nil {
		t.Fatalf("failed to unmarshal run.json: %v", err)
	}

	if runReport.Stats.Total != 2 {
		t.Errorf("expected 2 total tests, got %d", runReport.Stats.Total)
	}
	if runReport.Stats.Passed != 1 {
		t.Errorf("expected 1 passed test, got %d", runReport.Stats.Passed)
	}
	if runReport.Stats.Failed != 1 {
		t.Errorf("expected 1 failed test, got %d", runReport.Stats.Failed)
	}
	if len(runReport.Results) != 2 {
		t.Errorf("expected 2 result summaries, got %d", len(runReport.Results))
	}

	// Check result files were created
	resultsDir := filepath.Join(tmpDir, "results")
	result1Path := filepath.Join(resultsDir, result1.ID+".json")
	result2Path := filepath.Join(resultsDir, result2.ID+".json")

	if _, err := os.Stat(result1Path); os.IsNotExist(err) {
		t.Fatalf("result file was not created: %s", result1Path)
	}
	if _, err := os.Stat(result2Path); os.IsNotExist(err) {
		t.Fatalf("result file was not created: %s", result2Path)
	}

	// Read and validate individual result files
	for _, tc := range []struct {
		domainResult *domain.TestResult
		title        string
		status       string
	}{
		{result1, "Test 1", "passed"},
		{result2, "Test 2", "failed"},
	} {
		rPath := filepath.Join(resultsDir, tc.domainResult.ID+".json")
		data, err = os.ReadFile(rPath)
		if err != nil {
			t.Fatalf("failed to read result file %s: %v", rPath, err)
		}

		var resultReport Result
		err = json.Unmarshal(data, &resultReport)
		if err != nil {
			t.Fatalf("failed to unmarshal result: %v", err)
		}

		if resultReport.Title != tc.title {
			t.Errorf("expected title %q, got %s", tc.title, resultReport.Title)
		}
		if resultReport.Execution.Status != tc.status {
			t.Errorf("expected status %q, got %s", tc.status, resultReport.Execution.Status)
		}
	}
}

func TestFileReporter_CompleteTestRun_EmptyResults(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := config.NewConfig()
	cfg.Report.Connection.Local.Path = tmpDir

	reporterConfig := FileReporterConfig{
		Config: cfg,
	}

	reporter := NewFileReporter(reporterConfig)
	ctx := context.Background()

	err := reporter.StartTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to start test run: %v", err)
	}

	err = reporter.CompleteTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to complete test run: %v", err)
	}

	// Check run.json
	runPath := filepath.Join(tmpDir, "run.json")
	if _, err := os.Stat(runPath); os.IsNotExist(err) {
		t.Fatalf("run.json was not created: %s", runPath)
	}

	data, err := os.ReadFile(runPath)
	if err != nil {
		t.Fatalf("failed to read run.json: %v", err)
	}

	var runReport RunReport
	err = json.Unmarshal(data, &runReport)
	if err != nil {
		t.Fatalf("failed to unmarshal run.json: %v", err)
	}

	if runReport.Stats.Total != 0 {
		t.Errorf("expected 0 total tests, got %d", runReport.Stats.Total)
	}
}

func TestFileReporter_CompleteTestRun_CreateDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "reports", "subdir")

	cfg := config.NewConfig()
	cfg.Report.Connection.Local.Path = newDir

	reporterConfig := FileReporterConfig{
		Config: cfg,
	}

	reporter := NewFileReporter(reporterConfig)
	ctx := context.Background()

	err := reporter.StartTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to start test run: %v", err)
	}

	err = reporter.CompleteTestRun(ctx)
	if err != nil {
		t.Fatalf("failed to complete test run: %v", err)
	}

	// Check that directory and run.json were created
	runPath := filepath.Join(newDir, "run.json")
	if _, err := os.Stat(runPath); os.IsNotExist(err) {
		t.Fatalf("run.json was not created: %s", runPath)
	}
}

func TestFileReporter_WorkflowIntegration(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := config.NewConfig()
	cfg.Report.Connection.Local.Path = tmpDir

	reporterConfig := FileReporterConfig{
		Config: cfg,
	}

	reporter := NewFileReporter(reporterConfig)
	ctx := context.Background()

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

	// Validate run.json
	runPath := filepath.Join(tmpDir, "run.json")
	data, err := os.ReadFile(runPath)
	if err != nil {
		t.Fatalf("failed to read run.json: %v", err)
	}

	var runReport RunReport
	err = json.Unmarshal(data, &runReport)
	if err != nil {
		t.Fatalf("failed to unmarshal run.json: %v", err)
	}

	if runReport.Stats.Total != 3 {
		t.Errorf("expected 3 total tests, got %d", runReport.Stats.Total)
	}
	if runReport.Stats.Passed != 1 {
		t.Errorf("expected 1 passed test, got %d", runReport.Stats.Passed)
	}
	if runReport.Stats.Failed != 1 {
		t.Errorf("expected 1 failed test, got %d", runReport.Stats.Failed)
	}
	if runReport.Stats.Skipped != 1 {
		t.Errorf("expected 1 skipped test, got %d", runReport.Stats.Skipped)
	}

	// Check individual result files
	resultsDir := filepath.Join(tmpDir, "results")
	for _, r := range results {
		resultPath := filepath.Join(resultsDir, r.ID+".json")
		if _, err := os.Stat(resultPath); os.IsNotExist(err) {
			t.Errorf("result file was not created: %s", resultPath)
		}
	}
}
