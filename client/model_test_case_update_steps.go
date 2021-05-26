package client

type TestCaseUpdateSteps struct {
	Action         string `json:"action,omitempty"`
	ExpectedResult string `json:"expected_result,omitempty"`
	InputData      string `json:"input_data,omitempty"`
	Position       int32  `json:"position,omitempty"`
}
