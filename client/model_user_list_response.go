package client

type UserListResponse struct {
	Status bool                    `json:"status,omitempty"`
	Result *UserListResponseResult `json:"result,omitempty"`
}
