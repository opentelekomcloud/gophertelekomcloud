package monitors

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get retrieves a particular Health Monitor based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("healthmonitors", id), &r.Body, nil)
	return
}
