package client

type SharedStepSteps struct {
	Data           string           `json:"data,omitempty"`
	Hash           string           `json:"hash,omitempty"`
	Action         string           `json:"action,omitempty"`
	ExpectedResult string           `json:"expected_result,omitempty"`
	Attachments    []AttachmentHash `json:"attachments,omitempty"`
}
