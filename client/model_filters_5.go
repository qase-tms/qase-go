package client

type Filters5 struct {
	// A list of status values separated by comma. Possible values: active, complete, abort.
	Status      string `json:"status,omitempty"`
	Milestone   int32  `json:"milestone,omitempty"`
	Environment int32  `json:"environment,omitempty"`
}
