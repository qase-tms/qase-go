package client

import (
	"time"
)

type Result struct {
	Hash        string        `json:"hash,omitempty"`
	Comment     string        `json:"comment,omitempty"`
	Stacktrace  string        `json:"stacktrace,omitempty"`
	RunId       int32         `json:"run_id,omitempty"`
	CaseId      int32         `json:"case_id,omitempty"`
	Steps       []ResultSteps `json:"steps,omitempty"`
	Status      string        `json:"status,omitempty"`
	IsApiResult bool          `json:"is_api_result,omitempty"`
	TimeSpentMs int32         `json:"time_spent_ms,omitempty"`
	EndTime     time.Time     `json:"end_time,omitempty"`
	Attachments []Attachment  `json:"attachments,omitempty"`
}
