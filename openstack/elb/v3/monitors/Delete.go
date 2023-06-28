package monitors

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular Monitor based on its unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("healthmonitors", id), nil)
	return
}
