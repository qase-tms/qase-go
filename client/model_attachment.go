package client

type Attachment struct {
	Size     int32  `json:"size,omitempty"`
	Mime     string `json:"mime,omitempty"`
	Filename string `json:"filename,omitempty"`
	Url      string `json:"url,omitempty"`
}
