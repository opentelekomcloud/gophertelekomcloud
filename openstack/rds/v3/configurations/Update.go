package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToConfigUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains all the values needed to update a Configuration.
type UpdateOpts struct {
	// Configuration Name
	Name string `json:"name,omitempty"`
	// Configuration Description
	Description string `json:"description,omitempty"`
	// Configuration Values
	Values map[string]string `json:"values,omitempty"`
}

// ToConfigUpdateMap builds a update request body from UpdateOpts.
func (opts UpdateOpts) ToConfigUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Update accepts a UpdateOpts struct and uses the values to update a Configuration.The response code from api is 200
func Update(c *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToConfigUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders}
	raw, err := c.Put(c.ServiceURL("configurations", id), b, nil, reqOpt)
	return
}
