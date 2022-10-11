package streams

type DeleteStreamRequest struct {
	// Name of the stream to be deleted.
	// Maximum: 60
	StreamName string `json:"stream_name"`
}

// DELETE /v2/{project_id}/streams/{stream_name}

type DeleteStreamResponse struct {
	HttpStatusCode int `json:"-"`
}
