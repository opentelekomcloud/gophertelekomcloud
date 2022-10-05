package public

// BatchQueryJobReqPage

// POST /v3/{project_id}/jobs/batch-status

type BatchListJobStatusResponse struct {
	// 任务状态信息
	Results []QueryJobStatusResp `json:"results,omitempty"`
	// 返回任务数量
	Count int32 `json:"count,omitempty"`
}

type QueryJobStatusResp struct {
	// Task ID.
	Id string `json:"id,omitempty"`
	// Task status. Values:
	// CREATING: The task is being created.
	// CREATE_FAILED: The task fails to be created.
	// CONFIGURATION: The task is being configured.
	// STARTJOBING: The task is being started.
	// WAITING_FOR_START: The task is waiting to be started.
	// START_JOB_FAILED: The task fails to be started.
	// PAUSEING: The task is being paused.
	// FULL_TRANSFER_STARTED: Full migration is in progress, and the DR scenario is initialized.
	// FULL_TRANSFER_FAILED: Full migration failed. Initialization failed in the DR scenario.
	// FULL_TRANSFER_COMPLETE: Full migration is complete, and the initialization is complete in the DR scenario.
	// INCRE_TRANSFER_STARTED: Incremental migration is being performed, and the DR task is in progress.
	// INCRE_TRANSFER_FAILED: Incremental migration fails and a DR exception occurs.
	// RELEASE_RESOURCE_STARTED: The task is being stopped.
	// RELEASE_RESOURCE_FAILED: Stop task failed.
	// RELEASE_RESOURCE_COMPLETE: The task is stopped.
	// CHANGE_JOB_STARTED: The task is being changed.
	// CHANGE_JOB_FAILED: Change task failed.
	// CHILD_TRANSFER_STARTING: The subtask is being started.
	// CHILD_TRANSFER_STARTED: The subtask is being migrated.
	// CHILD_TRANSFER_COMPLETE: The subtask migration is complete.
	// CHILD_TRANSFER_FAILED: Migrate subtask failed.
	// RELEASE_CHILD_TRANSFER_STARTED: The subtask is being stopped.
	// RELEASE_CHILD_TRANSFER_COMPLETE: The subtask is complete.
	Status string `json:"status,omitempty"`

	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}
