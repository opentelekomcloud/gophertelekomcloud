package v1

type DeleteSnapshotRequest struct {
	//
	SnapshotId string `json:"snapshot_id"`
}

// DELETE /v1.0/{project_id}/snapshots/{snapshot_id}

type DeleteSnapshotResponse struct {
}
