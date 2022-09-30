package public

type BatchPrecheckReq struct {
	Jobs []PreCheckInfo `json:"jobs"`
}

type PreCheckInfo struct {
	// Task ID.
	JobId string `json:"job_id"`
	// Pre-check mode. Value: forStartJob
	PrecheckMode string `json:"precheck_mode"`
}

// POST /v3/{project_id}/jobs/batch-precheck

type BatchCheckJobsResponse struct {
	Results []PostPreCheckResp `json:"results,omitempty"`
	Count   int32              `json:"count,omitempty"`
}

type PostPreCheckResp struct {
	// Task ID
	Id string `json:"id,omitempty"`
	// Pre-check ID.
	PrecheckId string `json:"precheck_id,omitempty"`
	// Success or failure status.
	// Values: success failed
	Status string `json:"status,omitempty"`
	// Error code, which is optional and indicates the returned information about the failure status.
	ErrorCode string `json:"error_code,omitempty"`
	// Error message, which is optional and indicates the returned information about the failure status.
	ErrorMsg string `json:"error_msg,omitempty"`
}
