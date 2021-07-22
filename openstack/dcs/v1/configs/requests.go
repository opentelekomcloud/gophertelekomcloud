package configs

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Get(client *golangsdk.ServiceClient, instanceID string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, instanceID), &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type UpdateOpts struct {
}
