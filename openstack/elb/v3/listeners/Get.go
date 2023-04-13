package listeners

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get retrieves a particular Listeners based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("listeners", id), &r.Body, nil)
	return
}
