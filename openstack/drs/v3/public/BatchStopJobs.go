package public

type BatchPauseJobReq struct {
	// The value cannot contain empty objects. The value of job_id must comply with the UUID rule.
	Jobs []PauseInfo `json:"jobs"`
}

type PauseInfo struct {
	// Task ID.
	JobId string `json:"job_id"`
	// Pause type. target: Stop replay. all: Stop log capturing and replay.
	// Values: target all
	PauseMode string `json:"pause_mode"`
}

// POST /v3/{project_id}/jobs/batch-pause-task

// BatchJobsResponse
