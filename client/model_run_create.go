package client

type RunCreate struct {
	Title           string   `json:"title"`
	Description     string   `json:"description,omitempty"`
	IncludeAllCases bool     `json:"include_all_cases,omitempty"`
	Cases           []int64  `json:"cases,omitempty"`
	IsAutotest      bool     `json:"is_autotest,omitempty"`
	EnvironmentId   int64    `json:"environment_id,omitempty"`
	MilestoneId     int64    `json:"milestone_id,omitempty"`
	PlanId          int64    `json:"plan_id,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	// A map of custom fields values (id => value)
	CustomField map[string]string `json:"custom_field,omitempty"`
}
