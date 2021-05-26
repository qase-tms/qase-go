package client

type RunCreate struct {
	Title           string   `json:"title"`
	Description     string   `json:"description,omitempty"`
	IncludeAllCases bool     `json:"include_all_cases,omitempty"`
	Cases           []int32  `json:"cases,omitempty"`
	EnvironmentId   int32    `json:"environment_id,omitempty"`
	MilestoneId     int32    `json:"milestone_id,omitempty"`
	PlanId          int32    `json:"plan_id,omitempty"`
	Tags            []string `json:"tags,omitempty"`
}
