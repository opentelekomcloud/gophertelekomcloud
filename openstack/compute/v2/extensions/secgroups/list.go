package secgroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List will return a collection of all the security groups for a particular
// tenant.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	return commonList(client, rootURL(client))
}

// ListByServer will return a collection of all the security groups which are
// associated with a particular server.
func ListByServer(client *golangsdk.ServiceClient, serverID string) pagination.Pager {
	return commonList(client, listByServerURL(client, serverID))
}
