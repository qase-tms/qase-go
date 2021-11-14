package client

import (
	"time"
)

type Suite struct {
	Id            int64     `json:"id,omitempty"`
	Title         string    `json:"title,omitempty"`
	Description   string    `json:"description,omitempty"`
	Preconditions string    `json:"preconditions,omitempty"`
	Position      int32     `json:"position,omitempty"`
	CasesCount    int32     `json:"cases_count,omitempty"`
	ParentId      int64     `json:"parent_id,omitempty"`
	Created       time.Time `json:"created,omitempty"`
	Updated       time.Time `json:"updated,omitempty"`
}
