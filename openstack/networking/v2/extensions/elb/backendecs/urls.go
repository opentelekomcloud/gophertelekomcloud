package backendecs

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath     = "elbaas"
	resourcePath = "listeners"
	memberPath   = "members"
)

func rootURL(c *golangsdk.ServiceClient, lId string) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath, lId, memberPath)
}

func actionURL(c *golangsdk.ServiceClient, lId string) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath, lId, memberPath, "action")
}
