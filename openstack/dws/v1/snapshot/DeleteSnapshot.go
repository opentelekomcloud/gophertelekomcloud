package snapshot

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func DeleteSnapshot(client *golangsdk.ServiceClient, snapshotId string) error {
	// DELETE /v1.0/{project_id}/snapshots/{snapshot_id}
	_, err := client.Delete(client.ServiceURL("snapshots", snapshotId), nil)
	return err
}
