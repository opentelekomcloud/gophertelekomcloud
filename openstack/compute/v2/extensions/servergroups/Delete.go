package servergroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete requests the deletion of a previously allocated ServerGroup.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	raw, err := client.Delete(deleteURL(client, id), nil)
	return
}
