package public

type BatchRetryReq struct {
	Jobs []RetryInfo `json:"jobs"`
}

type RetryInfo struct {
	// Task ID.
	JobId string `json:"job_id"`
	// This parameter is mandatory and must be set to true.
	IsSyncReEdit bool `json:"is_sync_re_edit,omitempty"`
}

// POST /v3/{project_id}/jobs/batch-retry-task

type BatchRestoreTaskResponse struct {
	Results []IdJobResp `json:"results,omitempty"`
	Count   int32       `json:"count,omitempty"`
}
