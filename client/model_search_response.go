package client

type SearchResponse struct {
	Status bool                  `json:"status,omitempty"`
	Result *SearchResponseResult `json:"result,omitempty"`
}
