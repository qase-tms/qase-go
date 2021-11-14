package client

type Environment struct {
	Id          int64  `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Host        string `json:"host,omitempty"`
}
