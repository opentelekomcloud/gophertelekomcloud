package metadata

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, snapshotId string, key string) (err error) {
	// DELETE /v3/{project_id}/snapshots/{snapshot_id}/metadata/{key}
	_, err = client.Delete(client.ServiceURL("snapshots", snapshotId, "metadata", key), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
