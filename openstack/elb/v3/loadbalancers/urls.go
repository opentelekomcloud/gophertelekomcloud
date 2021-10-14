package loadbalancers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	resourcePath = "loadbalancers"
	statusPath   = "statuses"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func statusURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, statusPath)
}
