package loadbalancers

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath     = "elbaas"
	resourcePath = "loadbalancers"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath, id)
}
