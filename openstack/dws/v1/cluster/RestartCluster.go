package cluster

type RestartClusterRequest struct {
	// ID of the cluster to be restarted. For details about how to obtain the ID, see 7.6 Obtaining the Cluster ID.
	ClusterId string `json:"cluster_id"`

	Body RestartClusterRequestBody `json:"body,omitempty"`
}

type RestartClusterRequestBody struct {
	// Restart flag.
	Restart interface{} `json:"restart"`
}

// POST /v1.0/{project_id}/clusters/{cluster_id}/restart

type RestartClusterResponse struct {
}
