package client

type SuiteListResponse struct {
	Status bool                     `json:"status,omitempty"`
	Result *SuiteListResponseResult `json:"result,omitempty"`
}
