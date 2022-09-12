package usage

import (
	"net/url"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// SingleTenantOpts are options for fetching usage of a single tenant.
type SingleTenantOpts struct {
	// The ending time to calculate usage statistics on compute and storage resources.
	End *time.Time `q:"end"`
	// The beginning time to calculate usage statistics on compute and storage resources.
	Start *time.Time `q:"start"`
}

// ToUsageSingleTenantQuery formats a SingleTenantOpts into a query string.
func (opts SingleTenantOpts) ToUsageSingleTenantQuery() (string, error) {
	params := make(url.Values)
	if opts.Start != nil {
		params.Add("start", opts.Start.Format(golangsdk.RFC3339MilliNoZ))
	}

	if opts.End != nil {
		params.Add("end", opts.End.Format(golangsdk.RFC3339MilliNoZ))
	}

	q := &url.URL{RawQuery: params.Encode()}
	return q.String(), nil
}

// SingleTenant returns usage data about a single tenant.
func SingleTenant(client *golangsdk.ServiceClient, tenantID string, opts SingleTenantOpts) pagination.Pager {
	query, err := opts.ToUsageSingleTenantQuery()
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("os-simple-tenant-usage", tenantID)+query,
		func(r pagination.PageResult) pagination.Page {
			return SingleTenantPage{pagination.LinkedPageBase{PageResult: r}}
		})
}

// SingleTenantPage stores a single, only page of TenantUsage results from a SingleTenant call.
type SingleTenantPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether a SingleTenantPage is empty.
func (r SingleTenantPage) IsEmpty() (bool, error) {
	ks, err := ExtractSingleTenant(r)
	return ks == nil, err
}

// NextPageURL uses the response's embedded link reference to navigate to the next page of results.
func (r SingleTenantPage) NextPageURL() (string, error) {
	var res []golangsdk.Link
	err := extract.IntoSlicePtr(r.BodyReader(), &res, "tenant_usage_links")
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(res)
}

// ExtractSingleTenant interprets a SingleTenantPage as a TenantUsage result.
func ExtractSingleTenant(page pagination.Page) (*TenantUsage, error) {
	var res TenantUsage
	err := extract.IntoStructPtr(page.(SingleTenantPage).BodyReader(), &res, "tenant_usage")
	return &res, err
}
