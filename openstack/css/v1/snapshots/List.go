package snapshots

import "github.com/opentelekomcloud/gophertelekomcloud"

// List retrieves the Snapshots with the provided ID. To extract the Snapshot
// objects from the response, call the Extract method on the GetResult.
func List(client *golangsdk.ServiceClient, clusterId string) (r ListResult) {
	_, r.Err = client.Get(client.ServiceURL("clusters", clusterId, "index_snapshots"), &r.Body, nil)
	return
}
