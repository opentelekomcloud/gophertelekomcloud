package policies

import "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(client.ServiceURL("scaling_policy", id), nil)
	return
}
