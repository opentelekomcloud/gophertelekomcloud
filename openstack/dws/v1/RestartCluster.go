package v1

type RestartClusterRequest struct {
	//
	ClusterId string `json:"cluster_id"`

	Body RestartClusterRequestBody `json:"body,omitempty"`
}

type RestartClusterRequestBody struct {
	//
	Restart interface{} `json:"restart"`
}

// POST /v1.0/{project_id}/clusters/{cluster_id}/restart

type RestartClusterResponse struct {
}
