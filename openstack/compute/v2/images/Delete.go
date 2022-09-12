package images

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete deletes the specified image ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	raw, err := client.Delete(client.ServiceURL("images", id), nil)
	return
}
