package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("scaling_policy", id), &r.Body, nil)
	return
}
