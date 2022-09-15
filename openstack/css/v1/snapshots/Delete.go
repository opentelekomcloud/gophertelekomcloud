package snapshots

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will delete the existing Snapshot ID with the provided ID.
func Delete(client *golangsdk.ServiceClient, clusterId, id string) (r ErrorResult) {
	_, r.Err = client.Delete(client.ServiceURL("clusters", clusterId, "index_snapshot", id), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return
}
