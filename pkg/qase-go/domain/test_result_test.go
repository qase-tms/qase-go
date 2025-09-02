package domain

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestNewTestResult(t *testing.T) {
	tests := []struct {
		name  string
		title string
	}{
		{
			name:  "basic title",
			title: "Test Basic Functionality",
		},
		{
			name:  "empty title",
			title: "",
		},
		{
			name:  "unicode title",
			title: "测试中文标题",
		},
		{
			name:  "long title",
			title: "This is a very long test title that should be handled properly by the system and not cause any issues with memory allocation or performance",
		},
		{
			name:  "special characters",
			title: "Test with @#$%^&*()_+ special chars",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewTestResult(tt.title)

			// Verify basic fields
			if result.Title != tt.title {
				t.Errorf("Title: got %q, want %q", result.Title, tt.title)
			}
			if result.ID != "" {
				t.Errorf("ID should be empty initially, got %q", result.ID)
			}
			if result.Signature != "" {
				t.Errorf("Signature should be empty initially, got %q", result.Signature)
			}
			if result.RunID != nil {
				t.Errorf("RunID should be nil initially, got %v", result.RunID)
			}
			if result.TestopsID != nil {
				t.Errorf("TestopsID should be nil initially, got %v", result.TestopsID)
			}
			if result.Relations != nil {
				t.Errorf("Relations should be nil initially, got %v", result.Relations)
			}
			if result.Message != nil {
				t.Errorf("Message should be nil initially, got %v", result.Message)
			}
			if result.Muted {
				t.Error("Muted should be false initially")
			}

			// Verify initialized maps and slices
			if result.Fields == nil {
				t.Error("Fields should be initialized")
			}
			if len(result.Fields) != 0 {
				t.Errorf("Fields should be empty initially, got %d items", len(result.Fields))
			}
			if result.Params == nil {
				t.Error("Params should be initialized")
			}
			if len(result.Params) != 0 {
				t.Errorf("Params should be empty initially, got %d items", len(result.Params))
			}
			if result.GroupParams == nil {
				t.Error("GroupParams should be initialized")
			}
			if len(result.GroupParams) != 0 {
				t.Errorf("GroupParams should be empty initially, got %d items", len(result.GroupParams))
			}
			if result.Attachments == nil {
				t.Error("Attachments should be initialized")
			}
			if len(result.Attachments) != 0 {
				t.Errorf("Attachments should be empty initially, got %d items", len(result.Attachments))
			}
			if result.Steps == nil {
				t.Error("Steps should be initialized")
			}
			if len(result.Steps) != 0 {
				t.Errorf("Steps should be empty initially, got %d items", len(result.Steps))
			}
		})
	}
}

func TestTestResult_SetRunID(t *testing.T) {
	testResult := NewTestResult("Test")

	tests := []struct {
		name  string
		runID int64
	}{
		{"positive ID", 123},
		{"zero ID", 0},
		{"large ID", 9223372036854775807}, // max int64
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testResult.SetRunID(tt.runID)

			if testResult.RunID == nil {
				t.Error("RunID should not be nil after SetRunID")
				return
			}
			if *testResult.RunID != tt.runID {
				t.Errorf("RunID: got %d, want %d", *testResult.RunID, tt.runID)
			}
		})
	}
}

func TestTestResult_SetTestopsID(t *testing.T) {
	testResult := NewTestResult("Test")

	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "single int64",
			input:    int64(123),
			expected: int64(123),
		},
		{
			name:     "slice of int64",
			input:    []int64{1, 2, 3},
			expected: []int64{1, 2, 3},
		},
		{
			name:     "nil value",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty slice",
			input:    []int64{},
			expected: []int64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testResult.SetTestopsID(tt.input)

			if !reflect.DeepEqual(testResult.TestopsID, tt.expected) {
				t.Errorf("TestopsID: got %v, want %v", testResult.TestopsID, tt.expected)
			}
		})
	}
}

func TestTestResult_SetTestopsIDSingle(t *testing.T) {
	testResult := NewTestResult("Test")

	tests := []struct {
		name string
		id   int64
	}{
		{"positive ID", 456},
		{"zero ID", 0},
		{"negative ID", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testResult.SetTestopsIDSingle(tt.id)

			if testResult.TestopsID != tt.id {
				t.Errorf("TestopsID: got %v, want %v", testResult.TestopsID, tt.id)
			}
		})
	}
}

func TestTestResult_SetTestopsIDMultiple(t *testing.T) {
	testResult := NewTestResult("Test")

	tests := []struct {
		name string
		ids  []int64
	}{
		{"multiple IDs", []int64{1, 2, 3, 4, 5}},
		{"single ID in slice", []int64{42}},
		{"empty slice", []int64{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testResult.SetTestopsIDMultiple(tt.ids)

			result, ok := testResult.TestopsID.([]int64)
			if !ok {
				t.Errorf("TestopsID should be []int64, got %T", testResult.TestopsID)
				return
			}
			if !reflect.DeepEqual(result, tt.ids) {
				t.Errorf("TestopsID: got %v, want %v", result, tt.ids)
			}
		})
	}
}

func TestTestResult_AddMessage(t *testing.T) {
	tests := []struct {
		name    string
		message string
	}{
		{"normal message", "Test completed successfully"},
		{"empty message", ""},
		{"unicode message", "Message avec accents éàü"},
		{"multiline message", "Line 1\nLine 2\nLine 3"},
		{"special chars", "Message with @#$%^&*()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testResult := NewTestResult("Test")
			testResult.AddMessage(tt.message)

			if testResult.Message == nil {
				t.Error("Message should not be nil after AddMessage")
				return
			}
			if *testResult.Message != tt.message {
				t.Errorf("Message: got %q, want %q", *testResult.Message, tt.message)
			}
		})
	}
}

func TestTestResult_AddMessageAccumulation(t *testing.T) {
	testResult := NewTestResult("Test")

	// Add first message
	testResult.AddMessage("First message")
	if testResult.Message == nil || *testResult.Message != "First message" {
		t.Error("First message should be set correctly")
	}

	// Add second message
	testResult.AddMessage("Second message")
	expected := "First message\nSecond message"
	if testResult.Message == nil || *testResult.Message != expected {
		t.Errorf("Messages should accumulate: got %q, want %q", *testResult.Message, expected)
	}

	// Add third message
	testResult.AddMessage("Third message")
	expected = "First message\nSecond message\nThird message"
	if testResult.Message == nil || *testResult.Message != expected {
		t.Errorf("Messages should accumulate: got %q, want %q", *testResult.Message, expected)
	}
}

func TestTestResult_AddSystemMessage(t *testing.T) {
	testResult := NewTestResult("Test")

	// Add user messages first
	testResult.AddMessage("User message 1")
	testResult.AddMessage("User message 2")

	// Add system message
	testResult.AddSystemMessage("System error message")

	expected := "User message 1\nUser message 2\n\n**System error message**"
	if testResult.Message == nil || *testResult.Message != expected {
		t.Errorf("System message should be separated and bold: got %q, want %q", *testResult.Message, expected)
	}
}

func TestTestResult_SetField(t *testing.T) {
	testResult := NewTestResult("Test")

	tests := []struct {
		name  string
		key   string
		value string
	}{
		{"basic field", "priority", "high"},
		{"empty value", "category", ""},
		{"unicode key", "título", "test"},
		{"unicode value", "description", "café"},
		{"special chars", "field@#$", "value!@#"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testResult.SetField(tt.key, tt.value)

			value, exists := testResult.Fields[tt.key]
			if !exists {
				t.Errorf("Field %q should exist", tt.key)
				return
			}
			if value != tt.value {
				t.Errorf("Field %q: got %q, want %q", tt.key, value, tt.value)
			}
		})
	}

	// Test multiple fields
	t.Run("multiple fields", func(t *testing.T) {
		testResult := NewTestResult("Test")
		testResult.SetField("priority", "high")
		testResult.SetField("category", "smoke")
		testResult.SetField("author", "dev")

		if len(testResult.Fields) != 3 {
			t.Errorf("Expected 3 fields, got %d", len(testResult.Fields))
		}

		expectedFields := map[string]string{
			"priority": "high",
			"category": "smoke",
			"author":   "dev",
		}

		for key, expectedValue := range expectedFields {
			if value, exists := testResult.Fields[key]; !exists {
				t.Errorf("Field %q should exist", key)
			} else if value != expectedValue {
				t.Errorf("Field %q: got %q, want %q", key, value, expectedValue)
			}
		}
	})
}

func TestTestResult_SetParam(t *testing.T) {
	testResult := NewTestResult("Test")

	tests := []struct {
		name  string
		key   string
		value string
	}{
		{"basic param", "browser", "chrome"},
		{"empty value", "environment", ""},
		{"numeric value", "timeout", "5000"},
		{"boolean value", "debug", "true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testResult.SetParam(tt.key, tt.value)

			value, exists := testResult.Params[tt.key]
			if !exists {
				t.Errorf("Param %q should exist", tt.key)
				return
			}
			if value != tt.value {
				t.Errorf("Param %q: got %q, want %q", tt.key, value, tt.value)
			}
		})
	}
}

func TestTestResult_SetGroupParam(t *testing.T) {
	testResult := NewTestResult("Test")

	tests := []struct {
		name  string
		key   string
		value string
	}{
		{"basic group param", "suite", "regression"},
		{"empty value", "tag", ""},
		{"complex value", "config", "env=prod,region=us-east"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testResult.SetGroupParam(tt.key, tt.value)

			value, exists := testResult.GroupParams[tt.key]
			if !exists {
				t.Errorf("GroupParam %q should exist", tt.key)
				return
			}
			if value != tt.value {
				t.Errorf("GroupParam %q: got %q, want %q", tt.key, value, tt.value)
			}
		})
	}
}

func TestTestResult_AddAttachment(t *testing.T) {
	testResult := NewTestResult("Test")

	attachment1 := *NewAttachment("att-1", "file1.txt", "text/plain", []byte("content1"))
	attachment2 := *NewAttachment("att-2", "file2.jpg", "image/jpeg", []byte{0xFF, 0xD8})

	testResult.AddAttachment(attachment1)
	if len(testResult.Attachments) != 1 {
		t.Errorf("Expected 1 attachment, got %d", len(testResult.Attachments))
	}
	if testResult.Attachments[0].ID != "att-1" {
		t.Errorf("First attachment ID: got %q, want %q", testResult.Attachments[0].ID, "att-1")
	}

	testResult.AddAttachment(attachment2)
	if len(testResult.Attachments) != 2 {
		t.Errorf("Expected 2 attachments, got %d", len(testResult.Attachments))
	}
	if testResult.Attachments[1].ID != "att-2" {
		t.Errorf("Second attachment ID: got %q, want %q", testResult.Attachments[1].ID, "att-2")
	}
}

func TestTestResult_AddStep(t *testing.T) {
	testResult := NewTestResult("Test")

	expected1 := "Button is clicked"
	expected2 := "Result is correct"

	step1 := TestStep{
		ID:       "step-1",
		StepType: StepTypeText,
		Data: StepTextData{
			Action:         "Click button",
			ExpectedResult: &expected1,
		},
	}

	step2 := TestStep{
		ID:       "step-2",
		StepType: StepTypeText,
		Data: StepTextData{
			Action:         "Verify result",
			ExpectedResult: &expected2,
		},
	}

	testResult.AddStep(step1)
	if len(testResult.Steps) != 1 {
		t.Errorf("Expected 1 step, got %d", len(testResult.Steps))
	}
	if testResult.Steps[0].ID != "step-1" {
		t.Errorf("First step ID: got %q, want %q", testResult.Steps[0].ID, "step-1")
	}

	testResult.AddStep(step2)
	if len(testResult.Steps) != 2 {
		t.Errorf("Expected 2 steps, got %d", len(testResult.Steps))
	}
	if testResult.Steps[1].ID != "step-2" {
		t.Errorf("Second step ID: got %q, want %q", testResult.Steps[1].ID, "step-2")
	}
}

func TestTestResult_SetSuite(t *testing.T) {
	testResult := NewTestResult("Test")

	publicID1 := int64(100)
	publicID2 := int64(200)

	suiteData := []SuiteData{
		{Title: "Suite 1", PublicID: &publicID1},
		{Title: "Suite 2", PublicID: &publicID2},
		{Title: "Suite 3", PublicID: nil},
	}

	testResult.SetSuite(suiteData)

	if testResult.Relations == nil {
		t.Error("Relations should not be nil after SetSuite")
		return
	}
	if testResult.Relations.Suite == nil {
		t.Error("Relations.Suite should not be nil after SetSuite")
		return
	}
	if len(testResult.Relations.Suite.Data) != 3 {
		t.Errorf("Expected 3 suite data items, got %d", len(testResult.Relations.Suite.Data))
		return
	}

	// Verify first suite data
	if testResult.Relations.Suite.Data[0].Title != "Suite 1" {
		t.Errorf("First suite title: got %q, want %q", testResult.Relations.Suite.Data[0].Title, "Suite 1")
	}
	if testResult.Relations.Suite.Data[0].PublicID == nil || *testResult.Relations.Suite.Data[0].PublicID != 100 {
		t.Errorf("First suite public ID: got %v, want %d", testResult.Relations.Suite.Data[0].PublicID, 100)
	}

	// Verify third suite data (with nil PublicID)
	if testResult.Relations.Suite.Data[2].Title != "Suite 3" {
		t.Errorf("Third suite title: got %q, want %q", testResult.Relations.Suite.Data[2].Title, "Suite 3")
	}
	if testResult.Relations.Suite.Data[2].PublicID != nil {
		t.Errorf("Third suite public ID should be nil, got %v", testResult.Relations.Suite.Data[2].PublicID)
	}
}

func TestTestResult_AddSuiteData(t *testing.T) {
	testResult := NewTestResult("Test")

	// Add first suite data
	publicID1 := int64(123)
	testResult.AddSuiteData("Main Suite", &publicID1)

	if testResult.Relations == nil {
		t.Error("Relations should not be nil after AddSuiteData")
		return
	}
	if testResult.Relations.Suite == nil {
		t.Error("Relations.Suite should not be nil after AddSuiteData")
		return
	}
	if len(testResult.Relations.Suite.Data) != 1 {
		t.Errorf("Expected 1 suite data item, got %d", len(testResult.Relations.Suite.Data))
		return
	}

	// Add second suite data
	testResult.AddSuiteData("Sub Suite", nil)

	if len(testResult.Relations.Suite.Data) != 2 {
		t.Errorf("Expected 2 suite data items, got %d", len(testResult.Relations.Suite.Data))
		return
	}

	// Verify first suite data
	firstSuite := testResult.Relations.Suite.Data[0]
	if firstSuite.Title != "Main Suite" {
		t.Errorf("First suite title: got %q, want %q", firstSuite.Title, "Main Suite")
	}
	if firstSuite.PublicID == nil || *firstSuite.PublicID != 123 {
		t.Errorf("First suite public ID: got %v, want %d", firstSuite.PublicID, 123)
	}

	// Verify second suite data
	secondSuite := testResult.Relations.Suite.Data[1]
	if secondSuite.Title != "Sub Suite" {
		t.Errorf("Second suite title: got %q, want %q", secondSuite.Title, "Sub Suite")
	}
	if secondSuite.PublicID != nil {
		t.Errorf("Second suite public ID should be nil, got %v", secondSuite.PublicID)
	}
}

func TestTestExecution_TimestampHandling(t *testing.T) {
	testResult := NewTestResult("Test")

	now := time.Now()
	startTime := now.UnixMilli()
	endTime := now.Add(time.Second * 5).UnixMilli()
	duration := int64(5000) // 5 seconds in milliseconds

	testResult.Execution.Status = StatusPassed
	testResult.Execution.StartTime = &startTime
	testResult.Execution.EndTime = &endTime
	testResult.Execution.Duration = &duration

	if testResult.Execution.StartTime == nil || *testResult.Execution.StartTime != startTime {
		t.Errorf("StartTime: got %v, want %d", testResult.Execution.StartTime, startTime)
	}
	if testResult.Execution.EndTime == nil || *testResult.Execution.EndTime != endTime {
		t.Errorf("EndTime: got %v, want %d", testResult.Execution.EndTime, endTime)
	}
	if testResult.Execution.Duration == nil || *testResult.Execution.Duration != duration {
		t.Errorf("Duration: got %v, want %d", testResult.Execution.Duration, duration)
	}
}

func TestTestResult_JSONSerialization(t *testing.T) {
	tests := []struct {
		name       string
		testResult *TestResult
		validate   func(*testing.T, []byte)
	}{
		{
			name: "minimal test result",
			testResult: &TestResult{
				ID:          "test-123",
				Title:       "Simple Test",
				Signature:   "test.simple",
				Execution:   TestExecution{Status: StatusPassed},
				Fields:      make(map[string]string),
				Attachments: make([]Attachment, 0),
				Steps:       make([]TestStep, 0),
				Params:      make(map[string]string),
				GroupParams: make(map[string]string),
				Muted:       false,
			},
			validate: func(t *testing.T, data []byte) {
				var result TestResult
				if err := json.Unmarshal(data, &result); err != nil {
					t.Fatalf("Failed to unmarshal: %v", err)
				}
				if result.ID != "test-123" {
					t.Errorf("ID: got %q, want %q", result.ID, "test-123")
				}
				if result.Title != "Simple Test" {
					t.Errorf("Title: got %q, want %q", result.Title, "Simple Test")
				}
			},
		},
		{
			name: "complete test result",
			testResult: func() *TestResult {
				tr := NewTestResult("Complete Test")
				tr.ID = "test-456"
				tr.Signature = "test.complete"
				tr.SetRunID(789)
				tr.SetTestopsIDSingle(101112)
				tr.AddMessage("Test completed successfully")
				tr.SetField("priority", "high")
				tr.SetParam("browser", "chrome")
				tr.SetGroupParam("suite", "smoke")
				tr.Muted = false

				// Set execution details
				now := time.Now().UnixMilli()
				tr.Execution.Status = StatusPassed
				tr.Execution.StartTime = &now
				tr.Execution.EndTime = &now
				tr.Execution.Duration = &[]int64{1000}[0]

				// Add attachment
				attachment := *NewAttachment("att-1", "screenshot.png", "image/png", []byte("fake-image-data"))
				tr.AddAttachment(attachment)

				// Add suite data
				publicID := int64(999)
				tr.AddSuiteData("Main Suite", &publicID)

				return tr
			}(),
			validate: func(t *testing.T, data []byte) {
				var result TestResult
				if err := json.Unmarshal(data, &result); err != nil {
					t.Fatalf("Failed to unmarshal: %v", err)
				}

				// Verify deserialized data
				if result.ID != "test-456" {
					t.Errorf("ID: got %q, want %q", result.ID, "test-456")
				}
				if result.RunID == nil || *result.RunID != 789 {
					t.Errorf("RunID: got %v, want %d", result.RunID, 789)
				}
				if result.TestopsID != float64(101112) { // JSON numbers become float64
					t.Errorf("TestopsID: got %v, want %v", result.TestopsID, 101112)
				}
				if result.Message == nil || *result.Message != "Test completed successfully" {
					t.Errorf("Message: got %v, want %q", result.Message, "Test completed successfully")
				}
				if len(result.Fields) != 1 || result.Fields["priority"] != "high" {
					t.Errorf("Fields: got %v, want %v", result.Fields, map[string]string{"priority": "high"})
				}
				if len(result.Attachments) != 1 {
					t.Errorf("Attachments count: got %d, want %d", len(result.Attachments), 1)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.testResult)
			if err != nil {
				t.Fatalf("Failed to marshal test result: %v", err)
			}

			tt.validate(t, data)
		})
	}
}

func TestTestResult_EdgeCases(t *testing.T) {
	t.Run("nil maps and slices", func(t *testing.T) {
		tr := &TestResult{
			Title:       "Test",
			Fields:      nil,
			Params:      nil,
			GroupParams: nil,
			Attachments: nil,
			Steps:       nil,
		}

		// In Go, writing to nil map causes panic - this is expected behavior
		// We test that NewTestResult properly initializes maps to avoid this
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic when writing to nil map")
			}
		}()

		// This should panic with nil map (demonstrating the importance of proper initialization)
		tr.SetField("key", "value")
	})

	t.Run("proper initialization prevents panics", func(t *testing.T) {
		// NewTestResult should properly initialize all maps and slices
		tr := NewTestResult("Test")

		// These operations should not panic
		tr.SetField("key", "value")
		tr.SetParam("param", "value")
		tr.SetGroupParam("group", "value")

		attachment := *NewAttachment("att-1", "file.txt", "text/plain", []byte("content"))
		tr.AddAttachment(attachment)

		step := TestStep{ID: "step-1", StepType: StepTypeText}
		tr.AddStep(step)

		// Verify all operations worked
		if tr.Fields["key"] != "value" {
			t.Error("Field not set properly")
		}
		if tr.Params["param"] != "value" {
			t.Error("Param not set properly")
		}
		if tr.GroupParams["group"] != "value" {
			t.Error("GroupParam not set properly")
		}
		if len(tr.Attachments) != 1 {
			t.Error("Attachment not added properly")
		}
		if len(tr.Steps) != 1 {
			t.Error("Step not added properly")
		}
	})

	t.Run("very long strings", func(t *testing.T) {
		longString := string(make([]byte, 10000))
		tr := NewTestResult(longString)

		if tr.Title != longString {
			t.Error("Long title should be preserved")
		}

		tr.AddMessage(longString)
		if tr.Message == nil || *tr.Message != longString {
			t.Error("Long message should be preserved")
		}
	})

	t.Run("unicode handling", func(t *testing.T) {
		tr := NewTestResult("测试 ✓ ñáéíóú")
		tr.SetField("描述", "测试描述")
		tr.SetParam("parámetro", "valor")
		tr.SetGroupParam("groupe", "paramètre")
		tr.AddMessage("Message d'erreur")

		// Verify unicode is preserved
		if tr.Title != "测试 ✓ ñáéíóú" {
			t.Error("Unicode title should be preserved")
		}
		if tr.Fields["描述"] != "测试描述" {
			t.Error("Unicode field should be preserved")
		}
	})
}
