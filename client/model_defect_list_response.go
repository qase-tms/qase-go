package client

type DefectListResponse struct {
	Status bool                      `json:"status,omitempty"`
	Result *DefectListResponseResult `json:"result,omitempty"`
}
