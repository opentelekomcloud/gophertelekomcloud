package v1

type DeleteClusterRequest struct {
	//
	ClusterId string `json:"cluster_id"`

	Body DeleteClusterRequestBody `json:"body,omitempty"`
}
type DeleteClusterRequestBody struct {
	//
	KeepLastManualSnapshot int32 `json:"keep_last_manual_snapshot"`
}

// DELETE /v1.0/{project_id}/clusters/{cluster_id}

type DeleteClusterResponse struct {
}
