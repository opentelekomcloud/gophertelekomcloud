package backup

import "github.com/opentelekomcloud/gophertelekomcloud"

const rootPath = "providers"
const providerID = "fc4d5750-22e7-4798-8a46-f48f62c4c1da"

func rootURL(c *golangsdk.ServiceClient, resourceID string) string {
	return c.ServiceURL(rootPath, providerID, "resources", resourceID, "action")
}

func resourceURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, providerID, "resources", "action")
}

func getURL(c *golangsdk.ServiceClient, checkpointItemID string) string {
	return c.ServiceURL("checkpoint_items", checkpointItemID)
}

func deleteURL(c *golangsdk.ServiceClient, checkpointID string) string {
	return c.ServiceURL(rootPath, providerID, "checkpoints", checkpointID)
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("checkpoint_items")
}
