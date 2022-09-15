package clusters

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func List(client *golangsdk.ServiceClient) (p pagination.Pager) {
	return pagination.NewPager(client, client.ServiceURL("clusters"), func(r pagination.PageResult) pagination.Page {
		return ClusterPage{pagination.SinglePageBase(r)}
	})
}
