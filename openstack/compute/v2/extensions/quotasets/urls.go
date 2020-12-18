package quotasets

import "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "os-quota-sets"

func getURL(c *golangsdk.ServiceClient, tenantID string) string {
	return c.ServiceURL(resourcePath, tenantID)
}

func getDetailURL(c *golangsdk.ServiceClient, tenantID string) string {
	return c.ServiceURL(resourcePath, tenantID, "detail")
}

func updateURL(c *golangsdk.ServiceClient, tenantID string) string {
	return getURL(c, tenantID)
}

func deleteURL(c *golangsdk.ServiceClient, tenantID string) string {
	return getURL(c, tenantID)
}
