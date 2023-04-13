package certificates

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular Certificate based on its unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(client.ServiceURL("certificates", id), nil)
	return
}
