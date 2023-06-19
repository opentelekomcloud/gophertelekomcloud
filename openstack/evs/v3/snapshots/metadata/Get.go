package metadata

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Get If metadata contains the __system__enableActive field, the snapshot is automatically created during the backup of a server.
func Get(client *golangsdk.ServiceClient, snapshotId string) (map[string]string, error) {
	// GET /v3/{project_id}/snapshots/{snapshot_id}/metadata
	raw, err := client.Get(client.ServiceURL("snapshots", snapshotId, "metadata"), nil, nil)
	return extraMetadata(err, raw)
}
