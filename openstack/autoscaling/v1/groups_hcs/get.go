package groups_hcs

import "github.com/opentelekomcloud/gophertelekomcloud"

// GetGroup is a method of getting the detailed information of the group by id
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("scaling_group", id), &r.Body, nil)
	return
}
