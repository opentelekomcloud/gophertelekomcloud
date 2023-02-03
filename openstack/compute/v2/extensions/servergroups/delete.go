package servergroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete requests the deletion of a previously allocated ServerGroup.
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("os-server-groups", id), nil)
	return
}
