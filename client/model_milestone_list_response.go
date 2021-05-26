package client

type MilestoneListResponse struct {
	Status bool                         `json:"status,omitempty"`
	Result *MilestoneListResponseResult `json:"result,omitempty"`
}
