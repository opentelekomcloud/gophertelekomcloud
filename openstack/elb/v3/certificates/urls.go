package certificates

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath     = "elb"
	resourcePath = "certificates"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
