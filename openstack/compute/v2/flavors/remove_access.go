package flavors

import "github.com/opentelekomcloud/gophertelekomcloud"

// RemoveAccess removes/revokes a tenant/project access to a flavor.
func RemoveAccess(client *golangsdk.ServiceClient, id string, opts RemoveAccessOptsBuilder) (r RemoveAccessResult) {
	b, err := opts.ToFlavorRemoveAccessMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("flavors", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
