package job

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

func Delete(c *golangsdk.ServiceClient, id string) (err error) {
	_, err = c.Delete(c.ServiceURL("job-executions", id), &golangsdk.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders,
	})
	return
}
