package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// Get retrieves a particular Configuration based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (r GetResult) {
	raw, err := c.Get(c.ServiceURL("configurations", id), nil, openstack.StdRequestOpts())
	return
}
