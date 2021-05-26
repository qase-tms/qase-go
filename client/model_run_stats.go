package client

type RunStats struct {
	Total      int32 `json:"total,omitempty"`
	Untested   int32 `json:"untested,omitempty"`
	Passed     int32 `json:"passed,omitempty"`
	Failed     int32 `json:"failed,omitempty"`
	Blocked    int32 `json:"blocked,omitempty"`
	Skipped    int32 `json:"skipped,omitempty"`
	Retest     int32 `json:"retest,omitempty"`
	InProgress int32 `json:"in_progress,omitempty"`
	Invalid    int32 `json:"invalid,omitempty"`
}
