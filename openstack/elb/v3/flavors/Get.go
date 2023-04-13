package flavors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Get returns additional information about a Flavor, given its ID.
func Get(client *golangsdk.ServiceClient, flavorID string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("flavors", flavorID), &r.Body, nil)
	return
}
