package client

type ProjectResponse struct {
	Status bool     `json:"status,omitempty"`
	Result *Project `json:"result,omitempty"`
}
