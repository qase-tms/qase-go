package client

type PlanDetailedCases struct {
	CaseId     int64 `json:"case_id,omitempty"`
	AssigneeId int64 `json:"assignee_id,omitempty"`
}
