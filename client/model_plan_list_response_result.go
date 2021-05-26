package client

type PlanListResponseResult struct {
	Total    int32  `json:"total,omitempty"`
	Filtered int32  `json:"filtered,omitempty"`
	Count    int32  `json:"count,omitempty"`
	Entities []Plan `json:"entities,omitempty"`
}
