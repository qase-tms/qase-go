package client

type ProjectListResponseResult struct {
	Total    int32     `json:"total,omitempty"`
	Filtered int32     `json:"filtered,omitempty"`
	Count    int32     `json:"count,omitempty"`
	Entities []Project `json:"entities,omitempty"`
}
