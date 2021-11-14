package client

type ResultCreate struct {
	CaseId      int64               `json:"case_id,omitempty"`
	Case_       *ResultCreateCase   `json:"case,omitempty"`
	Status      string              `json:"status,omitempty"`
	Time        int64               `json:"time,omitempty"`
	TimeMs      int64               `json:"time_ms,omitempty"`
	Defect      bool                `json:"defect,omitempty"`
	Attachments []string            `json:"attachments,omitempty"`
	Stacktrace  string              `json:"stacktrace,omitempty"`
	Comment     string              `json:"comment,omitempty"`
	Param       []string            `json:"param,omitempty"`
	Steps       []ResultCreateSteps `json:"steps,omitempty"`
}
