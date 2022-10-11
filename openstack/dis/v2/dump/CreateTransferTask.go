package dump

type CreateTransferTaskRequest struct {
	// Name of the stream.
	// Maximum: 60
	StreamName string `json:"stream_name"`

	Body *CreateTransferTaskRequest `json:"body,omitempty"`
}

// POST /v2/{project_id}/streams/{stream_name}/transfer-tasks

type CreateTransferTaskResponse struct {
	HttpStatusCode int `json:"-"`
}
