package domain

// StepType represents the type of test step
type StepType string

const (
	StepTypeText StepType = "text"
	// StepTypeGherkin StepType = "gherkin" // Not implemented as requested
)

// TestStep represents a test step matching JavaScript TestStepType
type TestStep struct {
	ID          string        `json:"id"`
	StepType    StepType      `json:"step_type"`
	Data        StepTextData  `json:"data"`
	ParentID    *string       `json:"parent_id"`
	Execution   StepExecution `json:"execution"`
	Attachments []Attachment  `json:"attachments"`
	Steps       []TestStep    `json:"steps"`
}

// StepTextData represents text step data matching JavaScript StepTextData
type StepTextData struct {
	Action         string  `json:"action"`
	ExpectedResult *string `json:"expected_result"`
	Data           *string `json:"data"`
}

// StepExecution represents step execution details matching JavaScript StepExecution
type StepExecution struct {
	StartTime  *int64     `json:"start_time"`
	Status     StepStatus `json:"status"`
	EndTime    *int64     `json:"end_time"`
	Duration   *int64     `json:"duration"`
	Stacktrace *string    `json:"stacktrace,omitempty"`
}

// NewTestStep creates a new test step with TEXT type
func NewTestStep(action string) *TestStep {
	return &TestStep{
		ID:       "",
		StepType: StepTypeText,
		Data: StepTextData{
			Action:         action,
			ExpectedResult: nil,
			Data:           nil,
		},
		ParentID: nil,
		Execution: StepExecution{
			Status:     StepStatusPassed,
			StartTime:  nil,
			EndTime:    nil,
			Duration:   nil,
			Stacktrace: nil,
		},
		Attachments: make([]Attachment, 0),
		Steps:       make([]TestStep, 0),
	}
}

// NewTestStepWithStatus creates a new test step with specified status
func NewTestStepWithStatus(action string, status StepStatus) *TestStep {
	step := NewTestStep(action)
	step.Execution.Status = status
	return step
}

// SetID sets the step ID
func (ts *TestStep) SetID(id string) {
	ts.ID = id
}

// SetParentID sets the parent step ID
func (ts *TestStep) SetParentID(parentID string) {
	ts.ParentID = &parentID
}

// SetAction sets the step action
func (ts *TestStep) SetAction(action string) {
	ts.Data.Action = action
}

// SetExpectedResult sets the expected result
func (ts *TestStep) SetExpectedResult(result string) {
	ts.Data.ExpectedResult = &result
}

// SetData sets the step data
func (ts *TestStep) SetData(data string) {
	ts.Data.Data = &data
}

// SetStatus sets the execution status
func (ts *TestStep) SetStatus(status StepStatus) {
	ts.Execution.Status = status
}

// SetTiming sets the start and end time
func (ts *TestStep) SetTiming(startTime, endTime int64) {
	ts.Execution.StartTime = &startTime
	ts.Execution.EndTime = &endTime
	if endTime > startTime {
		duration := endTime - startTime
		ts.Execution.Duration = &duration
	}
}

// SetDuration sets the execution time in milliseconds
func (ts *TestStep) SetDuration(ms int64) {
	ts.Execution.Duration = &ms
}

// AddAttachment adds an attachment
func (ts *TestStep) AddAttachment(attachment Attachment) {
	ts.Attachments = append(ts.Attachments, attachment)
}

// AddStep adds a nested step
func (ts *TestStep) AddStep(step TestStep) {
	ts.Steps = append(ts.Steps, step)
}

// GetStatus returns the execution status
func (ts *TestStep) GetStatus() StepStatus {
	return ts.Execution.Status
}

// GetAction returns the step action
func (ts *TestStep) GetAction() string {
	return ts.Data.Action
}

// IsSuccess returns true if the step passed
func (ts *TestStep) IsSuccess() bool {
	return ts.Execution.Status.IsSuccess()
}

// IsFailure returns true if the step failed
func (ts *TestStep) IsFailure() bool {
	return ts.Execution.Status.IsFailure()
}
