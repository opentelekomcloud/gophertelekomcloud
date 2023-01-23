package snapshot

import (
	"fmt"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dws/v1/cluster"
)

type Snapshot struct {
	// Snapshot name, which must be unique and start with a letter.
	// It consists of 4 to 64 characters, which are case-insensitive and contain letters, digits, hyphens (-), and underscores (_) only.
	Name string `json:"name" required:"true"`
	// ID of the cluster for which you want to create a snapshot. For details about how to obtain the ID, see 7.6 Obtaining the Cluster ID.
	ClusterId string `json:"cluster_id" required:"true"`
	// Snapshot description. If no value is specified, the description is empty. Enter a maximum of 256 characters.
	// The following special characters are not allowed: !<>'=&"
	Description string `json:"description,omitempty"`
}

func CreateSnapshot(client *golangsdk.ServiceClient, opts Snapshot) (string, error) {
	b, err := build.RequestBody(opts, "snapshot")
	if err != nil {
		return "", err
	}

	// POST /v1.0/{project_id}/snapshots
	raw, err := client.Post(client.ServiceURL("snapshots"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		Id string `json:"id,omitempty"`
	}
	err = extract.IntoStructPtr(raw.Body, &res, "snapshot")
	return res.Id, err
}

func WaitForSnapshot(c *golangsdk.ServiceClient, cid, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		current, err := cluster.ListClusterDetails(c, cid)
		if err != nil {
			return false, err
		}

		curSnap, err := ListSnapshotDetails(c, id)
		if err != nil {
			return false, err
		}

		if curSnap.Status == "AVAILABLE" && current.Status == "AVAILABLE" && current.TaskStatus != "SNAPSHOTTING" {
			return true, nil
		}

		if curSnap.Status == "UNAVAILABLE" {
			return false, fmt.Errorf("snapshot creation failed: " + current.FailedReasons.ErrorMsg)
		}

		time.Sleep(10 * time.Second)

		return false, nil
	})
}
