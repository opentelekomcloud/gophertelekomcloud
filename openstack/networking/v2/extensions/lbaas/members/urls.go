package members

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath     = "lb"
	resourcePath = "members"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
