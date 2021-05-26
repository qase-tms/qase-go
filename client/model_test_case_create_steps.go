package client

type TestCaseCreateSteps struct {
	Action         string   `json:"action,omitempty"`
	ExpectedResult string   `json:"expected_result,omitempty"`
	InputData      string   `json:"input_data,omitempty"`
	Position       int32    `json:"position,omitempty"`
	Attachments    []string `json:"attachments,omitempty"`
}
