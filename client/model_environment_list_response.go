package client

type EnvironmentListResponse struct {
	Status bool                           `json:"status,omitempty"`
	Result *EnvironmentListResponseResult `json:"result,omitempty"`
}
