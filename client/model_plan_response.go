package client

type PlanResponse struct {
	Status bool          `json:"status,omitempty"`
	Result *PlanDetailed `json:"result,omitempty"`
}
