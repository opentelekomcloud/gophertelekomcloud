package flavors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListAccesses retrieves the tenants which have access to a flavor.
func ListAccesses(client *golangsdk.ServiceClient, id string) pagination.Pager {
	url := client.ServiceURL("flavors", id, "os-flavor-access")

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return AccessPage{pagination.SinglePageBase(r)}
	})
}
