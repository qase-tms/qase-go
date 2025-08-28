package domain

import (
	"fmt"
	"strings"
)

// TestResult represents test execution result matching JavaScript TestResultType
type TestResult struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Signature   string            `json:"signature"`
	RunID       *int64            `json:"run_id"`
	TestopsID   interface{}       `json:"testops_id"` // number | number[] | null
	Execution   TestExecution     `json:"execution"`
	Fields      map[string]string `json:"fields"`
	Attachments []Attachment      `json:"attachments"`
	Steps       []TestStep        `json:"steps"`
	Params      map[string]string `json:"params"`
	GroupParams map[string]string `json:"group_params"`
	Relations   *Relation         `json:"relations"`
	Muted       bool              `json:"muted"`
	Message     *string           `json:"message"`
}

// TestExecution represents test execution details
type TestExecution struct {
	Status     TestResultStatus `json:"status"`
	StartTime  *int64           `json:"start_time,omitempty"` // Unix timestamp
	EndTime    *int64           `json:"end_time,omitempty"`   // Unix timestamp
	Duration   *int64           `json:"duration,omitempty"`   // Duration in milliseconds
	Stacktrace *string          `json:"stacktrace,omitempty"`
	Thread     *string          `json:"thread,omitempty"`
	// Additional fields for detailed error information
	ErrorDetails *ErrorDetails `json:"error_details,omitempty"`
}

// String implements the Stringer interface for TestExecution
func (te *TestExecution) String() string {
	var parts []string

	// Status
	parts = append(parts, fmt.Sprintf("Status: %q", te.Status))

	// Timing
	if te.StartTime != nil {
		parts = append(parts, fmt.Sprintf("StartTime: %d", *te.StartTime))
	}
	if te.EndTime != nil {
		parts = append(parts, fmt.Sprintf("EndTime: %d", *te.EndTime))
	}
	if te.Duration != nil {
		parts = append(parts, fmt.Sprintf("Duration: %d", *te.Duration))
	}

	// Stacktrace
	if te.Stacktrace != nil {
		parts = append(parts, fmt.Sprintf("Stacktrace: %q", *te.Stacktrace))
	}

	// Thread
	if te.Thread != nil {
		parts = append(parts, fmt.Sprintf("Thread: %q", *te.Thread))
	}

	// Error Details
	if te.ErrorDetails != nil {
		parts = append(parts, fmt.Sprintf("ErrorDetails: %s", te.ErrorDetails.String()))
	}

	return fmt.Sprintf("TestExecution{%s}", strings.Join(parts, ", "))
}

// ErrorDetails represents detailed error information from test assertions
type ErrorDetails struct {
	ErrorType    string `json:"error_type,omitempty"`    // Type of assertion error (e.g., "Should be true", "Not equal")
	Expected     string `json:"expected,omitempty"`      // Expected value
	Actual       string `json:"actual,omitempty"`        // Actual value
	ErrorMessage string `json:"error_message,omitempty"` // Detailed error message
	File         string `json:"file,omitempty"`          // File where error occurred
	Line         int    `json:"line,omitempty"`          // Line number where error occurred
}

// String implements the Stringer interface for ErrorDetails
func (ed *ErrorDetails) String() string {
	var parts []string

	if ed.ErrorType != "" {
		parts = append(parts, fmt.Sprintf("ErrorType: %q", ed.ErrorType))
	}
	if ed.Expected != "" {
		parts = append(parts, fmt.Sprintf("Expected: %q", ed.Expected))
	}
	if ed.Actual != "" {
		parts = append(parts, fmt.Sprintf("Actual: %q", ed.Actual))
	}
	if ed.ErrorMessage != "" {
		parts = append(parts, fmt.Sprintf("ErrorMessage: %q", ed.ErrorMessage))
	}
	if ed.File != "" {
		parts = append(parts, fmt.Sprintf("File: %q", ed.File))
	}
	if ed.Line != 0 {
		parts = append(parts, fmt.Sprintf("Line: %d", ed.Line))
	}

	return fmt.Sprintf("ErrorDetails{%s}", strings.Join(parts, ", "))
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
	PublicID *int64 `json:"public_id"`
}

// String implements the Stringer interface for SuiteData
func (sd *SuiteData) String() string {
	var parts []string

	if sd.Title != "" {
		parts = append(parts, fmt.Sprintf("Title: %q", sd.Title))
	}
	if sd.PublicID != nil {
		parts = append(parts, fmt.Sprintf("PublicID: %d", *sd.PublicID))
	}

	return fmt.Sprintf("SuiteData{%s}", strings.Join(parts, ", "))
}

// String implements the Stringer interface for Suite
func (s *Suite) String() string {
	if len(s.Data) == 0 {
		return "Suite{}"
	}

	dataStr := make([]string, 0, len(s.Data))
	for i, data := range s.Data {
		dataStr = append(dataStr, fmt.Sprintf("[%d]: %s", i, data.String()))
	}

	return fmt.Sprintf("Suite{Data: [%s]}", strings.Join(dataStr, ", "))
}

// String implements the Stringer interface for Relation
func (r *Relation) String() string {
	if r.Suite == nil {
		return "Relation{}"
	}
	return fmt.Sprintf("Relation{Suite: %s}", r.Suite.String())
}

// NewTestResult creates a new test result with the given title
func NewTestResult(title string) *TestResult {
	return &TestResult{
		ID:          "",
		Title:       title,
		Signature:   "",
		RunID:       nil,
		TestopsID:   nil,
		Execution:   TestExecution{},
		Fields:      make(map[string]string),
		Attachments: make([]Attachment, 0),
		Steps:       make([]TestStep, 0),
		Params:      make(map[string]string),
		GroupParams: make(map[string]string),
		Relations:   nil,
		Muted:       false,
		Message:     nil,
	}
}

// SetRunID sets the run ID
func (tr *TestResult) SetRunID(runID int64) {
	tr.RunID = &runID
}

// SetTestopsID sets the testops ID (can be int64 or []int64)
func (tr *TestResult) SetTestopsID(id interface{}) {
	tr.TestopsID = id
}

// SetTestopsIDSingle sets a single testops ID
func (tr *TestResult) SetTestopsIDSingle(id int64) {
	tr.TestopsID = id
}

// SetTestopsIDMultiple sets multiple testops IDs
func (tr *TestResult) SetTestopsIDMultiple(ids []int64) {
	tr.TestopsID = ids
}

// SetMessage sets the message
func (tr *TestResult) SetMessage(message string) {
	tr.Message = &message
}

// SetField sets a field value
func (tr *TestResult) SetField(key, value string) {
	tr.Fields[key] = value
}

// SetParam sets a parameter value
func (tr *TestResult) SetParam(key, value string) {
	tr.Params[key] = value
}

// SetGroupParam sets a group parameter value
func (tr *TestResult) SetGroupParam(key, value string) {
	tr.GroupParams[key] = value
}

// AddAttachment adds an attachment
func (tr *TestResult) AddAttachment(attachment Attachment) {
	tr.Attachments = append(tr.Attachments, attachment)
}

// AddStep adds a test step
func (tr *TestResult) AddStep(step TestStep) {
	tr.Steps = append(tr.Steps, step)
}

// SetSuite sets the suite relation
func (tr *TestResult) SetSuite(suiteData []SuiteData) {
	if tr.Relations == nil {
		tr.Relations = &Relation{}
	}
	tr.Relations.Suite = &Suite{Data: suiteData}
}

// AddSuiteData adds suite data to relations
func (tr *TestResult) AddSuiteData(title string, publicID *int64) {
	if tr.Relations == nil {
		tr.Relations = &Relation{}
	}
	if tr.Relations.Suite == nil {
		tr.Relations.Suite = &Suite{Data: make([]SuiteData, 0)}
	}
	tr.Relations.Suite.Data = append(tr.Relations.Suite.Data, SuiteData{
		Title:    title,
		PublicID: publicID,
	})
}

// SetErrorDetails sets the error details
func (tr *TestResult) SetErrorDetails(errorDetails *ErrorDetails) {
	tr.Execution.ErrorDetails = errorDetails
}

// SetErrorDetailsFromString creates and sets error details from a simple error message
func (tr *TestResult) SetErrorDetailsFromString(errorType, errorMessage string) {
	tr.Execution.ErrorDetails = &ErrorDetails{
		ErrorType:    errorType,
		ErrorMessage: errorMessage,
	}
}

// SetErrorDetailsWithValues creates and sets error details with expected/actual values
func (tr *TestResult) SetErrorDetailsWithValues(errorType, expected, actual, errorMessage string) {
	tr.Execution.ErrorDetails = &ErrorDetails{
		ErrorType:    errorType,
		Expected:     expected,
		Actual:       actual,
		ErrorMessage: errorMessage,
	}
}

// SetErrorLocation sets the file and line information for the error
func (tr *TestResult) SetErrorLocation(file string, line int) {
	if tr.Execution.ErrorDetails == nil {
		tr.Execution.ErrorDetails = &ErrorDetails{}
	}
	tr.Execution.ErrorDetails.File = file
	tr.Execution.ErrorDetails.Line = line
}

// String implements the Stringer interface for better logging
func (tr *TestResult) String() string {
	var parts []string

	// Basic info
	if tr.ID != "" {
		parts = append(parts, fmt.Sprintf("ID: %q", tr.ID))
	}
	if tr.Title != "" {
		parts = append(parts, fmt.Sprintf("Title: %q", tr.Title))
	}
	if tr.Signature != "" {
		parts = append(parts, fmt.Sprintf("Signature: %q", tr.Signature))
	}

	// Run ID
	if tr.RunID != nil {
		parts = append(parts, fmt.Sprintf("RunID: %d", *tr.RunID))
	}

	// Testops ID
	if tr.TestopsID != nil {
		switch v := tr.TestopsID.(type) {
		case int64:
			parts = append(parts, fmt.Sprintf("TestopsID: %d", v))
		case []int64:
			parts = append(parts, fmt.Sprintf("TestopsID: %v", v))
		default:
			parts = append(parts, fmt.Sprintf("TestopsID: %v", v))
		}
	}

	// Execution
	parts = append(parts, fmt.Sprintf("Execution: %s", tr.Execution.String()))

	// Fields
	if len(tr.Fields) > 0 {
		fieldsStr := make([]string, 0, len(tr.Fields))
		for k, v := range tr.Fields {
			fieldsStr = append(fieldsStr, fmt.Sprintf("%q: %q", k, v))
		}
		parts = append(parts, fmt.Sprintf("Fields: {%s}", strings.Join(fieldsStr, ", ")))
	}

	// Attachments
	if len(tr.Attachments) > 0 {
		attachmentsStr := make([]string, 0, len(tr.Attachments))
		for i, att := range tr.Attachments {
			attachmentsStr = append(attachmentsStr, fmt.Sprintf("[%d]: %s", i, att.String()))
		}
		parts = append(parts, fmt.Sprintf("Attachments: [%s]", strings.Join(attachmentsStr, ", ")))
	}

	// Steps
	if len(tr.Steps) > 0 {
		stepsStr := make([]string, 0, len(tr.Steps))
		for i, step := range tr.Steps {
			stepsStr = append(stepsStr, fmt.Sprintf("[%d]: %s", i, step.String()))
		}
		parts = append(parts, fmt.Sprintf("Steps: [%s]", strings.Join(stepsStr, ", ")))
	}

	// Params
	if len(tr.Params) > 0 {
		paramsStr := make([]string, 0, len(tr.Params))
		for k, v := range tr.Params {
			paramsStr = append(paramsStr, fmt.Sprintf("%q: %q", k, v))
		}
		parts = append(parts, fmt.Sprintf("Params: {%s}", strings.Join(paramsStr, ", ")))
	}

	// Group Params
	if len(tr.GroupParams) > 0 {
		groupParamsStr := make([]string, 0, len(tr.GroupParams))
		for k, v := range tr.GroupParams {
			groupParamsStr = append(groupParamsStr, fmt.Sprintf("%q: %q", k, v))
		}
		parts = append(parts, fmt.Sprintf("GroupParams: {%s}", strings.Join(groupParamsStr, ", ")))
	}

	// Relations
	if tr.Relations != nil {
		parts = append(parts, fmt.Sprintf("Relations: %s", tr.Relations.String()))
	}

	// Muted
	parts = append(parts, fmt.Sprintf("Muted: %t", tr.Muted))

	// Message
	if tr.Message != nil {
		parts = append(parts, fmt.Sprintf("Message: %q", *tr.Message))
	}

	return fmt.Sprintf("TestResult{%s}", strings.Join(parts, ", "))
}
