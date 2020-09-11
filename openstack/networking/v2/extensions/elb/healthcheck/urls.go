package healthcheck

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath     = "elbaas"
	resourcePath = "healthcheck"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath, id)
}
