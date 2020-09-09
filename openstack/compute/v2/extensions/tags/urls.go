package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func createURL(c *golangsdk.ServiceClient, server_id string) string {
	return c.ServiceURL("servers", server_id, "tags")
}

func getURL(c *golangsdk.ServiceClient, server_id string) string {
	return c.ServiceURL("servers", server_id, "tags")
}

func deleteURL(c *golangsdk.ServiceClient, server_id string) string {
	return c.ServiceURL("servers", server_id, "tags")
}
