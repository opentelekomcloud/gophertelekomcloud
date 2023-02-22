package images

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get implements image get request.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("images", id), &r.Body, nil)
	return
}
