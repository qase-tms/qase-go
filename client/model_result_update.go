package client

type ResultUpdate struct {
	Status      string              `json:"status,omitempty"`
	TimeMs      int64               `json:"time_ms,omitempty"`
	Defect      bool                `json:"defect,omitempty"`
	Attachments []string            `json:"attachments,omitempty"`
	Stacktrace  string              `json:"stacktrace,omitempty"`
	Comment     string              `json:"comment,omitempty"`
	Steps       []ResultUpdateSteps `json:"steps,omitempty"`
}
