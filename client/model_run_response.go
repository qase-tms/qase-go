package client

type RunResponse struct {
	Status bool `json:"status,omitempty"`
	Result *Run `json:"result,omitempty"`
}
