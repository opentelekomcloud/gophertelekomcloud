package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// Delete will permanently delete a particular Configuration based on its unique ID.
func Delete(c *golangsdk.ServiceClient, id string) (r DeleteResult) {
	raw, err := c.Delete(c.ServiceURL("configurations", id),
		&golangsdk.RequestOpts{OkCodes: []int{200}, MoreHeaders: openstack.StdRequestOpts().MoreHeaders})
	return
}
