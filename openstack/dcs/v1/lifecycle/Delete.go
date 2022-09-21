package lifecycle

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete an instance by id
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}
