package client

type AttachmentHash struct {
	Size     int32  `json:"size,omitempty"`
	Mime     string `json:"mime,omitempty"`
	Filename string `json:"filename,omitempty"`
	Url      string `json:"url,omitempty"`
	Hash     string `json:"hash,omitempty"`
}
