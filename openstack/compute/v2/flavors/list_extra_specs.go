package flavors

import "github.com/opentelekomcloud/gophertelekomcloud"

// ListExtraSpecs requests all the extra-specs for the given flavor ID.
func ListExtraSpecs(client *golangsdk.ServiceClient, flavorID string) (r ListExtraSpecsResult) {
	raw, err := client.Get(client.ServiceURL("flavors", flavorID, "os-extra_specs"), nil, nil)
	return
}
