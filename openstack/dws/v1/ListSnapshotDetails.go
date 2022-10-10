package v1

type ListSnapshotDetailsRequest struct {
	// Snapshot ID
	SnapshotId string `json:"snapshot_id"`
}

// GET /v1.0/{project_id}/snapshots/{snapshot_id}

type ListSnapshotDetailsResponse struct {
	// Snapshot object
	Snapshot SnapshotDetail `json:"snapshot,omitempty"`
}

type SnapshotDetail struct {
	// Snapshot ID
	Id string `json:"id"`
	// Snapshot name
	Name string `json:"name"`
	// Snapshot description
	Description string `json:"description"`
	// Time when a snapshot starts to be created. Format: ISO8601: YYYY-MM-DDThh:mm:ssZ
	Started string `json:"started"`
	// Time when a snapshot is complete. Format: ISO8601: YYYY-MM-DDThh:mm:ssZ
	Finished string `json:"finished"`
	// Snapshot size, in GB
	Size float64 `json:"size"`
	// Snapshot status:
	// CREATING
	// AVAILABLE
	// UNAVAILABLE
	Status string `json:"status"`
	// Snapshot type. It can be:
	// MANUAL
	// AUTOMATED
	Type string `json:"type"`
	// ID of the cluster for which snapshots are created For details about how to obtain the ID, see 7.6 Obtaining the Cluster ID.
	ClusterId string `json:"cluster_id"`
}
