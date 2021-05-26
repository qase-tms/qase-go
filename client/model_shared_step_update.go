package client

type SharedStepUpdate struct {
	Title          string                  `json:"title,omitempty"`
	Action         string                  `json:"action,omitempty"`
	ExpectedResult string                  `json:"expected_result,omitempty"`
	Data           string                  `json:"data,omitempty"`
	Steps          []SharedStepCreateSteps `json:"steps,omitempty"`
}
