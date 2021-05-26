package client

type ProjectListResponse struct {
	Status bool                       `json:"status,omitempty"`
	Result *ProjectListResponseResult `json:"result,omitempty"`
}
