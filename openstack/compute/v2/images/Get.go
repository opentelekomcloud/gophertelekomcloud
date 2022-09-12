package images

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get returns data about a specific image by its ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("images", id), nil, nil)
	return
}
