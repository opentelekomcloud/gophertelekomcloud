package instance

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, instanceID string) (err error) {
	_, err = client.Delete(client.ServiceURL("enterprise-router", "instances", instanceID), &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
