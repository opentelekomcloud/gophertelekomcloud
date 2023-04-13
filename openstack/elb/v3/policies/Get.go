package policies

import "github.com/opentelekomcloud/gophertelekomcloud"

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("l7policies", id), &r.Body, nil)
	return
}
