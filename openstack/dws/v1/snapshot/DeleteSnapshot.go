package snapshot

type DeleteSnapshotRequest struct {
	// Snapshot ID
	SnapshotId string `json:"snapshot_id"`
}

// DELETE /v1.0/{project_id}/snapshots/{snapshot_id}

type DeleteSnapshotResponse struct {
}
