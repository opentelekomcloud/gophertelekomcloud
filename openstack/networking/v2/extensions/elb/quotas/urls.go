package quotas

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath     = "elbaas"
	resourcePath = "quotas"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath)
}
