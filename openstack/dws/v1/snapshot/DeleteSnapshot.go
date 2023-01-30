package snapshot

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

func DeleteSnapshot(client *golangsdk.ServiceClient, snapshotId string) error {
	// DELETE /v1.0/{project_id}/snapshots/{snapshot_id}
	_, err := client.Delete(client.ServiceURL("snapshots", snapshotId), openstack.StdRequestOpts())
	return err
}
