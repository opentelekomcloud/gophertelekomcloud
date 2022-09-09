package attachinterfaces

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete makes a request against the nova API to detach a single interface from the server.
// It needs server and port IDs to make a such request.
func Delete(client *golangsdk.ServiceClient, serverID, portID string) (err error) {
	_, err = client.Delete(client.ServiceURL("servers", serverID, "os-interface", portID), nil)
	return
}
