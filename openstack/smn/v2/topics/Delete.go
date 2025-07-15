package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete a topic via id
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("topics", id), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
