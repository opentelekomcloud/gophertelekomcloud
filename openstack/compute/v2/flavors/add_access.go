package flavors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// AccessOpts represents options for adding access to a flavor.
type AccessOpts struct {
	// Tenant is the project/tenant ID to grant access.
	Tenant string `json:"tenant"`
}

// AddAccess grants a tenant/project access to a flavor.
func AddAccess(client *golangsdk.ServiceClient, id string, opts AccessOpts) ([]FlavorAccess, error) {
	b, err := build.RequestBody(opts, "addTenantAccess")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("flavors", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extraAcc(err, raw)
}

// FlavorAccess represents an ACL of tenant access to a specific Flavor.
type FlavorAccess struct {
	// FlavorID is the unique ID of the flavor.
	FlavorID string `json:"flavor_id"`
	// TenantID is the unique ID of the tenant.
	TenantID string `json:"tenant_id"`
}
