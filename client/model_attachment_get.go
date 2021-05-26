package client

type AttachmentGet struct {
	Hash      string `json:"hash,omitempty"`
	File      string `json:"file,omitempty"`
	Mime      string `json:"mime,omitempty"`
	Size      int32  `json:"size,omitempty"`
	Extension string `json:"extension,omitempty"`
	FullPath  string `json:"full_path,omitempty"`
}
