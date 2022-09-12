package usage

import (
	"net/url"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// AllTenantsOpts are options for fetching usage of all tenants.
type AllTenantsOpts struct {
	// Detailed will return detailed results.
	Detailed bool
	// The ending time to calculate usage statistics on compute and storage resources.
	End *time.Time `q:"end"`
	// The beginning time to calculate usage statistics on compute and storage resources.
	Start *time.Time `q:"start"`
}

// ToUsageAllTenantsQuery formats a AllTenantsOpts into a query string.
func (opts AllTenantsOpts) ToUsageAllTenantsQuery() (string, error) {
	params := make(url.Values)
	if opts.Start != nil {
		params.Add("start", opts.Start.Format(golangsdk.RFC3339MilliNoZ))
	}

	if opts.End != nil {
		params.Add("end", opts.End.Format(golangsdk.RFC3339MilliNoZ))
	}

	if opts.Detailed {
		params.Add("detailed", "1")
	}

	q := &url.URL{RawQuery: params.Encode()}
	return q.String(), nil
}

// AllTenants returns usage data about all tenants.
func AllTenants(client *golangsdk.ServiceClient, opts AllTenantsOpts) pagination.Pager {
	query, err := opts.ToUsageAllTenantsQuery()
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("os-simple-tenant-usage")+query,
		func(r pagination.PageResult) pagination.Page {
			return AllTenantsPage{pagination.LinkedPageBase{PageResult: r}}
		})
}

// AllTenantsPage stores a single, only page of TenantUsage results from a AllTenants call.
type AllTenantsPage struct {
	pagination.LinkedPageBase
}

// ExtractAllTenants interprets a AllTenantsPage as a TenantUsage result.
func ExtractAllTenants(page pagination.Page) ([]TenantUsage, error) {
	var res []TenantUsage
	err := extract.IntoSlicePtr(page.(AllTenantsPage).BodyReader(), &res, "tenant_usages")
	return res, err
}

// IsEmpty determines whether an AllTenantsPage is empty.
func (r AllTenantsPage) IsEmpty() (bool, error) {
	usages, err := ExtractAllTenants(r)
	return len(usages) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to the next page of results.
func (r AllTenantsPage) NextPageURL() (string, error) {
	var res []golangsdk.Link
	err := extract.IntoSlicePtr(r.BodyReader(), &res, "tenant_usages_links")
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(res)
}
