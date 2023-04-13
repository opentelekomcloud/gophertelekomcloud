package monitors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToMonitorListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Monitor attributes you want to see returned. SortKey allows you to
// sort by a particular Monitor attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	ID            string `q:"id"`
	Name          string `q:"name"`
	TenantID      string `q:"tenant_id"`
	ProjectID     string `q:"project_id"`
	PoolID        string `q:"pool_id"`
	Type          string `q:"type"`
	Delay         int    `q:"delay"`
	Timeout       int    `q:"timeout"`
	MaxRetries    int    `q:"max_retries"`
	HTTPMethod    string `q:"http_method"`
	URLPath       string `q:"url_path"`
	ExpectedCodes string `q:"expected_codes"`
	AdminStateUp  *bool  `q:"admin_state_up"`
	Status        string `q:"status"`
	Limit         int    `q:"limit"`
	Marker        string `q:"marker"`
	SortKey       string `q:"sort_key"`
	SortDir       string `q:"sort_dir"`
}

// ToMonitorListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToMonitorListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List returns a Pager which allows you to iterate over a collection of
// health monitors. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those health monitors that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("healthmonitors")
	if opts != nil {
		query, err := opts.ToMonitorListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return MonitorPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}
