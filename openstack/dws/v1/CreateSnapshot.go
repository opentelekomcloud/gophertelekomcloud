package v1

type CreateSnapshotRequestBody struct {
	Snapshot Snapshot `json:"snapshot"`
}

type Snapshot struct {
	//
	Name string `json:"name"`
	//
	ClusterId string `json:"cluster_id"`
	//
	Description string `json:"description,omitempty"`
}

// POST /v1.0/{project_id}/snapshots

type CreateSnapshotResponse struct {
	Snapshot SnapshotResp `json:"snapshot,omitempty"`
}

type SnapshotResp struct {
	// Snapshot ID
	Id string `json:"id,omitempty"`
}
