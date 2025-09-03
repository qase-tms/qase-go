package domain

// TestResultStatus represents the status of a test result
type TestResultStatus string

// StepStatus represents the status of a test step
type StepStatus string

// Test result status constants
const (
	StatusPassed     TestResultStatus = "passed"
	StatusFailed     TestResultStatus = "failed"
	StatusBlocked    TestResultStatus = "blocked"
	StatusSkipped    TestResultStatus = "skipped"
	StatusInProgress TestResultStatus = "in_progress"
	StatusInvalid    TestResultStatus = "invalid"
)

// Step status constants
const (
	StepStatusPassed  StepStatus = "passed"
	StepStatusFailed  StepStatus = "failed"
	StepStatusBlocked StepStatus = "blocked"
	StepStatusSkipped StepStatus = "skipped"
)

// TestResultStatus methods

// IsValid returns true if the test result status is valid
func (s TestResultStatus) IsValid() bool {
	switch s {
	case StatusPassed, StatusFailed, StatusBlocked, StatusSkipped, StatusInProgress, StatusInvalid:
		return true
	default:
		return false
	}
}

// String returns the string representation
func (s TestResultStatus) String() string {
	return string(s)
}

// IsFinal returns true if the status represents a final state
func (s TestResultStatus) IsFinal() bool {
	return s != StatusInProgress
}

// IsSuccess returns true if the status represents success
func (s TestResultStatus) IsSuccess() bool {
	return s == StatusPassed
}

// IsFailure returns true if the status represents failure
func (s TestResultStatus) IsFailure() bool {
	return s == StatusFailed || s == StatusBlocked || s == StatusInvalid
}

// StepStatus methods

// IsValid returns true if the step status is valid
func (s StepStatus) IsValid() bool {
	switch s {
	case StepStatusPassed, StepStatusFailed, StepStatusBlocked, StepStatusSkipped:
		return true
	default:
		return false
	}
}

// String returns the string representation
func (s StepStatus) String() string {
	return string(s)
}

// IsSuccess returns true if the step status represents success
func (s StepStatus) IsSuccess() bool {
	return s == StepStatusPassed
}

// IsFailure returns true if the step status represents failure
func (s StepStatus) IsFailure() bool {
	return s == StepStatusFailed || s == StepStatusBlocked
}