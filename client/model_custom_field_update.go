package client

type CustomFieldUpdate struct {
	Title string                   `json:"title"`
	Value []CustomFieldCreateValue `json:"value,omitempty"`
	// Dictionary of old values and their replacemants
	ReplaceValues map[string]string `json:"replace_values,omitempty"`
	Placeholder   string            `json:"placeholder,omitempty"`
	DefaultValue  string            `json:"default_value,omitempty"`
	IsFilterable  bool              `json:"is_filterable,omitempty"`
	IsVisible     bool              `json:"is_visible,omitempty"`
	IsRequired    bool              `json:"is_required,omitempty"`
	ProjectsCodes []string          `json:"projects_codes"`
}
