package client

type ResultListResponse struct {
	Status bool                      `json:"status,omitempty"`
	Result *ResultListResponseResult `json:"result,omitempty"`
}
