package aggregates

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List makes a request against the API to list aggregates.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, aggregatesListURL(client), func(r pagination.PageResult) pagination.Page {
		return AggregatesPage{pagination.SinglePageBase(r)}
	})
}
