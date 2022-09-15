package snapshots

import "github.com/opentelekomcloud/gophertelekomcloud"

// Disable will disable the Snapshot function with the provided ID.
func Disable(client *golangsdk.ServiceClient, clusterId string) (r ErrorResult) {
	_, r.Err = client.Delete(client.ServiceURL("clusters", clusterId, "index_snapshots"), &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
	})
	return
}
