package client

type SharedStepResponse struct {
	Status bool        `json:"status,omitempty"`
	Result *SharedStep `json:"result,omitempty"`
}
