package qase

// StepMetadata represents step metadata
type StepMetadata struct {
	Name           string
	Description    string
	ExpectedResult string
	Data           string
}

// TestMetadata represents test metadata matching JavaScript interface
type TestMetadata struct {
	// Core test information
	DisplayName string
	Title       string
	Description string
	Comment     string

	// Test case IDs (matching JavaScript ids field)
	IDs []int64

	// Suite information
	Suite string

	// Metadata fields (matching JavaScript fields)
	Fields map[string]string

	// Parameters (matching JavaScript parameters)
	Parameters map[string]string

	// Group parameters (matching JavaScript groupParams)
	GroupParameters map[string]string

	// Test control
	Ignore bool
}
