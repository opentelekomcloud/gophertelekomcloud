package usage

import (
	"net/url"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// SingleTenantOpts are options for fetching usage of a single tenant.
type SingleTenantOpts struct {
	// The ending time to calculate usage statistics on compute and storage resources.
	End *time.Time `q:"end"`

	// The beginning time to calculate usage statistics on compute and storage resources.
	Start *time.Time `q:"start"`
}

// SingleTenantOptsBuilder allows extensions to add additional parameters to the
// SingleTenant request.
type SingleTenantOptsBuilder interface {
	ToUsageSingleTenantQuery() (string, error)
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
func SingleTenant(client *golangsdk.ServiceClient, tenantID string, opts SingleTenantOptsBuilder) pagination.Pager {
	u := getTenantURL(client, tenantID)
	if opts != nil {
		query, err := opts.ToUsageSingleTenantQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		u += query
	}
	return pagination.Pager{
		Client:     client,
		InitialURL: u,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return SingleTenantPage{pagination.LinkedPageBase{PageResult: r}}
		},
	}
}

// AllTenantsOpts are options for fetching usage of all tenants.
type AllTenantsOpts struct {
	// Detailed will return detailed results.
	Detailed bool

	// The ending time to calculate usage statistics on compute and storage resources.
	End *time.Time `q:"end"`

	// The beginning time to calculate usage statistics on compute and storage resources.
	Start *time.Time `q:"start"`
}

// AllTenantsOptsBuilder allows extensions to add additional parameters to the
// AllTenants request.
type AllTenantsOptsBuilder interface {
	ToUsageAllTenantsQuery() (string, error)
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
func AllTenants(client *golangsdk.ServiceClient, opts AllTenantsOptsBuilder) pagination.Pager {
	u := allTenantsURL(client)
	if opts != nil {
		query, err := opts.ToUsageAllTenantsQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		u += query
	}
	return pagination.Pager{
		Client:     client,
		InitialURL: u,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return AllTenantsPage{pagination.LinkedPageBase{PageResult: r}}
		},
	}
}
