package flavors

import "github.com/opentelekomcloud/gophertelekomcloud"

// ListExtraSpecs requests all the extraAcc-specs for the given flavor ID.
func ListExtraSpecs(client *golangsdk.ServiceClient, flavorID string) (map[string]string, error) {
	raw, err := client.Get(client.ServiceURL("flavors", flavorID, "os-extra_specs"), nil, nil)
	return extraSpes(err, raw)
}
