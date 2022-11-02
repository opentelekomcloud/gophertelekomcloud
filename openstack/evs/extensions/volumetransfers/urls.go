package volumetransfers

import "github.com/opentelekomcloud/gophertelekomcloud"

func transferURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("os-volume-transfer")
}

func acceptURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("os-volume-transfer", id, "accept")
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("os-volume-transfer", id)
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("os-volume-transfer", "detail")
}

func getURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("os-volume-transfer", id)
}
