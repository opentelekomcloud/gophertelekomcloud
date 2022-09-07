package attachinterfaces

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete makes a request against the nova API to detach a single interface from the server.
// It needs server and port IDs to make a such request.
func Delete(client *golangsdk.ServiceClient, serverID, portID string) (r DeleteResult) {
	raw, err := client.Delete(deleteInterfaceURL(client, serverID, portID), nil)
	return
}
