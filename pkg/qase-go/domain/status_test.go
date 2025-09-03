package domain

import "testing"

// TestResultStatus tests

func TestTestResultStatus_IsValid(t *testing.T) {
	tests := []struct {
		name   string
		status TestResultStatus
		want   bool
	}{
		{"passed is valid", StatusPassed, true},
		{"failed is valid", StatusFailed, true},
		{"blocked is valid", StatusBlocked, true},
		{"skipped is valid", StatusSkipped, true},
		{"in_progress is valid", StatusInProgress, true},
		{"invalid is valid", StatusInvalid, true},
		{"empty string is invalid", TestResultStatus(""), false},
		{"unknown status is invalid", TestResultStatus("unknown"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.IsValid(); got != tt.want {
				t.Errorf("TestResultStatus.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestResultStatus_String(t *testing.T) {
	tests := []struct {
		name   string
		status TestResultStatus
		want   string
	}{
		{"passed", StatusPassed, "passed"},
		{"failed", StatusFailed, "failed"},
		{"blocked", StatusBlocked, "blocked"},
		{"skipped", StatusSkipped, "skipped"},
		{"in_progress", StatusInProgress, "in_progress"},
		{"invalid", StatusInvalid, "invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.String(); got != tt.want {
				t.Errorf("TestResultStatus.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestResultStatus_IsFinal(t *testing.T) {
	tests := []struct {
		name   string
		status TestResultStatus
		want   bool
	}{
		{"passed is final", StatusPassed, true},
		{"failed is final", StatusFailed, true},
		{"blocked is final", StatusBlocked, true},
		{"skipped is final", StatusSkipped, true},
		{"invalid is final", StatusInvalid, true},
		{"in_progress is not final", StatusInProgress, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.IsFinal(); got != tt.want {
				t.Errorf("TestResultStatus.IsFinal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestResultStatus_IsSuccess(t *testing.T) {
	tests := []struct {
		name   string
		status TestResultStatus
		want   bool
	}{
		{"passed is success", StatusPassed, true},
		{"failed is not success", StatusFailed, false},
		{"blocked is not success", StatusBlocked, false},
		{"skipped is not success", StatusSkipped, false},
		{"in_progress is not success", StatusInProgress, false},
		{"invalid is not success", StatusInvalid, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.IsSuccess(); got != tt.want {
				t.Errorf("TestResultStatus.IsSuccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestResultStatus_IsFailure(t *testing.T) {
	tests := []struct {
		name   string
		status TestResultStatus
		want   bool
	}{
		{"passed is not failure", StatusPassed, false},
		{"failed is failure", StatusFailed, true},
		{"blocked is failure", StatusBlocked, true},
		{"skipped is not failure", StatusSkipped, false},
		{"in_progress is not failure", StatusInProgress, false},
		{"invalid is failure", StatusInvalid, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.IsFailure(); got != tt.want {
				t.Errorf("TestResultStatus.IsFailure() = %v, want %v", got, tt.want)
			}
		})
	}
}

// StepStatus tests

func TestStepStatus_IsValid(t *testing.T) {
	tests := []struct {
		name   string
		status StepStatus
		want   bool
	}{
		{"passed is valid", StepStatusPassed, true},
		{"failed is valid", StepStatusFailed, true},
		{"blocked is valid", StepStatusBlocked, true},
		{"skipped is valid", StepStatusSkipped, true},
		{"empty string is invalid", StepStatus(""), false},
		{"unknown status is invalid", StepStatus("unknown"), false},
		{"in_progress is invalid for steps", StepStatus("in_progress"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.IsValid(); got != tt.want {
				t.Errorf("StepStatus.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStepStatus_String(t *testing.T) {
	tests := []struct {
		name   string
		status StepStatus
		want   string
	}{
		{"passed", StepStatusPassed, "passed"},
		{"failed", StepStatusFailed, "failed"},
		{"blocked", StepStatusBlocked, "blocked"},
		{"skipped", StepStatusSkipped, "skipped"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.String(); got != tt.want {
				t.Errorf("StepStatus.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStepStatus_IsSuccess(t *testing.T) {
	tests := []struct {
		name   string
		status StepStatus
		want   bool
	}{
		{"passed is success", StepStatusPassed, true},
		{"failed is not success", StepStatusFailed, false},
		{"blocked is not success", StepStatusBlocked, false},
		{"skipped is not success", StepStatusSkipped, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.IsSuccess(); got != tt.want {
				t.Errorf("StepStatus.IsSuccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStepStatus_IsFailure(t *testing.T) {
	tests := []struct {
		name   string
		status StepStatus
		want   bool
	}{
		{"passed is not failure", StepStatusPassed, false},
		{"failed is failure", StepStatusFailed, true},
		{"blocked is failure", StepStatusBlocked, true},
		{"skipped is not failure", StepStatusSkipped, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.IsFailure(); got != tt.want {
				t.Errorf("StepStatus.IsFailure() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test constants are properly defined
func TestStatusConstants(t *testing.T) {
	// Test TestResultStatus constants
	if StatusPassed != "passed" {
		t.Errorf("StatusPassed = %v, want 'passed'", StatusPassed)
	}
	if StatusFailed != "failed" {
		t.Errorf("StatusFailed = %v, want 'failed'", StatusFailed)
	}
	if StatusBlocked != "blocked" {
		t.Errorf("StatusBlocked = %v, want 'blocked'", StatusBlocked)
	}
	if StatusSkipped != "skipped" {
		t.Errorf("StatusSkipped = %v, want 'skipped'", StatusSkipped)
	}
	if StatusInProgress != "in_progress" {
		t.Errorf("StatusInProgress = %v, want 'in_progress'", StatusInProgress)
	}
	if StatusInvalid != "invalid" {
		t.Errorf("StatusInvalid = %v, want 'invalid'", StatusInvalid)
	}

	// Test StepStatus constants
	if StepStatusPassed != "passed" {
		t.Errorf("StepStatusPassed = %v, want 'passed'", StepStatusPassed)
	}
	if StepStatusFailed != "failed" {
		t.Errorf("StepStatusFailed = %v, want 'failed'", StepStatusFailed)
	}
	if StepStatusBlocked != "blocked" {
		t.Errorf("StepStatusBlocked = %v, want 'blocked'", StepStatusBlocked)
	}
	if StepStatusSkipped != "skipped" {
		t.Errorf("StepStatusSkipped = %v, want 'skipped'", StepStatusSkipped)
	}
}

