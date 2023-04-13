package members

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToMembersListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections
// through the API. Filtering is achieved by passing in struct field values
// that map to the Member attributes you want to see returned. SortKey allows
// you to sort by a particular Member attribute. SortDir sets the direction,
// and is either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	Name            string `q:"name"`
	Weight          int    `q:"weight"`
	AdminStateUp    *bool  `q:"admin_state_up"`
	SubnetID        string `q:"subnet_sidr_id"`
	Address         string `q:"address"`
	ProtocolPort    int    `q:"protocol_port"`
	ID              string `q:"id"`
	OperatingStatus string `q:"operating_status"`
	Limit           int    `q:"limit"`
	Marker          string `q:"marker"`
	SortKey         string `q:"sort_key"`
	SortDir         string `q:"sort_dir"`
}

// ToMembersListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToMembersListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// members. It accepts a ListOptsBuilder, which allows you to filter and
// sort the returned collection for greater efficiency.
//
// Default policy settings return only those members that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(client *golangsdk.ServiceClient, poolID string, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("pools", poolID, "members")
	if opts != nil {
		query, err := opts.ToMembersListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return MemberPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}
