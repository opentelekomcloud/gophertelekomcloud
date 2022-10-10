package tag

type BatchCreateClusterTagsRequest struct {
	// Resource ID, for example, 7d85f602-a948-4a30-afd4-e84f47471c15.
	ClusterId string `json:"resource_id"`

	Body BatchCreateClusterTagsReq `json:"body,omitempty"`
}

type BatchCreateClusterTagsReq struct {
	// Identifies the operation. The value can be create or delete.
	// create: adds tags in batches.
	// delete: deletes tags in batches.
	Action string `json:"action"`
	// Tag list.
	Tags []Tag `json:"tags"`
}

// POST /v1.0/{project_id}/clusters/{resource_id}/tags/action
