package tag

type DeleteClusterTagRequest struct {
	// Resource ID
	ClusterId string `json:"resource_id"`
	// Tag key
	Key string `json:"key"`
}

// DELETE /v1.0/{project_id}/clusters/{resource_id}/tags/{key}

type DeleteClusterTagResponse struct {
}
