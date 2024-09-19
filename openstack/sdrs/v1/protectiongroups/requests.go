package protectiongroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// Get retrieves a particular Group based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, openstack.StdRequestOpts())
	return
}

// Delete will permanently delete a particular Group based on its unique ID.
func Delete(c *golangsdk.ServiceClient, id string) (r JobResult) {
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200}, MoreHeaders: openstack.StdRequestOpts().MoreHeaders}
	_, r.Err = c.DeleteWithResponse(resourceURL(c, id), &r.Body, reqOpt)
	return
}
