package client

type RunListResponse struct {
	Status bool                   `json:"status,omitempty"`
	Result *RunListResponseResult `json:"result,omitempty"`
}
