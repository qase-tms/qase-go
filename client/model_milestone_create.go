package client

type MilestoneCreate struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
	// unix timestamp
	DueDate int32 `json:"due_date,omitempty"`
}
