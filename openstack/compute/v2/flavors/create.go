package flavors

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create requests the creation of a new flavor.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToFlavorCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("flavors"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}
