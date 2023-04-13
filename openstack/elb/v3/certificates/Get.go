package certificates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Get retrieves a particular Loadbalancer based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("certificates", id), &r.Body, nil)
	return
}
