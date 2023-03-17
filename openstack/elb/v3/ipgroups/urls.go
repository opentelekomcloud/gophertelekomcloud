package ipgroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	resourcePath                   = "ipgroups"
	ipListResourcePath             = "iplist"
	ipListResourceActionUpdatePath = "create-or-update"
	ipListResourceActionDeletePath = "batch-delete"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func ipListUpdateResourceURL(c *golangsdk.ServiceClient, ipGroupId string) string {
	return c.ServiceURL(ipGroupId, ipListResourcePath, ipListResourceActionUpdatePath)
}

func ipListDeleteResourceURL(c *golangsdk.ServiceClient, ipGroupId string) string {
	return c.ServiceURL(ipGroupId, ipListResourcePath, ipListResourceActionDeletePath)
}
