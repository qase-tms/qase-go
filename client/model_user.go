package client

type User struct {
	Id     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Title  string `json:"title,omitempty"`
	Status int32  `json:"status,omitempty"`
}
