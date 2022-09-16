package tracker

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular tracker.
func Delete(client *golangsdk.ServiceClient) (err error) {
	_, err = client.Delete(client.ServiceURL("tracker"), nil)
	return
}
