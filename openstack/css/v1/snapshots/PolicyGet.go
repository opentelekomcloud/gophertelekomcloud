package snapshots

import "github.com/opentelekomcloud/gophertelekomcloud"

// PolicyGet retrieves the snapshot policy with the provided cluster ID.
// To extract the snapshot policy object from the response, call the Extract method on the GetResult.
func PolicyGet(client *golangsdk.ServiceClient, clusterId string) (r PolicyResult) {
	raw, err = client.Get(client.ServiceURL("clusters", clusterId, "index_snapshot/policy"), nil, nil)
	return
}
