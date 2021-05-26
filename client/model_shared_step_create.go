package client

type SharedStepCreate struct {
	Title          string                  `json:"title,omitempty"`
	Action         string                  `json:"action,omitempty"`
	ExpectedResult string                  `json:"expected_result,omitempty"`
	Data           string                  `json:"data,omitempty"`
	Attachments    []AttachmentHash        `json:"attachments,omitempty"`
	Steps          []SharedStepCreateSteps `json:"steps,omitempty"`
}
