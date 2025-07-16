package domain

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNewTestStep(t *testing.T) {
	tests := []struct {
		name   string
		action string
	}{
		{
			name:   "basic action",
			action: "Click the submit button",
		},
		{
			name:   "empty action",
			action: "",
		},
		{
			name:   "unicode action",
			action: "Cliquez sur le bouton avec accents éàü",
		},
		{
			name:   "long action",
			action: "This is a very long action description that should be handled properly by the system and not cause any issues with memory allocation or performance when creating test steps",
		},
		{
			name:   "special characters",
			action: "Click @#$%^&*()_+ special chars button",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step := NewTestStep(tt.action)

			// Verify basic fields
			if step.Data.Action != tt.action {
				t.Errorf("Action: got %q, want %q", step.Data.Action, tt.action)
			}
			if step.ID != "" {
				t.Errorf("ID should be empty initially, got %q", step.ID)
			}
			if step.StepType != StepTypeText {
				t.Errorf("StepType: got %q, want %q", step.StepType, StepTypeText)
			}
			if step.ParentID != nil {
				t.Errorf("ParentID should be nil initially, got %v", step.ParentID)
			}
			if step.Data.ExpectedResult != nil {
				t.Errorf("ExpectedResult should be nil initially, got %v", step.Data.ExpectedResult)
			}
			if step.Data.Data != nil {
				t.Errorf("Data should be nil initially, got %v", step.Data.Data)
			}

			// Verify execution defaults
			if step.Execution.Status != StepStatusPassed {
				t.Errorf("Default status: got %q, want %q", step.Execution.Status, StepStatusPassed)
			}
			if step.Execution.StartTime != nil {
				t.Errorf("StartTime should be nil initially, got %v", step.Execution.StartTime)
			}
			if step.Execution.EndTime != nil {
				t.Errorf("EndTime should be nil initially, got %v", step.Execution.EndTime)
			}
			if step.Execution.Duration != nil {
				t.Errorf("Duration should be nil initially, got %v", step.Execution.Duration)
			}

			// Verify initialized slices
			if step.Attachments == nil {
				t.Error("Attachments should be initialized")
			}
			if len(step.Attachments) != 0 {
				t.Errorf("Attachments should be empty initially, got %d items", len(step.Attachments))
			}
			if step.Steps == nil {
				t.Error("Steps should be initialized")
			}
			if len(step.Steps) != 0 {
				t.Errorf("Steps should be empty initially, got %d items", len(step.Steps))
			}
		})
	}
}

func TestNewTestStepWithStatus(t *testing.T) {
	tests := []struct {
		name   string
		action string
		status StepStatus
	}{
		{"passed step", "Click button", StepStatusPassed},
		{"failed step", "Verify result", StepStatusFailed},
		{"blocked step", "Wait for element", StepStatusBlocked},
		{"skipped step", "Optional action", StepStatusSkipped},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step := NewTestStepWithStatus(tt.action, tt.status)

			if step.Data.Action != tt.action {
				t.Errorf("Action: got %q, want %q", step.Data.Action, tt.action)
			}
			if step.Execution.Status != tt.status {
				t.Errorf("Status: got %q, want %q", step.Execution.Status, tt.status)
			}
			if step.StepType != StepTypeText {
				t.Errorf("StepType: got %q, want %q", step.StepType, StepTypeText)
			}
		})
	}
}

func TestTestStep_SetID(t *testing.T) {
	step := NewTestStep("Test action")

	tests := []struct {
		name string
		id   string
	}{
		{"basic ID", "step-123"},
		{"empty ID", ""},
		{"uuid-like ID", "550e8400-e29b-41d4-a716-446655440000"},
		{"numeric ID", "12345"},
		{"special chars ID", "step@#$%"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step.SetID(tt.id)

			if step.ID != tt.id {
				t.Errorf("ID: got %q, want %q", step.ID, tt.id)
			}
		})
	}
}

func TestTestStep_SetParentID(t *testing.T) {
	step := NewTestStep("Test action")

	tests := []struct {
		name     string
		parentID string
	}{
		{"basic parent ID", "parent-123"},
		{"empty parent ID", ""},
		{"uuid parent ID", "parent-550e8400-e29b-41d4-a716"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step.SetParentID(tt.parentID)

			if step.ParentID == nil {
				t.Error("ParentID should not be nil after SetParentID")
				return
			}
			if *step.ParentID != tt.parentID {
				t.Errorf("ParentID: got %q, want %q", *step.ParentID, tt.parentID)
			}
		})
	}
}

func TestTestStep_SetAction(t *testing.T) {
	step := NewTestStep("Initial action")

	tests := []struct {
		name   string
		action string
	}{
		{"update action", "Updated action"},
		{"empty action", ""},
		{"unicode action", "Action avec accents éàü"},
		{"long action", "Very long action description that should be handled properly"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step.SetAction(tt.action)

			if step.Data.Action != tt.action {
				t.Errorf("Action: got %q, want %q", step.Data.Action, tt.action)
			}
		})
	}
}

func TestTestStep_SetExpectedResult(t *testing.T) {
	step := NewTestStep("Test action")

	tests := []struct {
		name   string
		result string
	}{
		{"basic result", "Button should be clicked"},
		{"empty result", ""},
		{"unicode result", "Le bouton devrait être cliqué"},
		{"detailed result", "The submit button should change color to green and display a checkmark"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step.SetExpectedResult(tt.result)

			if step.Data.ExpectedResult == nil {
				t.Error("ExpectedResult should not be nil after SetExpectedResult")
				return
			}
			if *step.Data.ExpectedResult != tt.result {
				t.Errorf("ExpectedResult: got %q, want %q", *step.Data.ExpectedResult, tt.result)
			}
		})
	}
}

func TestTestStep_SetData(t *testing.T) {
	step := NewTestStep("Test action")

	tests := []struct {
		name string
		data string
	}{
		{"basic data", "Additional step data"},
		{"empty data", ""},
		{"json data", `{"key": "value", "number": 123}`},
		{"multiline data", "Line 1\nLine 2\nLine 3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step.SetData(tt.data)

			if step.Data.Data == nil {
				t.Error("Data should not be nil after SetData")
				return
			}
			if *step.Data.Data != tt.data {
				t.Errorf("Data: got %q, want %q", *step.Data.Data, tt.data)
			}
		})
	}
}

func TestTestStep_SetStatus(t *testing.T) {
	step := NewTestStep("Test action")

	tests := []struct {
		name   string
		status StepStatus
	}{
		{"passed", StepStatusPassed},
		{"failed", StepStatusFailed},
		{"blocked", StepStatusBlocked},
		{"skipped", StepStatusSkipped},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step.SetStatus(tt.status)

			if step.Execution.Status != tt.status {
				t.Errorf("Status: got %q, want %q", step.Execution.Status, tt.status)
			}
		})
	}
}

func TestTestStep_SetTiming(t *testing.T) {
	step := NewTestStep("Test action")

	tests := []struct {
		name      string
		startTime int64
		endTime   int64
		duration  int64
	}{
		{
			name:      "basic timing",
			startTime: 1000,
			endTime:   2000,
			duration:  1000,
		},
		{
			name:      "zero duration",
			startTime: 1000,
			endTime:   1000,
			duration:  0,
		},
		{
			name:      "large duration",
			startTime: 1000,
			endTime:   1000000,
			duration:  999000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step.SetTiming(tt.startTime, tt.endTime)

			if step.Execution.StartTime == nil || *step.Execution.StartTime != tt.startTime {
				t.Errorf("StartTime: got %v, want %d", step.Execution.StartTime, tt.startTime)
			}
			if step.Execution.EndTime == nil || *step.Execution.EndTime != tt.endTime {
				t.Errorf("EndTime: got %v, want %d", step.Execution.EndTime, tt.endTime)
			}
			if tt.endTime > tt.startTime {
				if step.Execution.Duration == nil || *step.Execution.Duration != tt.duration {
					t.Errorf("Duration: got %v, want %d", step.Execution.Duration, tt.duration)
				}
			}
		})
	}

	t.Run("end time before start time", func(t *testing.T) {
		step := NewTestStep("Test action")
		step.SetTiming(2000, 1000) // end before start

		if step.Execution.StartTime == nil || *step.Execution.StartTime != 2000 {
			t.Errorf("StartTime: got %v, want %d", step.Execution.StartTime, 2000)
		}
		if step.Execution.EndTime == nil || *step.Execution.EndTime != 1000 {
			t.Errorf("EndTime: got %v, want %d", step.Execution.EndTime, 1000)
		}
		// Duration should not be set when end < start
		if step.Execution.Duration != nil {
			t.Errorf("Duration should be nil when end < start, got %v", step.Execution.Duration)
		}
	})
}

func TestTestStep_SetDuration(t *testing.T) {
	step := NewTestStep("Test action")

	tests := []struct {
		name     string
		duration int64
	}{
		{"short duration", 100},
		{"zero duration", 0},
		{"long duration", 10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step.SetDuration(tt.duration)

			if step.Execution.Duration == nil || *step.Execution.Duration != tt.duration {
				t.Errorf("Duration: got %v, want %d", step.Execution.Duration, tt.duration)
			}
		})
	}
}

func TestTestStep_AddAttachment(t *testing.T) {
	step := NewTestStep("Test action")

	attachment1 := *NewAttachment("att-1", "step1.txt", "text/plain", []byte("step content 1"))
	attachment2 := *NewAttachment("att-2", "step2.jpg", "image/jpeg", []byte{0xFF, 0xD8})

	step.AddAttachment(attachment1)
	if len(step.Attachments) != 1 {
		t.Errorf("Expected 1 attachment, got %d", len(step.Attachments))
	}
	if step.Attachments[0].ID != "att-1" {
		t.Errorf("First attachment ID: got %q, want %q", step.Attachments[0].ID, "att-1")
	}

	step.AddAttachment(attachment2)
	if len(step.Attachments) != 2 {
		t.Errorf("Expected 2 attachments, got %d", len(step.Attachments))
	}
	if step.Attachments[1].ID != "att-2" {
		t.Errorf("Second attachment ID: got %q, want %q", step.Attachments[1].ID, "att-2")
	}
}

func TestTestStep_AddStep(t *testing.T) {
	parentStep := NewTestStep("Parent action")

	childStep1 := NewTestStep("Child action 1")
	childStep1.SetID("child-1")

	childStep2 := NewTestStep("Child action 2")
	childStep2.SetID("child-2")

	parentStep.AddStep(*childStep1)
	if len(parentStep.Steps) != 1 {
		t.Errorf("Expected 1 child step, got %d", len(parentStep.Steps))
	}
	if parentStep.Steps[0].ID != "child-1" {
		t.Errorf("First child step ID: got %q, want %q", parentStep.Steps[0].ID, "child-1")
	}

	parentStep.AddStep(*childStep2)
	if len(parentStep.Steps) != 2 {
		t.Errorf("Expected 2 child steps, got %d", len(parentStep.Steps))
	}
	if parentStep.Steps[1].ID != "child-2" {
		t.Errorf("Second child step ID: got %q, want %q", parentStep.Steps[1].ID, "child-2")
	}
}

func TestTestStep_GetStatus(t *testing.T) {
	tests := []struct {
		name   string
		status StepStatus
	}{
		{"passed", StepStatusPassed},
		{"failed", StepStatusFailed},
		{"blocked", StepStatusBlocked},
		{"skipped", StepStatusSkipped},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step := NewTestStep("Test action")
			step.SetStatus(tt.status)

			result := step.GetStatus()
			if result != tt.status {
				t.Errorf("GetStatus: got %q, want %q", result, tt.status)
			}
		})
	}
}

func TestTestStep_GetAction(t *testing.T) {
	tests := []struct {
		name   string
		action string
	}{
		{"basic action", "Click button"},
		{"empty action", ""},
		{"unicode action", "Clic sur le bouton"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step := NewTestStep(tt.action)

			result := step.GetAction()
			if result != tt.action {
				t.Errorf("GetAction: got %q, want %q", result, tt.action)
			}
		})
	}
}

func TestTestStep_IsSuccess(t *testing.T) {
	tests := []struct {
		name     string
		status   StepStatus
		expected bool
	}{
		{"passed is success", StepStatusPassed, true},
		{"failed is not success", StepStatusFailed, false},
		{"blocked is not success", StepStatusBlocked, false},
		{"skipped is not success", StepStatusSkipped, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step := NewTestStep("Test action")
			step.SetStatus(tt.status)

			result := step.IsSuccess()
			if result != tt.expected {
				t.Errorf("IsSuccess: got %t, want %t", result, tt.expected)
			}
		})
	}
}

func TestTestStep_IsFailure(t *testing.T) {
	tests := []struct {
		name     string
		status   StepStatus
		expected bool
	}{
		{"passed is not failure", StepStatusPassed, false},
		{"failed is failure", StepStatusFailed, true},
		{"blocked is failure", StepStatusBlocked, true},
		{"skipped is not failure", StepStatusSkipped, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step := NewTestStep("Test action")
			step.SetStatus(tt.status)

			result := step.IsFailure()
			if result != tt.expected {
				t.Errorf("IsFailure: got %t, want %t", result, tt.expected)
			}
		})
	}
}

func TestTestStep_NestedSteps(t *testing.T) {
	// Create a hierarchy of steps
	rootStep := NewTestStep("Root step")
	rootStep.SetID("root")

	level1Step1 := NewTestStep("Level 1 Step 1")
	level1Step1.SetID("level1-1")
	level1Step1.SetParentID("root")

	level1Step2 := NewTestStep("Level 1 Step 2")
	level1Step2.SetID("level1-2")
	level1Step2.SetParentID("root")

	level2Step := NewTestStep("Level 2 Step")
	level2Step.SetID("level2-1")
	level2Step.SetParentID("level1-1")

	// Build hierarchy
	level1Step1.AddStep(*level2Step)
	rootStep.AddStep(*level1Step1)
	rootStep.AddStep(*level1Step2)

	// Verify structure
	if len(rootStep.Steps) != 2 {
		t.Errorf("Root step should have 2 children, got %d", len(rootStep.Steps))
	}

	if len(rootStep.Steps[0].Steps) != 1 {
		t.Errorf("First child should have 1 grandchild, got %d", len(rootStep.Steps[0].Steps))
	}

	if len(rootStep.Steps[1].Steps) != 0 {
		t.Errorf("Second child should have 0 grandchildren, got %d", len(rootStep.Steps[1].Steps))
	}

	// Verify IDs
	if rootStep.Steps[0].ID != "level1-1" {
		t.Errorf("First child ID: got %q, want %q", rootStep.Steps[0].ID, "level1-1")
	}

	if rootStep.Steps[0].Steps[0].ID != "level2-1" {
		t.Errorf("Grandchild ID: got %q, want %q", rootStep.Steps[0].Steps[0].ID, "level2-1")
	}

	// Verify parent IDs
	if rootStep.Steps[0].ParentID == nil || *rootStep.Steps[0].ParentID != "root" {
		t.Errorf("First child parent ID: got %v, want %q", rootStep.Steps[0].ParentID, "root")
	}

	if rootStep.Steps[0].Steps[0].ParentID == nil || *rootStep.Steps[0].Steps[0].ParentID != "level1-1" {
		t.Errorf("Grandchild parent ID: got %v, want %q", rootStep.Steps[0].Steps[0].ParentID, "level1-1")
	}
}

func TestTestStep_JSONSerialization(t *testing.T) {
	tests := []struct {
		name     string
		step     *TestStep
		validate func(*testing.T, []byte)
	}{
		{
			name: "minimal step",
			step: &TestStep{
				ID:       "step-1",
				StepType: StepTypeText,
				Data: StepTextData{
					Action: "Click button",
				},
				Execution: StepExecution{
					Status: StepStatusPassed,
				},
				Attachments: make([]Attachment, 0),
				Steps:       make([]TestStep, 0),
			},
			validate: func(t *testing.T, data []byte) {
				var step TestStep
				if err := json.Unmarshal(data, &step); err != nil {
					t.Fatalf("Failed to unmarshal: %v", err)
				}
				if step.ID != "step-1" {
					t.Errorf("ID: got %q, want %q", step.ID, "step-1")
				}
				if step.Data.Action != "Click button" {
					t.Errorf("Action: got %q, want %q", step.Data.Action, "Click button")
				}
			},
		},
		{
			name: "complete step with timing",
			step: func() *TestStep {
				step := NewTestStep("Complete step action")
				step.SetID("step-complete")
				step.SetParentID("parent-step")
				step.SetExpectedResult("Expected result")
				step.SetData("Additional data")
				step.SetStatus(StepStatusPassed)
				step.SetTiming(1000, 2000)

				attachment := *NewAttachment("att-1", "step.png", "image/png", []byte("image-data"))
				step.AddAttachment(attachment)

				childStep := NewTestStep("Child step")
				childStep.SetID("child-1")
				step.AddStep(*childStep)

				return step
			}(),
			validate: func(t *testing.T, data []byte) {
				var step TestStep
				if err := json.Unmarshal(data, &step); err != nil {
					t.Fatalf("Failed to unmarshal: %v", err)
				}

				if step.ID != "step-complete" {
					t.Errorf("ID: got %q, want %q", step.ID, "step-complete")
				}
				if step.ParentID == nil || *step.ParentID != "parent-step" {
					t.Errorf("ParentID: got %v, want %q", step.ParentID, "parent-step")
				}
				if step.Data.ExpectedResult == nil || *step.Data.ExpectedResult != "Expected result" {
					t.Errorf("ExpectedResult: got %v, want %q", step.Data.ExpectedResult, "Expected result")
				}
				if step.Data.Data == nil || *step.Data.Data != "Additional data" {
					t.Errorf("Data: got %v, want %q", step.Data.Data, "Additional data")
				}
				if len(step.Attachments) != 1 {
					t.Errorf("Attachments count: got %d, want %d", len(step.Attachments), 1)
				}
				if len(step.Steps) != 1 {
					t.Errorf("Steps count: got %d, want %d", len(step.Steps), 1)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.step)
			if err != nil {
				t.Fatalf("Failed to marshal test step: %v", err)
			}

			tt.validate(t, data)
		})
	}
}

func TestTestStep_EdgeCases(t *testing.T) {
	t.Run("real-world timing", func(t *testing.T) {
		step := NewTestStep("Real timing test")
		
		now := time.Now()
		startTime := now.UnixMilli()
		endTime := now.Add(time.Millisecond * 500).UnixMilli()
		
		step.SetTiming(startTime, endTime)
		
		if step.Execution.StartTime == nil || *step.Execution.StartTime != startTime {
			t.Errorf("StartTime: got %v, want %d", step.Execution.StartTime, startTime)
		}
		
		expectedDuration := endTime - startTime
		if step.Execution.Duration == nil || *step.Execution.Duration != expectedDuration {
			t.Errorf("Duration: got %v, want %d", step.Execution.Duration, expectedDuration)
		}
	})

	t.Run("step type constants", func(t *testing.T) {
		step := NewTestStep("Test action")
		
		if step.StepType != StepTypeText {
			t.Errorf("StepType: got %q, want %q", step.StepType, StepTypeText)
		}
		
		// Verify the constant value
		if StepTypeText != "text" {
			t.Errorf("StepTypeText constant: got %q, want %q", StepTypeText, "text")
		}
	})

	t.Run("deep nesting", func(t *testing.T) {
		root := NewTestStep("Root")
		root.SetID("root")
		
		current := root
		for i := 1; i <= 5; i++ {
			child := NewTestStep("Step " + string(rune('0'+i)))
			child.SetID("step-" + string(rune('0'+i)))
			current.AddStep(*child)
			// For next iteration, work with the child that was just added
			current = &current.Steps[0]
		}
		
		// Verify depth
		current = root
		depth := 0
		for len(current.Steps) > 0 {
			depth++
			current = &current.Steps[0]
		}
		
		if depth != 5 {
			t.Errorf("Expected depth 5, got %d", depth)
		}
	})

	t.Run("unicode handling", func(t *testing.T) {
		step := NewTestStep("测试步骤")
		step.SetExpectedResult("Résultat attendu avec accents")
		step.SetData("Données avec caractères spéciaux: éàùç")
		
		if step.Data.Action != "测试步骤" {
			t.Error("Unicode action should be preserved")
		}
		if step.Data.ExpectedResult == nil || *step.Data.ExpectedResult != "Résultat attendu avec accents" {
			t.Error("Unicode expected result should be preserved")
		}
		if step.Data.Data == nil || *step.Data.Data != "Données avec caractères spéciaux: éàùç" {
			t.Error("Unicode data should be preserved")
		}
	})
}

