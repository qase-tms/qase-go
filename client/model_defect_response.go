package client

type DefectResponse struct {
	Status bool    `json:"status,omitempty"`
	Result *Defect `json:"result,omitempty"`
}
