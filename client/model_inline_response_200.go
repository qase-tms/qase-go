package client

type InlineResponse200 struct {
	Status bool                     `json:"status,omitempty"`
	Result *InlineResponse200Result `json:"result,omitempty"`
}
