package lifecycle

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func List(client *golangsdk.ServiceClient, opts ListDcsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToDcsListDetailQuery()

		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pageDcsList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return DcsPage{pagination.SinglePageBase(r)}
	})

	dcsheader := map[string]string{"Content-Type": "application/json"}
	pageDcsList.Headers = dcsheader
	return pageDcsList
}
