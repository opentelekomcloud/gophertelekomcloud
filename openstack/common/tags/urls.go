package tags

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	resourcePath = "tags"
	actionPath   = "action"
)

// supported resourceType: "vpcs", "subnets", "publicips"
// "DNS-public_zone", "DNS-private_zone", "DNS-ptr_record"
// "DNS-public_recordset", "DNS-private_recordset"
func actionURL(c *golangsdk.ServiceClient, resourceType, id string) string {
	if hasProjectID(c) {
		return c.ServiceURL(resourceType, id, resourcePath, actionPath)
	}
	return c.ServiceURL(c.ProjectID, resourceType, id, resourcePath, actionPath)
}

func getURL(c *golangsdk.ServiceClient, resourceType, id string) string {
	if hasProjectID(c) {
		return c.ServiceURL(resourceType, id, resourcePath)
	}
	return c.ServiceURL(c.ProjectID, resourceType, id, resourcePath)
}

func listURL(c *golangsdk.ServiceClient, resourceType string) string {
	if hasProjectID(c) {
		return c.ServiceURL(resourceType, resourcePath)
	}
	return c.ServiceURL(c.ProjectID, resourceType, resourcePath)
}

func hasProjectID(c *golangsdk.ServiceClient) bool {
	url := c.ResourceBaseURL()
	array := strings.Split(url, "/")

	// the baseURL must be end with "/"
	return array[len(array)-2] == c.ProjectID
}
