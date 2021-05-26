package client

import (
	"time"
)

type Defect struct {
	Id           int32              `json:"id,omitempty"`
	Title        string             `json:"title,omitempty"`
	ActualResult string             `json:"actual_result,omitempty"`
	Severity     string             `json:"severity,omitempty"`
	Status       string             `json:"status,omitempty"`
	MilestoneId  int32              `json:"milestone_id,omitempty"`
	CustomFields []CustomFieldValue `json:"custom_fields,omitempty"`
	Attachments  []Attachment       `json:"attachments,omitempty"`
	Created      time.Time          `json:"created,omitempty"`
	Updated      time.Time          `json:"updated,omitempty"`
	Deleted      time.Time          `json:"deleted,omitempty"`
	Resolved     time.Time          `json:"resolved,omitempty"`
	ProjectId    int32              `json:"project_id,omitempty"`
	MemberId     int32              `json:"member_id,omitempty"`
	ExternalData string             `json:"external_data,omitempty"`
	Tags         []TagValue         `json:"tags,omitempty"`
}
