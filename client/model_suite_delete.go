package client

type SuiteDelete struct {
	// If provided, child test cases would be moved to suite with such ID.
	DestinationId int32 `json:"destination_id,omitempty"`
}
