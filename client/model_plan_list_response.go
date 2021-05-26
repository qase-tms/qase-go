package client

type PlanListResponse struct {
	Status bool                    `json:"status,omitempty"`
	Result *PlanListResponseResult `json:"result,omitempty"`
}
