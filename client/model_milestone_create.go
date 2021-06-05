package client

type MilestoneCreate struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
	// unix timestamp
	DueDate int64 `json:"due_date,omitempty"`
}
