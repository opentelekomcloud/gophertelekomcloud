package availabilityzones

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListDetail will return the existing availability zones with detailed information.
func ListDetail(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, client.ServiceURL("os-availability-zone", "detail"), func(r pagination.PageResult) pagination.Page {
		return AvailabilityZonePage{pagination.SinglePageBase(r)}
	})
}
