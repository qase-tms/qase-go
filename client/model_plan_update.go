package client

type PlanUpdate struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Cases       []int32 `json:"cases,omitempty"`
}
