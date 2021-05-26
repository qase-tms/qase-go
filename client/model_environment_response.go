package client

type EnvironmentResponse struct {
	Status bool         `json:"status,omitempty"`
	Result *Environment `json:"result,omitempty"`
}
