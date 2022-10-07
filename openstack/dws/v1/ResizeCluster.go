package v1

type ResizeClusterRequest struct {
	//
	ClusterId string `json:"cluster_id"`

	Body ResizeClusterRequestBody `json:"body,omitempty"`
}

type ResizeClusterRequestBody struct {
	ScaleOut ScaleOut `json:"scale_out,omitempty"`
}

type ScaleOut struct {
	//
	Count int32 `json:"count"`
}

// POST /v1.0/{project_id}/clusters/{cluster_id}/resize

type ResizeClusterResponse struct {
}
