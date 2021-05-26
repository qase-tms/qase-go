package client

type ResultResponse struct {
	Status bool    `json:"status,omitempty"`
	Result *Result `json:"result,omitempty"`
}
