package instances

import "github.com/opentelekomcloud/gophertelekomcloud"

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("instances")
}

func deleteURL(c *golangsdk.ServiceClient, instanceID string) string {
	return c.ServiceURL("instances", instanceID)
}

func actionURL(c *golangsdk.ServiceClient, instanceID string) string {
	return c.ServiceURL("instances", instanceID, "action")
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("instances")
}
