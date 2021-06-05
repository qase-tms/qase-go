package client

type TestCaseUpdate struct {
	Description    string                `json:"description,omitempty"`
	Preconditions  string                `json:"preconditions,omitempty"`
	Postconditions string                `json:"postconditions,omitempty"`
	Title          string                `json:"title,omitempty"`
	Severity       int32                 `json:"severity,omitempty"`
	Priority       int32                 `json:"priority,omitempty"`
	Behavior       int32                 `json:"behavior,omitempty"`
	Type_          int32                 `json:"type,omitempty"`
	Layer          int32                 `json:"layer,omitempty"`
	IsFlaky        int32                 `json:"is_flaky,omitempty"`
	SuiteId        int64                 `json:"suite_id,omitempty"`
	MilestoneId    int64                 `json:"milestone_id,omitempty"`
	Automation     int32                 `json:"automation,omitempty"`
	Status         int32                 `json:"status,omitempty"`
	Steps          []TestCaseUpdateSteps `json:"steps,omitempty"`
	// A map of custom fields values (id => value)
	CustomField map[string]string `json:"custom_field,omitempty"`
}
