package availabilityzones

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List will return the existing availability zones.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, client.ServiceURL("os-availability-zone"), func(r pagination.PageResult) pagination.Page {
		return AvailabilityZonePage{pagination.SinglePageBase(r)}
	})
}
