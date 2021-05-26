package client

type Response struct {
	Status bool         `json:"status,omitempty"`
	Result *interface{} `json:"result,omitempty"`
}
