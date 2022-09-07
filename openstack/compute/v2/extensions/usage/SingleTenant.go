package usage

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// SingleTenant returns usage data about a single tenant.
func SingleTenant(client *golangsdk.ServiceClient, tenantID string, opts SingleTenantOptsBuilder) pagination.Pager {
	u := getTenantURL(client, tenantID)
	if opts != nil {
		query, err := opts.ToUsageSingleTenantQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		u += query
	}
	return pagination.NewPager(client, u, func(r pagination.PageResult) pagination.Page {
		return SingleTenantPage{pagination.LinkedPageBase{PageResult: r}}
	})
}
