package v1

type ListSnapshotsRequest struct {
}

// GET /v1.0/{project_id}/snapshots

type ListSnapshotsResponse struct {
	//
	Snapshots []Snapshots `json:"snapshots,omitempty"`
	//
	Count int32 `json:"count,omitempty"`
}

type Snapshots struct {
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
