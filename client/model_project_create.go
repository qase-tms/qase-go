package client

type ProjectCreate struct {
	// Project title.
	Title string `json:"title"`
	// Project code. Unique for team. Digits and special characters are not allowed.
	Code string `json:"code"`
	// Project description.
	Description string `json:"description,omitempty"`
	Access      string `json:"access,omitempty"`
	// Team group hash. Required if access param is set to group.
	Group string `json:"group,omitempty"`
}
