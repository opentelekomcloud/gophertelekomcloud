package quotasets

import "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "os-quota-sets"

func getURL(c *golangsdk.ServiceClient, projectID string) string {
	return c.ServiceURL(resourcePath, projectID)
}

func getDefaultsURL(c *golangsdk.ServiceClient, projectID string) string {
	return c.ServiceURL(resourcePath, projectID, "defaults")
}

func updateURL(c *golangsdk.ServiceClient, projectID string) string {
	return getURL(c, projectID)
}
