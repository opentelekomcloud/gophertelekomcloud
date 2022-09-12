package flavors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListDetail instructs OpenStack to provide a list of flavors.
// You may provide criteria by which List curtails its results for easier processing.
func ListDetail(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("flavors", "detail")
	if opts != nil {
		query, err := opts.ToFlavorListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return FlavorPage{pagination.LinkedPageBase{PageResult: r}}
	})
}
