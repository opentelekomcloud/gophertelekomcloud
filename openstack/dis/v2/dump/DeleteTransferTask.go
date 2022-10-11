package dump

type DeleteTransferTaskRequest struct {
	// Name of the stream.
	StreamName string `json:"stream_name"`
	// Name of the dump task to be deleted.
	TaskName string `json:"task_name"`
}

// DELETE /v2/{project_id}/streams/{stream_name}/transfer-tasks/{task_name}

type DeleteTransferTaskResponse struct {
	HttpStatusCode int `json:"-"`
}
