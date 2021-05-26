package client

type AttachmentResponse struct {
	Status bool           `json:"status,omitempty"`
	Result *AttachmentGet `json:"result,omitempty"`
}
