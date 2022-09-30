package public

type BatchStartJobReq struct {
	// Request list for starting tasks in batches.
	Jobs []StartInfo `json:"jobs"`
}

type StartInfo struct {
	// Task ID.
	JobId string `json:"job_id"`
	// Task start time. The timestamp is accurate to milliseconds, for example, 1608188903063.
	// If the value is empty, the task is started immediately.
	StartTime string `json:"start_time,omitempty"`
}

// POST /v3/{project_id}/jobs/batch-starting

type BatchStartJobsResponse struct {
	Results []IdJobResp `json:"results,omitempty"`
	Count   int32       `json:"count,omitempty"`
}

type IdJobResp struct {
	// Task ID.
	Id string `json:"id"`
	// Status Values: success failed
	Status string `json:"status,omitempty"`
	// Error code, which is optional and indicates the returned information about the failure status.
	ErrorCode string `json:"error_code,omitempty"`
	// Error message, which is optional and indicates the returned information about the failure status.
	ErrorMsg string `json:"error_msg,omitempty"`
}
