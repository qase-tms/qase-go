package client

import (
	"time"
)

type PlanDetailed struct {
	Id          int64               `json:"id,omitempty"`
	Title       string              `json:"title,omitempty"`
	Description string              `json:"description,omitempty"`
	CasesCount  int32               `json:"cases_count,omitempty"`
	Created     time.Time           `json:"created,omitempty"`
	Updated     time.Time           `json:"updated,omitempty"`
	AverageTime float64             `json:"average_time,omitempty"`
	Cases       []PlanDetailedCases `json:"cases,omitempty"`
}
