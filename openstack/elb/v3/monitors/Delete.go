package monitors

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular Monitor based on its unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(client.ServiceURL("healthmonitors", id), nil)
	return
}
