package client

type TestCaseListResponseResult struct {
	Total    int32      `json:"total,omitempty"`
	Filtered int32      `json:"filtered,omitempty"`
	Count    int32      `json:"count,omitempty"`
	Entities []TestCase `json:"entities,omitempty"`
}
