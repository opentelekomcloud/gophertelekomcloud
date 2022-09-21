package lifecycle

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get a instance with detailed information by id
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("instances", id), &r.Body, nil)
	return
}
