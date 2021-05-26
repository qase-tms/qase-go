package client

type Filters4 struct {
	// A single test run result status. Possible values: in_progress, passed, failed, blocked, skipped, invalid.
	Status string `json:"status,omitempty"`
	// A list of run IDs separated by comma.
	Run string `json:"run,omitempty"`
	// A list of case IDs separated by comma.
	CaseId string `json:"case_id,omitempty"`
	// A list of member IDs separated by comma.
	Member string `json:"member,omitempty"`
	Api    bool   `json:"api,omitempty"`
	// Will return all results created after provided datetime. Allowed format: `Y-m-d H:i:s`.
	FromEndTime string `json:"from_end_time,omitempty"`
	// Will return all results created before provided datetime. Allowed format: `Y-m-d H:i:s`.
	ToEndTime string `json:"to_end_time,omitempty"`
}
