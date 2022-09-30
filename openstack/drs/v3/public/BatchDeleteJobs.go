package public

type BatchDeleteJobReq struct {
	Jobs []DeleteJobReq `json:"jobs"`
}

type DeleteJobReq struct {
	// The value can be terminated, force_terminate, or delete. terminate indicates that the migration task is stopped,
	// force_terminate indicates that the migration task is forcibly stopped, and delete indicates that the migration task is deleted.
	// Values: terminate force_terminate delete
	DeleteType string `json:"delete_type"`
	// Task ID.
	JobId string `json:"job_id"`
}

type BatchDeleteJobsResponse struct {
	Results []IdJobResp `json:"results,omitempty"`
	Count   int32       `json:"count,omitempty"`
}

// DELETE /v3/{project_id}/jobs/batch-jobs

// BatchJobsResponse
