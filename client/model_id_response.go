package client

type IdResponse struct {
	Status bool              `json:"status,omitempty"`
	Result *IdResponseResult `json:"result,omitempty"`
}
