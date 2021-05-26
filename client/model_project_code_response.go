package client

type ProjectCodeResponse struct {
	Status bool                       `json:"status,omitempty"`
	Result *ProjectCodeResponseResult `json:"result,omitempty"`
}
