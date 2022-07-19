package agency

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	rootPath     = "OS-AGENCY"
	resourcePath = "agencies"
)

func rootURL(c *golangsdk.ServiceClient) string {
	url := c.ServiceURL(rootPath, resourcePath)
	return strings.Replace(url, "/v3/", "/v3.0/", 1)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	url := c.ServiceURL(rootPath, resourcePath, id)
	return strings.Replace(url, "/v3/", "/v3.0/", 1)
}

func roleURL(c *golangsdk.ServiceClient, resource, resourceID, agencyID, roleID string) string {
	url := c.ServiceURL(rootPath, resource, resourceID, resourcePath, agencyID, "roles", roleID)
	return strings.Replace(url, "/v3/", "/v3.0/", 1)
}

func listRolesURL(c *golangsdk.ServiceClient, resource, resourceID, agencyID string) string {
	url := c.ServiceURL(rootPath, resource, resourceID, resourcePath, agencyID, "roles")
	return strings.Replace(url, "/v3/", "/v3.0/", 1)
}
