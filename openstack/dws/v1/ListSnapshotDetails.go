package v1

type ListSnapshotDetailsRequest struct {
	//
	SnapshotId string `json:"snapshot_id"`
}

// GET /v1.0/{project_id}/snapshots/{snapshot_id}

type ListSnapshotDetailsResponse struct {
	Snapshot SnapshotDetail `json:"snapshot,omitempty"`
}

type SnapshotDetail struct {
	//
	Id string `json:"id"`
	//
	Name string `json:"name"`
	//
	Description string `json:"description"`
	//
	Started string `json:"started"`
	//
	Finished string `json:"finished"`
	//
	Size float64 `json:"size"`
	//
	Status string `json:"status"`
	//
	Type string `json:"type"`
	//
	ClusterId string `json:"cluster_id"`
}
