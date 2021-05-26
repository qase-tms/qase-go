package client

type TestStep struct {
	Hash                 string       `json:"hash,omitempty"`
	SharedStepHash       string       `json:"shared_step_hash,omitempty"`
	SharedStepNestedHash string       `json:"shared_step_nested_hash,omitempty"`
	Position             int32        `json:"position,omitempty"`
	Action               string       `json:"action,omitempty"`
	ExpectedResult       string       `json:"expected_result,omitempty"`
	Data                 string       `json:"data,omitempty"`
	Attachments          []Attachment `json:"attachments,omitempty"`
}
