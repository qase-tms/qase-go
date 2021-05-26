package client

type UserResponse struct {
	Status bool  `json:"status,omitempty"`
	Result *User `json:"result,omitempty"`
}
