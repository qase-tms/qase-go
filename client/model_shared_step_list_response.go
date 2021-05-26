package client

type SharedStepListResponse struct {
	Status bool                          `json:"status,omitempty"`
	Result *SharedStepListResponseResult `json:"result,omitempty"`
}
