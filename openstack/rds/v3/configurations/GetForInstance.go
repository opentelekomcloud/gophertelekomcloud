package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// GetForInstance retrieves Configuration applied to particular RDS instance
// configuration ID and Name will be empty
func GetForInstance(c *golangsdk.ServiceClient, instanceID string) (r GetResult) {
	_, r.Err = c.Get(c.ServiceURL("instances", instanceID, "configurations"), &r.Body, openstack.StdRequestOpts())
	return
}
