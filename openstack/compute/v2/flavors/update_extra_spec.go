package flavors

import "github.com/opentelekomcloud/gophertelekomcloud"

// UpdateExtraSpec will updates the value of the specified flavor's extra spec for the key in opts.
func UpdateExtraSpec(client *golangsdk.ServiceClient, flavorID string, opts UpdateExtraSpecOptsBuilder) (r UpdateExtraSpecResult) {
	b, key, err := opts.ToFlavorExtraSpecUpdateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL("flavors", flavorID, "os-extra_specs", key), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
