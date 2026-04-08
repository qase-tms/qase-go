package reporters

import (
	"context"
	"crypto/rand"
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

// RunReport represents the run.json structure per Qase report specification
type RunReport struct {
	Title       string              `json:"title"`
	Execution   RunExecution        `json:"execution"`
	Stats       RunStats            `json:"stats"`
	Results     []RunResultSummary  `json:"results"`
	Threads     []string            `json:"threads"`
	Suites      []string            `json:"suites"`
	Environment *string             `json:"environment"`
	HostData    map[string]string   `json:"host_data,omitempty"`
}

// RunExecution represents execution timing in run.json
type RunExecution struct {
	StartTime          int64 `json:"start_time"`
	EndTime            int64 `json:"end_time"`
	Duration           int64 `json:"duration"`
	CumulativeDuration int64 `json:"cumulative_duration"`
}

// RunStats represents test statistics in run.json
type RunStats struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Failed  int `json:"failed"`
	Skipped int `json:"skipped"`
	Blocked int `json:"blocked"`
	Invalid int `json:"invalid"`
	Muted   int `json:"muted"`
}

// RunResultSummary is a lightweight result entry in run.json
type RunResultSummary struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Status   string  `json:"status"`
	Duration int64   `json:"duration"`
	Thread   *string `json:"thread"`
}

// Result represents a single test result file in results/ directory
type Result struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Signature   *string           `json:"signature"`
	TestopsIDs  []int64           `json:"testops_ids"`
	Execution   Execution         `json:"execution"`
	Fields      map[string]string `json:"fields"`
	Attachments []Attachment      `json:"attachments"`
	Steps       []Step            `json:"steps"`
	Params      map[string]string `json:"params"`
	ParamGroups [][]string        `json:"param_groups"`
	Relations   *Relation         `json:"relations"`
	Tags        []string          `json:"tags"`
	Muted       bool              `json:"muted"`
	Message     *string           `json:"message"`
}

// Execution represents test execution details
type Execution struct {
	Status     string  `json:"status"`
	StartTime  *int64  `json:"start_time"`
	EndTime    *int64  `json:"end_time"`
	Duration   *int64  `json:"duration"`
	Stacktrace *string `json:"stacktrace"`
	Thread     *string `json:"thread"`
}

// Step represents a test step
type Step struct {
	ID        string        `json:"id"`
	StepType  string        `json:"step_type"`
	Data      StepData      `json:"data"`
	ParentID  *string       `json:"parent_id"`
	Execution StepExecution `json:"execution"`
	Steps     []Step        `json:"steps"`
}

// StepData represents step data
type StepData struct {
	Action         string  `json:"action"`
	ExpectedResult *string `json:"expected_result"`
	InputData      *string `json:"input_data"`
}

// StepExecution represents step execution details
type StepExecution struct {
	Status      string       `json:"status"`
	StartTime   *int64       `json:"start_time"`
	EndTime     *int64       `json:"end_time"`
	Duration    *int64       `json:"duration"`
	Attachments []Attachment `json:"attachments"`
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

// Attachment represents an attachment in a result
type Attachment struct {
	ID       string  `json:"id"`
	FileName *string `json:"file_name"`
	FilePath *string `json:"file_path"`
	MimeType *string `json:"mime_type"`
	Content  *string `json:"content"`
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
		startTime: time.Now().UnixMilli(),
	}
}

// StartTestRun initializes the test run
func (r *FileReporter) StartTestRun(ctx context.Context) error {
	r.startTime = time.Now().UnixMilli()
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

// CompleteTestRun generates run.json and individual result files
func (r *FileReporter) CompleteTestRun(ctx context.Context) error {
	r.endTime = time.Now().UnixMilli()

	outputPath := r.getOutputPath()
	resultsPath := filepath.Join(outputPath, "results")

	// Create output directories
	if err := os.MkdirAll(resultsPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Write individual result files
	for _, domainResult := range r.results {
		result := r.convertResult(domainResult)
		data, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal result %s: %w", domainResult.ID, err)
		}

		resultFile := filepath.Join(resultsPath, domainResult.ID+".json")
		if err := os.WriteFile(resultFile, data, 0644); err != nil {
			return fmt.Errorf("failed to write result file %s: %w", resultFile, err)
		}
	}

	// Write run.json (merging with existing if present)
	runFile := filepath.Join(outputPath, "run.json")
	runReport := r.generateRunReport()

	// Try to read and merge with existing run.json
	if existingData, err := os.ReadFile(runFile); err == nil {
		var existing RunReport
		if err := json.Unmarshal(existingData, &existing); err == nil {
			r.mergeRunReports(&existing, runReport)
			runReport = &existing
		}
	}

	data, err := json.MarshalIndent(runReport, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal run report: %w", err)
	}

	if err := os.WriteFile(runFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write run.json: %w", err)
	}

	return nil
}

// mergeRunReports merges new run report data into an existing one
func (r *FileReporter) mergeRunReports(existing *RunReport, new *RunReport) {
	// Merge stats
	existing.Stats.Total += new.Stats.Total
	existing.Stats.Passed += new.Stats.Passed
	existing.Stats.Failed += new.Stats.Failed
	existing.Stats.Skipped += new.Stats.Skipped
	existing.Stats.Blocked += new.Stats.Blocked
	existing.Stats.Invalid += new.Stats.Invalid
	existing.Stats.Muted += new.Stats.Muted

	// Use earliest start_time and latest end_time
	if new.Execution.StartTime < existing.Execution.StartTime {
		existing.Execution.StartTime = new.Execution.StartTime
	}
	if new.Execution.EndTime > existing.Execution.EndTime {
		existing.Execution.EndTime = new.Execution.EndTime
	}
	existing.Execution.Duration = existing.Execution.EndTime - existing.Execution.StartTime
	existing.Execution.CumulativeDuration += new.Execution.CumulativeDuration

	// Append results
	existing.Results = append(existing.Results, new.Results...)

	// Merge threads (unique)
	threadSet := make(map[string]bool)
	for _, t := range existing.Threads {
		threadSet[t] = true
	}
	for _, t := range new.Threads {
		if !threadSet[t] {
			existing.Threads = append(existing.Threads, t)
		}
	}

	// Merge suites (unique)
	suiteSet := make(map[string]bool)
	for _, s := range existing.Suites {
		suiteSet[s] = true
	}
	for _, s := range new.Suites {
		if !suiteSet[s] {
			existing.Suites = append(existing.Suites, s)
		}
	}
}

// generateRunReport creates the run.json structure
func (r *FileReporter) generateRunReport() *RunReport {
	stats := RunStats{}
	var cumulativeDuration int64
	resultSummaries := make([]RunResultSummary, 0, len(r.results))
	threadSet := make(map[string]bool)
	suiteSet := make(map[string]bool)

	for _, result := range r.results {
		switch result.Execution.Status {
		case domain.StatusPassed:
			stats.Passed++
		case domain.StatusFailed:
			stats.Failed++
		case domain.StatusSkipped:
			stats.Skipped++
		case domain.StatusBlocked:
			stats.Blocked++
		case domain.StatusInvalid:
			stats.Invalid++
		}
		if result.Muted {
			stats.Muted++
		}
		stats.Total++

		var duration int64
		if result.Execution.Duration != nil {
			duration = *result.Execution.Duration
			cumulativeDuration += duration
		}

		summary := RunResultSummary{
			ID:       result.ID,
			Title:    result.Title,
			Status:   string(result.Execution.Status),
			Duration: duration,
			Thread:   result.Execution.Thread,
		}
		resultSummaries = append(resultSummaries, summary)

		if result.Execution.Thread != nil {
			threadSet[*result.Execution.Thread] = true
		}

		if result.Relations != nil && result.Relations.Suite != nil {
			for _, s := range result.Relations.Suite.Data {
				suiteSet[s.Title] = true
			}
		}
	}

	threads := make([]string, 0, len(threadSet))
	for t := range threadSet {
		threads = append(threads, t)
	}

	suites := make([]string, 0, len(suiteSet))
	for s := range suiteSet {
		suites = append(suites, s)
	}

	return &RunReport{
		Title: "Qase Report",
		Execution: RunExecution{
			StartTime:          r.startTime,
			EndTime:            r.endTime,
			Duration:           r.endTime - r.startTime,
			CumulativeDuration: cumulativeDuration,
		},
		Stats:       stats,
		Results:     resultSummaries,
		Threads:     threads,
		Suites:      suites,
		Environment: nil,
	}
}

// generateID creates a random hex ID
func generateID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// convertResult converts a domain result to the spec-compliant result format
func (r *FileReporter) convertResult(domainResult *domain.TestResult) *Result {
	// Ensure result has an ID
	if domainResult.ID == "" {
		domainResult.ID = generateID()
	}

	var signature *string
	if domainResult.Signature != "" {
		sig := domainResult.Signature
		signature = &sig
	}

	testopsIDs := r.extractTestopsIDs(domainResult.TestopsID)

	fields := domainResult.Fields
	if fields == nil {
		fields = map[string]string{}
	}

	params := domainResult.Params
	if params == nil {
		params = map[string]string{}
	}

	return &Result{
		ID:          domainResult.ID,
		Title:       domainResult.Title,
		Signature:   signature,
		TestopsIDs:  testopsIDs,
		Execution:   r.convertExecution(domainResult.Execution),
		Fields:      fields,
		Attachments: r.convertAttachments(domainResult.Attachments),
		Steps:       r.convertSteps(domainResult.Steps, nil),
		Params:      params,
		ParamGroups: nil,
		Relations:   r.convertRelations(domainResult.Relations),
		Tags:        domainResult.Tags,
		Muted:       domainResult.Muted,
		Message:     domainResult.Message,
	}
}

// extractTestopsIDs extracts testops IDs from the flexible TestopsID field
func (r *FileReporter) extractTestopsIDs(testopsID interface{}) []int64 {
	if testopsID == nil {
		return nil
	}

	switch v := testopsID.(type) {
	case int64:
		return []int64{v}
	case []int64:
		return v
	case float64:
		return []int64{int64(v)}
	default:
		return nil
	}
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

// convertAttachments converts domain attachments to report attachment objects
func (r *FileReporter) convertAttachments(attachments []domain.Attachment) []Attachment {
	result := make([]Attachment, 0, len(attachments))
	for _, a := range attachments {
		att := Attachment{
			ID: a.ID,
		}
		if a.FileName != "" {
			fn := a.FileName
			att.FileName = &fn
		}
		if a.HasFilePath() {
			fp := a.GetFilePath()
			att.FilePath = &fp
		}
		if a.MimeType != "" {
			mt := a.MimeType
			att.MimeType = &mt
		}
		if len(a.Content) > 0 {
			c := string(a.Content)
			att.Content = &c
		}
		result = append(result, att)
	}
	return result
}

// convertSteps converts domain steps to report steps
func (r *FileReporter) convertSteps(steps []domain.TestStep, parentID *string) []Step {
	reportSteps := make([]Step, 0, len(steps))

	for _, step := range steps {
		reportStep := Step{
			ID:        step.ID,
			StepType:  string(step.StepType),
			Data:      r.convertStepData(step.Data),
			ParentID:  parentID,
			Execution: r.convertStepExecution(step),
			Steps:     r.convertSteps(step.Steps, &step.ID),
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
	}
}

// convertStepExecution converts domain step execution to report step execution
func (r *FileReporter) convertStepExecution(step domain.TestStep) StepExecution {
	return StepExecution{
		Status:      string(step.Execution.Status),
		StartTime:   step.Execution.StartTime,
		EndTime:     step.Execution.EndTime,
		Duration:    step.Execution.Duration,
		Attachments: r.convertAttachments(step.Attachments),
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
