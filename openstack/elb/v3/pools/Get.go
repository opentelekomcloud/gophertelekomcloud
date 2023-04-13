package pools

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get retrieves a particular pool based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("pools", id), &r.Body, nil)
	return
}
