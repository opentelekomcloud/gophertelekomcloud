package groups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete is a method of deleting a group by group id
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}
