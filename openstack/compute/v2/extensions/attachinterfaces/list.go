package attachinterfaces

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List makes a request against the nova API to list the server's interfaces.
func List(client *golangsdk.ServiceClient, serverID string) pagination.Pager {
	return pagination.NewPager(client, listInterfaceURL(client, serverID), func(r pagination.PageResult) pagination.Page {
		return InterfacePage{pagination.SinglePageBase(r)}
	})
}
