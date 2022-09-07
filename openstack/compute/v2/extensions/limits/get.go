package limits

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// GetOptsBuilder allows extensions to add additional parameters to the
// Get request.
type GetOptsBuilder interface {
	ToLimitsQuery() (string, error)
}

// GetOpts enables retrieving limits by a specific tenant.
type GetOpts struct {
	// The tenant ID to retrieve limits for.
	TenantID string `q:"tenant_id"`
}

// ToLimitsQuery formats a GetOpts into a query string.
func (opts GetOpts) ToLimitsQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

// Get returns the limits about the currently scoped tenant.
func Get(client *golangsdk.ServiceClient, opts GetOptsBuilder) (r GetResult) {
	url := client.ServiceURL("limits")
	if opts != nil {
		query, err := opts.ToLimitsQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}

	raw, err := client.Get(url, nil, nil)
	return
}
