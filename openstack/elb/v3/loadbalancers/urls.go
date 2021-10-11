package loadbalancers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3"
)

const (
	resourcePath = "loadbalancers"
	statusPath   = "statuses"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(v3.RootPath, resourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(v3.RootPath, resourcePath, id)
}

func statusURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(v3.RootPath, resourcePath, id, statusPath)
}
