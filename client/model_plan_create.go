package client

type PlanCreate struct {
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	Cases       []int32 `json:"cases"`
}
