package groups_hcs

import "github.com/opentelekomcloud/gophertelekomcloud"

// DeleteGroup is a method of deleting a group by group id
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(client.ServiceURL("scaling_group", id), nil)
	return
}
