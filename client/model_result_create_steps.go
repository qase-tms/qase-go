package client

type ResultCreateSteps struct {
	Position       int32    `json:"position"`
	Status         string   `json:"status"`
	Comment        string   `json:"comment,omitempty"`
	Attachments    []string `json:"attachments,omitempty"`
	Action         string   `json:"action,omitempty"`
	ExpectedResult string   `json:"expected_result,omitempty"`
	Data           string   `json:"data,omitempty"`
}
