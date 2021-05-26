package client

type AttachmentUploadsResponse struct {
	Status bool            `json:"status,omitempty"`
	Result []AttachmentGet `json:"result,omitempty"`
}
