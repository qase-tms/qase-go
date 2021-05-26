package client

type SuiteResponse struct {
	Status bool   `json:"status,omitempty"`
	Result *Suite `json:"result,omitempty"`
}
