package client

import (
	"time"
)

type CustomField struct {
	Id            int64     `json:"id,omitempty"`
	Title         string    `json:"title,omitempty"`
	Entity        string    `json:"entity,omitempty"`
	Type_         string    `json:"type,omitempty"`
	Placeholder   string    `json:"placeholder,omitempty"`
	DefaultValue  string    `json:"default_value,omitempty"`
	Value         string    `json:"value,omitempty"`
	IsRequired    bool      `json:"is_required,omitempty"`
	IsVisible     bool      `json:"is_visible,omitempty"`
	IsFilterable  bool      `json:"is_filterable,omitempty"`
	Created       time.Time `json:"created,omitempty"`
	Updated       time.Time `json:"updated,omitempty"`
	ProjectsCodes []string  `json:"projects_codes,omitempty"`
}
