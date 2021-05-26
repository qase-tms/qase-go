package client

// Could be used insetad of case_id
type ResultCreateCase struct {
	Title       string `json:"title,omitempty"`
	SuiteTitle  string `json:"suite_title,omitempty"`
	Description string `json:"description,omitempty"`
	Layer       string `json:"layer,omitempty"`
	Severity    string `json:"severity,omitempty"`
}
