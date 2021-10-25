package members

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const (
	resourcePath = "pools"
	memberPath   = "members"
)

func rootURL(c *golangsdk.ServiceClient, poolID string) string {
	return c.ServiceURL(resourcePath, poolID, memberPath)
}

func resourceURL(c *golangsdk.ServiceClient, poolID string, memberID string) string {
	return c.ServiceURL(resourcePath, poolID, memberPath, memberID)
}
