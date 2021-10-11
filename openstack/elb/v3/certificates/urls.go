package certificates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	v3 "github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3"
)

const (
	resourcePath = "certificates"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(v3.RootPath, resourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(v3.RootPath, resourcePath, id)
}
