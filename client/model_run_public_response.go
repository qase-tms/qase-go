package client

type RunPublicResponse struct {
	Status bool                     `json:"status,omitempty"`
	Result *RunPublicResponseResult `json:"result,omitempty"`
}
