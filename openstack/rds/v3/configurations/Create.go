package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToConfigCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains all the values needed to create a new configuration.
type CreateOpts struct {
	// Configuration Name
	Name string `json:"name" required:"true"`
	// Configuration Description
	Description string `json:"description,omitempty"`
	// Configuration Values
	Values map[string]string `json:"values,omitempty"`
	// Database Object
	DataStore DataStore `json:"datastore" required:"true"`
}

type DataStore struct {
	// DB Engine
	Type string `json:"type" required:"true"`
	// DB version
	Version string `json:"version" required:"true"`
}

// ToConfigCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToConfigCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create will create a new Config based on the values in CreateOpts.
func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToConfigCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(c.ServiceURL("configurations"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders, JSONBody: nil,
	})
	return
}
