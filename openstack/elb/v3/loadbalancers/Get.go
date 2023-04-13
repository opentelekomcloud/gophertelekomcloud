package loadbalancers

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get retrieves a particular Loadbalancer based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("loadbalancers", id), &r.Body, nil)
	return
}
