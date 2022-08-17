package groups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get is a method of getting the detailed information of the group by id
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}
