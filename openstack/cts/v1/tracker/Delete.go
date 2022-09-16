package tracker

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular tracker.
func Delete(client *golangsdk.ServiceClient) (r DeleteResult) {
	raw, err := client.Delete(client.ServiceURL("tracker"), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
