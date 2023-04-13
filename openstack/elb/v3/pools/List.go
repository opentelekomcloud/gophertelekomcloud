package pools

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToPoolListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Pool attributes you want to see returned. SortKey allows you to
// sort by a particular Pool attribute. SortDir sets the direction, and is
// either `asc` or `desc`. Marker and Limit are used for pagination.
type ListOpts struct {
	Description     []string `q:"description"`
	HealthMonitorID []string `q:"healthmonitor_id"`
	LBMethod        []string `q:"lb_algorithm"`
	Protocol        []string `q:"protocol"`
	AdminStateUp    *bool    `q:"admin_state_up"`
	Name            []string `q:"name"`
	ID              []string `q:"id"`
	LoadbalancerID  []string `q:"loadbalancer_id"`
	Limit           int      `q:"limit"`
	Marker          string   `q:"marker"`
	SortKey         string   `q:"sort_key"`
	SortDir         string   `q:"sort_dir"`
}

// ToPoolListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPoolListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List returns a Pager which allows you to iterate over a collection of
// pools. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those pools that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("pools")
	if opts != nil {
		query, err := opts.ToPoolListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return PoolPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}
