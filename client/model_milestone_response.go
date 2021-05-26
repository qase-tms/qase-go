package client

type MilestoneResponse struct {
	Status bool       `json:"status,omitempty"`
	Result *Milestone `json:"result,omitempty"`
}
