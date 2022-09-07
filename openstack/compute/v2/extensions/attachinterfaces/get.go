package attachinterfaces

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get requests details on a single interface attachment by the server and port IDs.
func Get(client *golangsdk.ServiceClient, serverID, portID string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("servers", serverID, "os-interface", portID), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
