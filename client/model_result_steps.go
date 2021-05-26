package client

type ResultSteps struct {
	Status      int32        `json:"status,omitempty"`
	Position    int32        `json:"position,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}
