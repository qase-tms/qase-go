package client

type CustomFieldsResponse struct {
	Status bool                        `json:"status,omitempty"`
	Result *CustomFieldsResponseResult `json:"result,omitempty"`
}
