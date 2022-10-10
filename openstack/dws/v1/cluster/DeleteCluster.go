package cluster

type DeleteClusterRequest struct {
	// ID of the cluster to be deleted. For details about how to obtain the ID, see 7.6 Obtaining the Cluster ID.
	ClusterId string `json:"cluster_id"`

	Body DeleteClusterRequestBody `json:"body,omitempty"`
}
type DeleteClusterRequestBody struct {
	// The number of latest manual snapshots that need to be retained for a cluster.
	KeepLastManualSnapshot int32 `json:"keep_last_manual_snapshot"`
}

// DELETE /v1.0/{project_id}/clusters/{cluster_id}

type DeleteClusterResponse struct {
}
