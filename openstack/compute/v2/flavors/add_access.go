package flavors

import "github.com/opentelekomcloud/gophertelekomcloud"

// AddAccess grants a tenant/project access to a flavor.
func AddAccess(client *golangsdk.ServiceClient, id string, opts AddAccessOptsBuilder) (r AddAccessResult) {
	b, err := opts.ToFlavorAddAccessMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("flavors", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
