package client

type Filters struct {
	// Provide a string that will be used to search by name.
	Search string `json:"search,omitempty"`
	// ID of milestone.
	MilestoneId int32 `json:"milestone_id,omitempty"`
	// ID of test suite.
	SuiteId int32 `json:"suite_id,omitempty"`
	// A list of severity values separated by comma. Possible values: undefined, blocker, critical, major, normal, minor, trivial
	Severity string `json:"severity,omitempty"`
	// A list of priority values separated by comma. Possible values: undefined, high, medium, low
	Priority string `json:"priority,omitempty"`
	// A list of type values separated by comma. Possible values: other, functional smoke, regression, security, usability, performance, acceptance
	Type_ string `json:"type,omitempty"`
	// A list of behavior values separated by comma. Possible values: undefined, positive negative, destructive
	Behavior string `json:"behavior,omitempty"`
	// A list of values separated by comma. Possible values: is-not-automated, automated to-be-automated
	Automation string `json:"automation,omitempty"`
	// A list of values separated by comma. Possible values: actual, draft deprecated
	Status string `json:"status,omitempty"`
}
