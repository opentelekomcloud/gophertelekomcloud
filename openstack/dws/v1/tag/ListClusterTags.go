package tag

type ListClusterTagsRequest struct {
	// Resource ID, for example, 7d85f602-a948-4a30-afd4-e84f47471c15.
	ClusterId string `json:"resource_id"`
}

// GET /v1.0/{project_id}/clusters/{cluster_id}/tags

type ListClusterTagsResponse struct {
	// Tag list.
	Tags           *[]TagPlain `json:"tags,omitempty"`
	HttpStatusCode int         `json:"-"`
}
