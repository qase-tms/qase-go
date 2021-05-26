package client

type ProjectCountsDefects struct {
	Total int32 `json:"total,omitempty"`
	Open  int32 `json:"open,omitempty"`
}
