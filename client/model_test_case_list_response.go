package client

type TestCaseListResponse struct {
	Status bool                        `json:"status,omitempty"`
	Result *TestCaseListResponseResult `json:"result,omitempty"`
}
