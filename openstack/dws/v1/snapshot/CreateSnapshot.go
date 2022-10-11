package snapshot

type CreateSnapshotOpts struct {
	// Snapshot object.
	Snapshot Snapshot `json:"snapshot"`
}

type Snapshot struct {
	// Snapshot name, which must be unique and start with a letter.
	// It consists of 4 to 64 characters, which are case-insensitive and contain letters, digits, hyphens (-), and underscores (_) only.
	Name string `json:"name"`
	// ID of the cluster for which you want to create a snapshot. For details about how to obtain the ID, see 7.6 Obtaining the Cluster ID.
	ClusterId string `json:"cluster_id"`
	// Snapshot description. If no value is specified, the description is empty. Enter a maximum of 256 characters.
	// The following special characters are not allowed: !<>'=&"
	Description string `json:"description,omitempty"`
}

// POST /v1.0/{project_id}/snapshots

type CreateSnapshotResponse struct {
	// Snapshot object.
	Snapshot SnapshotResp `json:"snapshot,omitempty"`
}

type SnapshotResp struct {
	// Snapshot ID
	Id string `json:"id,omitempty"`
}
