package clusters

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular cluster based on its unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("clusters", id), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts, JSONBody: nil,
	})

	return
}
