package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListAddressesByNetwork makes a request against the API to list the servers IP
// addresses for the given network.
func ListAddressesByNetwork(client *golangsdk.ServiceClient, id, network string) pagination.Pager {
	return pagination.NewPager(client, client.ServiceURL("servers", id, "ips", network), func(r pagination.PageResult) pagination.Page {
		return NetworkAddressPage{pagination.SinglePageBase(r)}
	})
}
