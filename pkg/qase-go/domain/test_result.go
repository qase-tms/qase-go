package domain

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

// ErrorDetails represents detailed error information from test assertions
type ErrorDetails struct {
	ErrorType    string `json:"error_type,omitempty"`    // Type of assertion error (e.g., "Should be true", "Not equal")
	Expected     string `json:"expected,omitempty"`      // Expected value
	Actual       string `json:"actual,omitempty"`        // Actual value
	ErrorMessage string `json:"error_message,omitempty"` // Detailed error message
	File         string `json:"file,omitempty"`          // File where error occurred
	Line         int    `json:"line,omitempty"`          // Line number where error occurred
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
