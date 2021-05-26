package client

type CustomFieldResponse struct {
	Status bool         `json:"status,omitempty"`
	Result *CustomField `json:"result,omitempty"`
}
