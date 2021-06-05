package client

type DefectUpdate struct {
	Title        string   `json:"title,omitempty"`
	ActualResult string   `json:"actual_result,omitempty"`
	Severity     int32    `json:"severity,omitempty"`
	MilestoneId  int64    `json:"milestone_id,omitempty"`
	Attachments  []string `json:"attachments,omitempty"`
	// A map of custom fields values (id => value)
	CustomField map[string]string `json:"custom_field,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
}
