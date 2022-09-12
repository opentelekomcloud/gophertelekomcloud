package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListAddresses makes a request against the API to list the servers IP
// addresses.
func ListAddresses(client *golangsdk.ServiceClient, id string) pagination.Pager {
	return pagination.NewPager(client, client.ServiceURL("servers", id, "ips"), func(r pagination.PageResult) pagination.Page {
		return AddressPage{pagination.SinglePageBase(r)}
	})
}
