package loadbalancers

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular LoadBalancer based on its
// unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(client.ServiceURL("loadbalancers", id), nil)
	return
}
