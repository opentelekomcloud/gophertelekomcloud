package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("instance", id), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
