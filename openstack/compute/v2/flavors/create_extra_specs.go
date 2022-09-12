package flavors

import "github.com/opentelekomcloud/gophertelekomcloud"

// CreateExtraSpecs will create or update the extra-specs key-value pairs for the specified Flavor.
func CreateExtraSpecs(client *golangsdk.ServiceClient, flavorID string, opts CreateExtraSpecsOptsBuilder) (r CreateExtraSpecsResult) {
	b, err := opts.ToFlavorExtraSpecsCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("flavors", flavorID, "os-extra_specs"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
