package snapshots

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will delete the existing Snapshot with the provided ID.
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v3/{project_id}/snapshots/{snapshot_id}
	_, err = client.Delete(client.ServiceURL("snapshots", id), nil)
	return
}
