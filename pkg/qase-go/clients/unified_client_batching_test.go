package clients

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
	"github.com/qase-tms/qase-go/pkg/qase-go/logging"
)

// recordingSender records every request made by UnifiedClient instead of
// talking to the API, so batching can be asserted on.
type recordingSender struct {
	calls    [][]*domain.TestResult
	failCall func(callIndex int) error
}

func (s *recordingSender) SendResult(ctx context.Context, projectCode string, runID int64, result *domain.TestResult) error {
	return s.SendResults(ctx, projectCode, runID, []*domain.TestResult{result})
}

func (s *recordingSender) SendResults(ctx context.Context, projectCode string, runID int64, results []*domain.TestResult) error {
	batch := make([]*domain.TestResult, len(results))
	copy(batch, results)
	s.calls = append(s.calls, batch)

	if s.failCall != nil {
		return s.failCall(len(s.calls) - 1)
	}

	return nil
}

func (s *recordingSender) batchSizes() []int {
	sizes := make([]int, 0, len(s.calls))
	for _, call := range s.calls {
		sizes = append(sizes, len(call))
	}
	return sizes
}

func newBatchingTestClient(sender *recordingSender, batchSize int) *UnifiedClient {
	cfg := config.NewConfig()
	cfg.TestOps.Project = "TEST"
	cfg.TestOps.Batch.Size = batchSize

	return &UnifiedClient{
		v2Client:    sender,
		config:      cfg,
		projectCode: cfg.TestOps.Project,
	}
}

func makeResults(count int) []*domain.TestResult {
	results := make([]*domain.TestResult, 0, count)
	for i := 0; i < count; i++ {
		results = append(results, domain.NewTestResult(fmt.Sprintf("Test %d", i)))
	}
	return results
}

// assertEveryResultSentOnce verifies that the recorded requests together carry
// exactly the given results, each of them exactly once.
func assertEveryResultSentOnce(t *testing.T, sender *recordingSender, expected []*domain.TestResult) {
	t.Helper()

	seen := make(map[string]int, len(expected))
	total := 0
	for _, call := range sender.calls {
		for _, result := range call {
			seen[result.Title]++
			total++
		}
	}

	if total != len(expected) {
		t.Errorf("sent %d results in total, want %d", total, len(expected))
	}

	for _, result := range expected {
		switch seen[result.Title] {
		case 1:
		case 0:
			t.Errorf("result %q was never sent", result.Title)
		default:
			t.Errorf("result %q was sent %d times, want exactly once", result.Title, seen[result.Title])
		}
	}
}

func TestUnifiedClient_UploadResults_SplitsResultsIntoBatches(t *testing.T) {
	sender := &recordingSender{}
	client := newBatchingTestClient(sender, 200)
	results := makeResults(250)

	if err := client.UploadResults(context.Background(), 123, results); err != nil {
		t.Fatalf("UploadResults() returned an unexpected error: %v", err)
	}

	if len(sender.calls) != 2 {
		t.Errorf("250 results with batch size 200 produced %d requests %v, want 2", len(sender.calls), sender.batchSizes())
	}

	assertEveryResultSentOnce(t, sender, results)
}

func TestUnifiedClient_UploadResults_ExactBatchSizeIsOneRequest(t *testing.T) {
	sender := &recordingSender{}
	client := newBatchingTestClient(sender, 200)
	results := makeResults(200)

	if err := client.UploadResults(context.Background(), 123, results); err != nil {
		t.Fatalf("UploadResults() returned an unexpected error: %v", err)
	}

	if len(sender.calls) != 1 {
		t.Errorf("200 results with batch size 200 produced %d requests %v, want 1", len(sender.calls), sender.batchSizes())
	}
}

func TestUnifiedClient_UploadResults_ClampsConfiguredBatchSize(t *testing.T) {
	sender := &recordingSender{}
	client := newBatchingTestClient(sender, 2000)
	results := makeResults(250)

	if err := client.UploadResults(context.Background(), 123, results); err != nil {
		t.Fatalf("UploadResults() returned an unexpected error: %v", err)
	}

	for i, size := range sender.batchSizes() {
		if size > config.MaxBatchSize {
			t.Errorf("request %d carried %d results, want at most %d", i, size, config.MaxBatchSize)
		}
	}

	if len(sender.calls) != 2 {
		t.Errorf("250 results with batch size 2000 produced %d requests %v, want 2 after clamping", len(sender.calls), sender.batchSizes())
	}

	assertEveryResultSentOnce(t, sender, results)
}

func TestUnifiedClient_UploadResults_ReportsNumberOfLostResults(t *testing.T) {
	sender := &recordingSender{
		failCall: func(callIndex int) error {
			if callIndex == 1 {
				return fmt.Errorf("boom")
			}
			return nil
		},
	}
	client := newBatchingTestClient(sender, 100)
	results := makeResults(250)

	err := client.UploadResults(context.Background(), 123, results)
	if err == nil {
		t.Fatal("UploadResults() returned nil when a batch failed, want an error reporting the lost results")
	}

	if !strings.Contains(err.Error(), "100") {
		t.Errorf("error %q does not report the number of lost results (100)", err.Error())
	}

	if !strings.Contains(err.Error(), "250") {
		t.Errorf("error %q does not report the total number of results (250)", err.Error())
	}
}

func TestUnifiedClient_UploadResults_KeepsFailedResultsIntact(t *testing.T) {
	sender := &recordingSender{
		failCall: func(callIndex int) error {
			return fmt.Errorf("boom")
		},
	}
	client := newBatchingTestClient(sender, 100)
	results := makeResults(250)

	_ = client.UploadResults(context.Background(), 123, results)

	if len(results) != 250 {
		t.Fatalf("caller slice holds %d results after a failed upload, want 250", len(results))
	}

	for i, result := range results {
		if result == nil {
			t.Fatalf("result %d was discarded from the caller slice after a failed upload", i)
		}
	}
}

// captureLogs installs a file-only global logger and returns everything it
// recorded during fn.
func captureLogs(t *testing.T, fn func()) string {
	t.Helper()

	logDir := t.TempDir()
	logger, err := logging.NewLogger(logging.LoggerConfig{
		LogToFile:   true,
		LogDir:      logDir,
		LogFileName: "capture.log",
		LogLevel:    logging.DEBUG,
	})
	if err != nil {
		t.Fatalf("failed to create logger: %v", err)
	}

	logging.SetGlobalLogger(logger)
	defer logging.SetGlobalLogger(nil)

	fn()

	content, err := os.ReadFile(filepath.Join(logDir, "capture.log"))
	if err != nil {
		t.Fatalf("failed to read captured logs: %v", err)
	}

	return string(content)
}

func TestUnifiedClient_UploadResults_LogsClampedBatchSize(t *testing.T) {
	sender := &recordingSender{}
	client := newBatchingTestClient(sender, 2000)

	logs := captureLogs(t, func() {
		if err := client.UploadResults(context.Background(), 123, makeResults(250)); err != nil {
			t.Errorf("UploadResults() returned an unexpected error: %v", err)
		}
	})

	if !strings.Contains(logs, "2000") || !strings.Contains(logs, "200") {
		t.Errorf("clamping of batch size 2000 down to 200 was not logged, logs were:\n%s", logs)
	}
}

func TestUnifiedClient_UploadResults_LogsNumberOfLostResults(t *testing.T) {
	sender := &recordingSender{
		failCall: func(callIndex int) error { return fmt.Errorf("boom") },
	}
	client := newBatchingTestClient(sender, 100)

	logs := captureLogs(t, func() {
		if err := client.UploadResults(context.Background(), 123, makeResults(250)); err == nil {
			t.Error("UploadResults() returned nil when every batch failed")
		}
	})

	if !strings.Contains(logs, "250 results uploaded, 250 lost") {
		t.Errorf("the number of lost results was not logged, logs were:\n%s", logs)
	}
}
