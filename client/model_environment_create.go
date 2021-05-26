package client

type EnvironmentCreate struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Slug        string `json:"slug"`
	Host        string `json:"host,omitempty"`
}
