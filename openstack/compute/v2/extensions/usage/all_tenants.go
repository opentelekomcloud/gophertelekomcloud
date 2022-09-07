package usage

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// AllTenants returns usage data about all tenants.
func AllTenants(client *golangsdk.ServiceClient, opts AllTenantsOptsBuilder) pagination.Pager {
	u := client.ServiceURL("os-simple-tenant-usage")
	if opts != nil {
		query, err := opts.ToUsageAllTenantsQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		u += query
	}
	return pagination.NewPager(client, u, func(r pagination.PageResult) pagination.Page {
		return AllTenantsPage{pagination.LinkedPageBase{PageResult: r}}
	})
}
