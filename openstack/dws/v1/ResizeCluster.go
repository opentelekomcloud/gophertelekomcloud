package v1

type ResizeClusterRequest struct {
	// ID of the cluster to be scaled out. For details about how to obtain the ID, see 7.6 Obtaining the Cluster ID.
	ClusterId string `json:"cluster_id"`

	Body ResizeClusterRequestBody `json:"body,omitempty"`
}

type ResizeClusterRequestBody struct {
	// Scale out an object.
	ScaleOut ScaleOut `json:"scale_out,omitempty"`
}

type ScaleOut struct {
	// Number of nodes to be added
	Count int32 `json:"count"`
}

// POST /v1.0/{project_id}/clusters/{cluster_id}/resize

type ResizeClusterResponse struct {
}
