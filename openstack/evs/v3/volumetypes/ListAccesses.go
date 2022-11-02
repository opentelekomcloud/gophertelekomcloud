package volumetypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListAccesses retrieves the tenants which have access to a volume type.
func ListAccesses(client *golangsdk.ServiceClient, id string) pagination.Pager {
	url := client.ServiceURL("types", id, "os-volume-type-access")

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return AccessPage{pagination.SinglePageBase(r)}
	})
}
