package shares

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath     = "sfs-turbo"
	resourcePath = "shares"
)

// createURL used to assemble the URI of creating API
func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

// resourceURL used to assemble the URI of deleting API or querying details API
func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

// listURL used to assemble the URI of querying all file system details API
func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath, "detail")
}

// For manage the specified file system, e.g.: extend
func actionURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id, "action")
}
