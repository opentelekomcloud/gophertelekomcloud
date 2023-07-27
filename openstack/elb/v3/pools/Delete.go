package pools

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular pool based on its unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("pools", id), nil)
	return
}
