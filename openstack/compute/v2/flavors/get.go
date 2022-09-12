package flavors

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get retrieves details of a single flavor. Use ExtractFlavor to convert its result into a Flavor.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("flavors", id), nil, nil)
	return
}
