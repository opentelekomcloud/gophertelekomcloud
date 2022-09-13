package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List makes a request against the API to list servers accessible to you.
func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	query, err := opts.ToServerListQuery()
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("servers", "detail")+query,
		func(r pagination.PageResult) pagination.Page {
			return ServerPage{pagination.LinkedPageBase{PageResult: r}}
		})
}
