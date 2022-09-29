package public

type BatchSpecialTestConnectionOpts struct {
	Jobs []BatchJobActionReq `json:"jobs"`
}

type BatchJobActionReq struct {
	// Specific operation to be performed.
	Action string `json:"action"`
	// Task ID. (In cluster mode, the value is the ID of the parent task.).
	JobId string `json:"job_id"`
	// The parameter that corresponds to the operation.
	// Please refer to Docs API for details.
	Property string `json:"property"`
}

// POST /v3/{project_id}/jobs/cluster/batch-connection

type BatchValidateClustersConnectionsResponse struct {
	Results []CheckJobResp `json:"results,omitempty"`
	Count   int32          `json:"count,omitempty"`
}
