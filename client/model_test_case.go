package client

import (
	"time"
)

type TestCase struct {
	Id             int64                `json:"id,omitempty"`
	Position       int32                `json:"position,omitempty"`
	Title          string               `json:"title,omitempty"`
	Description    string               `json:"description,omitempty"`
	Preconditions  string               `json:"preconditions,omitempty"`
	Postconditions string               `json:"postconditions,omitempty"`
	Severity       int32                `json:"severity,omitempty"`
	Priority       int32                `json:"priority,omitempty"`
	Type_          int32                `json:"type,omitempty"`
	Layer          int32                `json:"layer,omitempty"`
	IsFlaky        int32                `json:"is_flaky,omitempty"`
	Behavior       int32                `json:"behavior,omitempty"`
	Automation     int32                `json:"automation,omitempty"`
	Status         int32                `json:"status,omitempty"`
	MilestoneId    int64                `json:"milestone_id,omitempty"`
	SuiteId        int64                `json:"suite_id,omitempty"`
	CustomFields   []CustomFieldValue   `json:"custom_fields,omitempty"`
	Attachments    []Attachment         `json:"attachments,omitempty"`
	Steps          []TestStep           `json:"steps,omitempty"`
	Params         *AnyOfTestCaseParams `json:"params,omitempty"`
	Created        time.Time            `json:"created,omitempty"`
	Updated        time.Time            `json:"updated,omitempty"`
	Tags           []TagValue           `json:"tags,omitempty"`
	Deleted        time.Time            `json:"deleted,omitempty"`
	MemberId       int64                `json:"member_id,omitempty"`
	ProjectId      int64                `json:"project_id,omitempty"`
}
