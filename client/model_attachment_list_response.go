package client

type AttachmentListResponse struct {
	Status bool                          `json:"status,omitempty"`
	Result *AttachmentListResponseResult `json:"result,omitempty"`
}
