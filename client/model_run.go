package client

import (
	"time"
)

type Run struct {
	Id          int64     `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Status      int32     `json:"status,omitempty"`
	StatusText  string    `json:"status_text,omitempty"`
	StartTime   time.Time `json:"start_time,omitempty"`
	EndTime     time.Time `json:"end_time,omitempty"`
	Public      bool      `json:"public,omitempty"`
	Stats       *RunStats `json:"stats,omitempty"`
	// Time in ms.
	TimeSpent    int64              `json:"time_spent,omitempty"`
	Environment  *RunEnvironment    `json:"environment,omitempty"`
	Milestone    *RunMilestone      `json:"milestone,omitempty"`
	CustomFields []CustomFieldValue `json:"custom_fields,omitempty"`
	Tags         []TagValue         `json:"tags,omitempty"`
	Cases        []int64            `json:"cases,omitempty"`
}
