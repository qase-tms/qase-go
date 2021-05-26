package client

type ProjectCounts struct {
	Cases      int32                 `json:"cases,omitempty"`
	Suites     int32                 `json:"suites,omitempty"`
	Milestones int32                 `json:"milestones,omitempty"`
	Runs       *ProjectCountsRuns    `json:"runs,omitempty"`
	Defects    *ProjectCountsDefects `json:"defects,omitempty"`
}
