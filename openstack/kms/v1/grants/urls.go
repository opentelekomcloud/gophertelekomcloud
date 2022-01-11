package grants

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const (
	resourcePath = "kms"
)

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "create-grant")
}

func deleteURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "revoke-grant")
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "list-grants")
}
