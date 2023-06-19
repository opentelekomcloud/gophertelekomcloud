package snapshots

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get retrieves the Snapshot with the provided ID. To extract the Snapshot
// object from the response, call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, id string) (*Snapshot, error) {
	// GET /v3/{project_id}/snapshots/{snapshot_id}
	raw, err := client.Get(client.ServiceURL("snapshots", id), nil, nil)
	return extra(err, raw)
}
