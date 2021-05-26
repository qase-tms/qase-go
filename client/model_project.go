package client

type Project struct {
	Title  string         `json:"title,omitempty"`
	Code   string         `json:"code,omitempty"`
	Counts *ProjectCounts `json:"counts,omitempty"`
}
