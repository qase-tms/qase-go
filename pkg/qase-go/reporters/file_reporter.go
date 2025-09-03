package reporters

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

// FileReporterConfig contains configuration for file reporter
type FileReporterConfig struct {
	// Config is the main configuration
	Config *config.Config

	// CustomPath overrides the path from config (optional)
	CustomPath string

	// CustomFilename overrides the filename (optional)
	CustomFilename string
}

// Report represents the Qase report structure according to the specification
type Report struct {
	QaseReport ReportData `json:"qase_report"`
}

// ReportData contains the main report data
type ReportData struct {
	Version     string       `json:"version"`
	Summary     Summary      `json:"summary"`
	Results     []Result     `json:"results"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// Summary contains test execution summary
type Summary struct {
	Status       string `json:"status"`
	StartTime    int64  `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	Duration     int64  `json:"duration"`
	TotalTests   int    `json:"total_tests"`
	PassedTests  int    `json:"passed_tests"`
	FailedTests  int    `json:"failed_tests"`
	SkippedTests int    `json:"skipped_tests"`
}

// Result represents a single test result in the report
type Result struct {
	ID          string            `json:"id,omitempty"`
	Title       string            `json:"title"`
	Signature   string            `json:"signature,omitempty"`
	RunID       *int64            `json:"run_id,omitempty"`
	TestopsID   interface{}       `json:"testops_id,omitempty"`
	Execution   Execution         `json:"execution"`
	Fields      map[string]string `json:"fields,omitempty"`
	Attachments []string          `json:"attachments,omitempty"`
	Steps       []Step            `json:"steps,omitempty"`
	Params      map[string]string `json:"params,omitempty"`
	GroupParams map[string]string `json:"group_params,omitempty"`
	Relations   *Relation         `json:"relations,omitempty"`
	Muted       bool              `json:"muted,omitempty"`
	Message     *string           `json:"message,omitempty"`
}

// Execution represents test execution details
type Execution struct {
	Status     string  `json:"status"`
	StartTime  *int64  `json:"start_time,omitempty"`
	EndTime    *int64  `json:"end_time,omitempty"`
	Duration   *int64  `json:"duration,omitempty"`
	Stacktrace *string `json:"stacktrace,omitempty"`
	Thread     *string `json:"thread,omitempty"`
}

// Step represents a test step
type Step struct {
	Data      StepData      `json:"data"`
	Execution StepExecution `json:"execution"`
	Steps     []Step        `json:"steps,omitempty"`
}

// StepData represents step data
type StepData struct {
	Action         string   `json:"action"`
	ExpectedResult *string  `json:"expected_result,omitempty"`
	InputData      *string  `json:"input_data,omitempty"`
	Attachments    []string `json:"attachments,omitempty"`
}

// StepExecution represents step execution details
type StepExecution struct {
	Status    string `json:"status"`
	StartTime *int64 `json:"start_time,omitempty"`
	EndTime   *int64 `json:"end_time,omitempty"`
	Duration  *int64 `json:"duration,omitempty"`
}

// Relation represents test relations
type Relation struct {
	Suite *Suite `json:"suite,omitempty"`
}

// Suite represents suite information
type Suite struct {
	Data []SuiteData `json:"data"`
}

// SuiteData represents suite data
type SuiteData struct {
	Title    string `json:"title"`
	PublicID *int64 `json:"public_id,omitempty"`
}

// Attachment represents an attachment
type Attachment struct {
	ID       string `json:"id"`
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
	Size     int64  `json:"size,omitempty"`
	URL      string `json:"url,omitempty"`
}

// FileReporter handles saving test results to files
type FileReporter struct {
	config    FileReporterConfig
	results   []*domain.TestResult
	startTime int64
	endTime   int64
}

// NewFileReporter creates a new file reporter with the given configuration
func NewFileReporter(config FileReporterConfig) *FileReporter {
	return &FileReporter{
		config:    config,
		results:   make([]*domain.TestResult, 0),
		startTime: time.Now().Unix(),
	}
}

// StartTestRun initializes the test run
func (r *FileReporter) StartTestRun(ctx context.Context) error {
	r.startTime = time.Now().Unix()
	return nil
}

// AddResult adds a test result to the internal collection
func (r *FileReporter) AddResult(result *domain.TestResult) error {
	if result == nil {
		return fmt.Errorf("result cannot be nil")
	}

	if result.Title == "" {
		return fmt.Errorf("result title cannot be empty")
	}

	r.results = append(r.results, result)
	return nil
}

// CompleteTestRun generates and saves the report file
func (r *FileReporter) CompleteTestRun(ctx context.Context) error {
	r.endTime = time.Now().Unix()

	// Get output path and filename from config
	outputPath := r.getOutputPath()
	filename := r.getFilename()
	format := r.getFormat()

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Generate report
	report := r.generateReport()

	// Save report to file
	filepath := filepath.Join(outputPath, filename+"."+format)

	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal report: %w", err)
	}

	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("failed to write report file: %w", err)
	}

	return nil
}

// generateReport creates the report structure from collected results
func (r *FileReporter) generateReport() *Report {
	summary := r.generateSummary()
	results := r.convertResults()
	attachments := r.collectAttachments()

	return &Report{
		QaseReport: ReportData{
			Version:     "1.0.0",
			Summary:     summary,
			Results:     results,
			Attachments: attachments,
		},
	}
}

// generateSummary creates the summary from collected results
func (r *FileReporter) generateSummary() Summary {
	summary := Summary{
		StartTime:    r.startTime,
		EndTime:      r.endTime,
		Duration:     r.endTime - r.startTime,
		TotalTests:   len(r.results),
		PassedTests:  0,
		FailedTests:  0,
		SkippedTests: 0,
	}

	// Count test results by status
	for _, result := range r.results {
		switch result.Execution.Status {
		case domain.StatusPassed:
			summary.PassedTests++
		case domain.StatusFailed:
			summary.FailedTests++
		case domain.StatusSkipped:
			summary.SkippedTests++
		}
	}

	// Determine overall status
	if summary.FailedTests > 0 {
		summary.Status = "failed"
	} else if summary.PassedTests > 0 {
		summary.Status = "passed"
	} else {
		summary.Status = "skipped"
	}

	return summary
}

// convertResults converts domain results to report results
func (r *FileReporter) convertResults() []Result {
	results := make([]Result, 0, len(r.results))

	for _, domainResult := range r.results {
		result := Result{
			ID:          domainResult.ID,
			Title:       domainResult.Title,
			Signature:   domainResult.Signature,
			RunID:       domainResult.RunID,
			TestopsID:   domainResult.TestopsID,
			Execution:   r.convertExecution(domainResult.Execution),
			Fields:      domainResult.Fields,
			Attachments: r.convertAttachments(domainResult.Attachments),
			Steps:       r.convertSteps(domainResult.Steps),
			Params:      domainResult.Params,
			GroupParams: domainResult.GroupParams,
			Relations:   r.convertRelations(domainResult.Relations),
			Muted:       domainResult.Muted,
			Message:     domainResult.Message,
		}
		results = append(results, result)
	}

	return results
}

// convertExecution converts domain execution to report execution
func (r *FileReporter) convertExecution(exec domain.TestExecution) Execution {
	return Execution{
		Status:     string(exec.Status),
		StartTime:  exec.StartTime,
		EndTime:    exec.EndTime,
		Duration:   exec.Duration,
		Stacktrace: exec.Stacktrace,
		Thread:     exec.Thread,
	}
}

// convertAttachments converts domain attachments to report attachments
func (r *FileReporter) convertAttachments(attachments []domain.Attachment) []string {
	ids := make([]string, 0, len(attachments))
	for _, attachment := range attachments {
		ids = append(ids, attachment.ID)
	}
	return ids
}

// convertSteps converts domain steps to report steps
func (r *FileReporter) convertSteps(steps []domain.TestStep) []Step {
	reportSteps := make([]Step, 0, len(steps))

	for _, step := range steps {
		reportStep := Step{
			Data:      r.convertStepData(step.Data),
			Execution: r.convertStepExecution(step.Execution),
			Steps:     r.convertSteps(step.Steps),
		}
		reportSteps = append(reportSteps, reportStep)
	}

	return reportSteps
}

// convertStepData converts domain step data to report step data
func (r *FileReporter) convertStepData(data domain.StepTextData) StepData {
	return StepData{
		Action:         data.Action,
		ExpectedResult: data.ExpectedResult,
		InputData:      data.Data,
		Attachments:    []string{}, // Steps don't have attachments in domain model
	}
}

// convertStepExecution converts domain step execution to report step execution
func (r *FileReporter) convertStepExecution(exec domain.StepExecution) StepExecution {
	return StepExecution{
		Status:    string(exec.Status),
		StartTime: exec.StartTime,
		EndTime:   exec.EndTime,
		Duration:  exec.Duration,
	}
}

// convertRelations converts domain relations to report relations
func (r *FileReporter) convertRelations(relations *domain.Relation) *Relation {
	if relations == nil {
		return nil
	}

	if relations.Suite == nil {
		return &Relation{}
	}

	suiteData := make([]SuiteData, 0, len(relations.Suite.Data))
	for _, data := range relations.Suite.Data {
		suiteData = append(suiteData, SuiteData{
			Title:    data.Title,
			PublicID: data.PublicID,
		})
	}

	return &Relation{
		Suite: &Suite{
			Data: suiteData,
		},
	}
}

// getOutputPath returns the output path from config
func (r *FileReporter) getOutputPath() string {
	if r.config.CustomPath != "" {
		return r.config.CustomPath
	}
	if r.config.Config != nil && r.config.Config.Report.Connection.Local.Path != "" {
		return r.config.Config.Report.Connection.Local.Path
	}
	return "./build/qase-report"
}

// getFilename returns the filename for the report
func (r *FileReporter) getFilename() string {
	if r.config.CustomFilename != "" {
		return r.config.CustomFilename
	}
	return "qase-report"
}

// getFormat returns the output format
func (r *FileReporter) getFormat() string {
	if r.config.Config != nil && r.config.Config.Report.Connection.Local.Format != "" {
		return r.config.Config.Report.Connection.Local.Format
	}
	return "json"
}

// collectAttachments collects all unique attachments from all results
func (r *FileReporter) collectAttachments() []Attachment {
	attachmentMap := make(map[string]Attachment)

	for _, result := range r.results {
		for _, domainAttachment := range result.Attachments {
			attachment := Attachment{
				ID:       domainAttachment.ID,
				FileName: domainAttachment.FileName,
				MimeType: domainAttachment.MimeType,
				Size:     domainAttachment.Size,
				// URL field is not available in domain.Attachment
			}
			attachmentMap[domainAttachment.ID] = attachment
		}
	}

	attachments := make([]Attachment, 0, len(attachmentMap))
	for _, attachment := range attachmentMap {
		attachments = append(attachments, attachment)
	}

	return attachments
}
