package qase

// StepMetadata represents step metadata
type StepMetadata struct {
	Name           string
	ExpectedResult string
	Data           string
}

// TestMetadata represents test metadata matching JavaScript interface
type TestMetadata struct {
	// Core test information
	Title       string
	Description string
	Comment     string

	// Qase test case IDs
	QaseIDs []int64

	// Suite information
	Suite []string

	// Metadata fields (matching JavaScript fields)
	Fields map[string]string

	// Parameters (matching JavaScript parameters)
	Parameters map[string]string

	// Group parameters (matching JavaScript groupParams)
	GroupParameters map[string]string

	// Tags for the test case
	Tags []string

	// Test control
	Ignore bool
}
