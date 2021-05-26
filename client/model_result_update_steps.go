package client

type ResultUpdateSteps struct {
	Position    int32    `json:"position,omitempty"`
	Status      string   `json:"status,omitempty"`
	Comment     string   `json:"comment,omitempty"`
	Attachments []string `json:"attachments,omitempty"`
}
