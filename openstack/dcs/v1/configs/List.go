package configs

import "github.com/opentelekomcloud/gophertelekomcloud"

func List(client *golangsdk.ServiceClient, instanceID string) (r ListResult) {
	_, r.Err = client.Get(client.ServiceURL("instances", instanceID, "configs"), &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
