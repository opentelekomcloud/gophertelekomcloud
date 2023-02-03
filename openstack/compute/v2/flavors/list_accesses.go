package flavors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListAccesses retrieves the tenants which have access to a flavor.
func ListAccesses(client *golangsdk.ServiceClient, id string) ([]FlavorAccess, error) {
	raw, err := client.Get(client.ServiceURL("flavors", id, "os-flavor-access"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []FlavorAccess
	err = extract.IntoSlicePtr(raw.Body, &res, "flavor_access")
	return res, err
}
