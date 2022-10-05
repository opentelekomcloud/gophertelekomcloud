package public

type BatchQueryProgressReq struct {
	// Request for querying task progress in batches.
	Jobs []string `json:"jobs"`
}

// POST /v3/{project_id}/jobs/batch-progress

type BatchListProgressesResponse struct {
	Count   int32               `json:"count,omitempty"`
	Results []QueryProgressResp `json:"results,omitempty"`
}

type QueryProgressResp struct {
	// Task ID.
	JobId string `json:"job_id,omitempty"`
	// Migration percentage.
	Progress string `json:"progress,omitempty"`
	// Incremental migration delay.
	IncreTransDelay string `json:"incre_trans_delay,omitempty"`
	// Migration type. Values:
	// FULL_TRANS: full migration
	// INCR_TRANS: incremental migration
	// FULL_INCR_TRANS: full+incremental migration
	TaskMode string `json:"task_mode,omitempty"`
	// Task status.
	TransferStatus string `json:"transfer_status,omitempty"`
	// Migration time in timestamp format.
	ProcessTime string `json:"process_time,omitempty"`
	// Estimated remaining time.
	RemainingTime string `json:"remaining_time,omitempty"`
	// Data, structure, and index migration progress information body.
	ProgressMap map[string]ProgressInfo `json:"progress_map,omitempty"`

	ErrorCode string `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`
}

type ProgressInfo struct {
	// Progress.
	Completed string `json:"completed,omitempty"`
	// Estimated remaining time.
	RemainingTime string `json:"remaining_time,omitempty"`
}