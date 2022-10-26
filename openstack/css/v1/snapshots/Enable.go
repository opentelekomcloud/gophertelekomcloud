package snapshots

import "github.com/opentelekomcloud/gophertelekomcloud"

// Enable will enable the Snapshot function with the provided ID.
func Enable(client *golangsdk.ServiceClient, clusterId string) (err error) {
	_, err = client.Post(client.ServiceURL("clusters", clusterId, "index_snapshot/auto_setting"), nil, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
	})
	return
}
