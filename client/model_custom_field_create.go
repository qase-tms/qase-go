package client

type CustomFieldCreate struct {
	Title string                   `json:"title"`
	Value []CustomFieldCreateValue `json:"value,omitempty"`
	// Possible values: 0 - case; 1 - run; 2 - defect;
	Entity int32 `json:"entity"`
	// Possible values: 0 - number; 1 - string; 2 - text; 3 - selectbox; 4 - checkbox; 5 - radio; 6 - multiselect; 7 - url; 8 - user; 9 - datetime;
	Type_         int32    `json:"type"`
	Placeholder   string   `json:"placeholder,omitempty"`
	DefaultValue  string   `json:"default_value,omitempty"`
	IsFilterable  bool     `json:"is_filterable,omitempty"`
	IsVisible     bool     `json:"is_visible,omitempty"`
	IsRequired    bool     `json:"is_required,omitempty"`
	ProjectsCodes []string `json:"projects_codes"`
}
