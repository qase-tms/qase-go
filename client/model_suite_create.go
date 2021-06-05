package client

type SuiteCreate struct {
	// Test suite title.
	Title string `json:"title"`
	// Test suite description.
	Description string `json:"description,omitempty"`
	// Test suite preconditions
	Preconditions string `json:"preconditions,omitempty"`
	// Parent suite ID
	ParentId int64 `json:"parent_id,omitempty"`
}
