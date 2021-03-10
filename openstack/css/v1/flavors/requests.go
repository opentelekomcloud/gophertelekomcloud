package flavors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func List(client *golangsdk.ServiceClient) (p pagination.Pager) {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return VersionPage{pagination.SinglePageBase(r)}
	})
}
