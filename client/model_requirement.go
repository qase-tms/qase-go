package client

import (
	"time"
)

type Requirement struct {
	Id          int32     `json:"id,omitempty"`
	ParentId    int32     `json:"parent_id,omitempty"`
	MemberId    int32     `json:"member_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status,omitempty"`
	Type_       string    `json:"type,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	Updated     time.Time `json:"updated,omitempty"`
}
