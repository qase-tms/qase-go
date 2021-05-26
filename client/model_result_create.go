package client

type ResultCreate struct {
	CaseId      int32               `json:"case_id,omitempty"`
	Case_       *ResultCreateCase   `json:"case,omitempty"`
	Status      string              `json:"status,omitempty"`
	Time        int32               `json:"time,omitempty"`
	TimeMs      int32               `json:"time_ms,omitempty"`
	Defect      bool                `json:"defect,omitempty"`
	Attachments []string            `json:"attachments,omitempty"`
	Stacktrace  string              `json:"stacktrace,omitempty"`
	Comment     string              `json:"comment,omitempty"`
	Param       []string            `json:"param,omitempty"`
	Steps       []ResultCreateSteps `json:"steps,omitempty"`
}
