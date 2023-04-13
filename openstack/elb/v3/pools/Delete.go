package pools

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular pool based on its unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(client.ServiceURL("pools", id), nil)
	return
}
