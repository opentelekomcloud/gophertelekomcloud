package public

type BatchReplaceDefinerReq struct {
	Jobs []ReplaceDefinerInfo `json:"jobs"`
}

type ReplaceDefinerInfo struct {
	// Task ID.
	JobId string `json:"job_id"`
	// Whether to replace the definer with the destination database user.
	ReplaceDefiner bool `json:"replace_definer"`
}

// POST /v3/{project_id}/jobs/batch-replace-definer

// BatchJobsResponse
