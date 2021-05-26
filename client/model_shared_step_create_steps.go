package client

type SharedStepCreateSteps struct {
	Hash           string           `json:"hash,omitempty"`
	Action         string           `json:"action,omitempty"`
	ExpectedResult string           `json:"expected_result,omitempty"`
	Data           string           `json:"data,omitempty"`
	Attachments    []AttachmentHash `json:"attachments,omitempty"`
}
