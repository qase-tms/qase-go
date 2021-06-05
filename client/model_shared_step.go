package client

import (
	"time"
)

type SharedStep struct {
	Hash           string            `json:"hash,omitempty"`
	Title          string            `json:"title,omitempty"`
	Action         string            `json:"action,omitempty"`
	ExpectedResult string            `json:"expected_result,omitempty"`
	Steps          []SharedStepSteps `json:"steps,omitempty"`
	Data           string            `json:"data,omitempty"`
	Cases          []int64           `json:"cases,omitempty"`
	CasesCount     int32             `json:"cases_count,omitempty"`
	Created        time.Time         `json:"created,omitempty"`
	Updated        time.Time         `json:"updated,omitempty"`
}
