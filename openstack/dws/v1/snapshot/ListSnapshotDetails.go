package snapshot

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

func ListSnapshotDetails(client *golangsdk.ServiceClient, snapshotId string) (*SnapshotDetail, error) {
	// GET /v1.0/{project_id}/snapshots/{snapshot_id}
	raw, err := client.Get(client.ServiceURL("snapshots", snapshotId), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res SnapshotDetail
	err = extract.IntoStructPtr(raw.Body, &res, "snapshot")
	return &res, err
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
