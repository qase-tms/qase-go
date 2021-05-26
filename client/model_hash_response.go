package client

type HashResponse struct {
	Status bool                `json:"status,omitempty"`
	Result *HashResponseResult `json:"result,omitempty"`
}
