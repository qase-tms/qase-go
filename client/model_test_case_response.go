package client

type TestCaseResponse struct {
	Status bool      `json:"status,omitempty"`
	Result *TestCase `json:"result,omitempty"`
}
