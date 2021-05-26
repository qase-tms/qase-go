package client

type CustomFieldsResponseResult struct {
	Total    int32         `json:"total,omitempty"`
	Filtered int32         `json:"filtered,omitempty"`
	Count    int32         `json:"count,omitempty"`
	Entities []CustomField `json:"entities,omitempty"`
}
