package images

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete implements image delete request.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(client.ServiceURL("images", id), nil)
	return
}
