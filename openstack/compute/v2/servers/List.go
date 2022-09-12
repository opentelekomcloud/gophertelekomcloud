package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List makes a request against the API to list servers accessible to you.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("servers", "detail")
	if opts != nil {
		query, err := opts.ToServerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ServerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}
