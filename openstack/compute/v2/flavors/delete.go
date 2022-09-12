package flavors

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete deletes the specified flavor ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	raw, err := client.Delete(client.ServiceURL("flavors", id), nil)
	return
}
