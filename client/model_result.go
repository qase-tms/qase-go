package client

import (
	"time"
)

type Result struct {
	Hash        string        `json:"hash,omitempty"`
	Comment     string        `json:"comment,omitempty"`
	Stacktrace  string        `json:"stacktrace,omitempty"`
	RunId       int64         `json:"run_id,omitempty"`
	CaseId      int64         `json:"case_id,omitempty"`
	Steps       []ResultSteps `json:"steps,omitempty"`
	Status      string        `json:"status,omitempty"`
	IsApiResult bool          `json:"is_api_result,omitempty"`
	TimeSpentMs int64         `json:"time_spent_ms,omitempty"`
	EndTime     time.Time     `json:"end_time,omitempty"`
	Attachments []Attachment  `json:"attachments,omitempty"`
}
