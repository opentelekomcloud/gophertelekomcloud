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

// DisableOpts contains all the values needed to disable protection for a Group.
type DisableOpts struct {
	// Empty
}

// ToGroupDisableMap builds a create request body from DisableOpts.
func (opts DisableOpts) ToGroupDisableMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "stop-server-group")
}

// Disable will disable protection for a protection Group.
func Disable(c *golangsdk.ServiceClient, id string) (r JobResult) {
	opts := DisableOpts{}
	b, err := opts.ToGroupDisableMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200}}
	_, r.Err = c.Post(actionURL(c, id), b, &r.Body, reqOpt)
	return
}
