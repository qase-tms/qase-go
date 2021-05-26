package client

type RunEnvironment struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Host        string `json:"host,omitempty"`
}
