package attachments

import "github.com/opentelekomcloud/gophertelekomcloud"

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("attachments")
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("attachments", "detail")
}

func getURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("attachments", id)
}

func updateURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("attachments", id)
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("attachments", id)
}

func completeURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("attachments", id, "action")
}
