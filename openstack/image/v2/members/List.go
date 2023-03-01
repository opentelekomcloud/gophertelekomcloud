package members

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List members returns list of members for specified image id.
func List(client *golangsdk.ServiceClient, imageID string) pagination.Pager {
	return pagination.NewPager(client, client.ServiceURL("images", imageID, "members"), func(r pagination.PageResult) pagination.Page {
		return MemberPage{pagination.SinglePageBase(r)}
	})
}
