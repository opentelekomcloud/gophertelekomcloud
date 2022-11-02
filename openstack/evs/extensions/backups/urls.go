package backups

import "github.com/opentelekomcloud/gophertelekomcloud"

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("backups")
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("backups", id)
}

func getURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("backups", id)
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("backups")
}

func listDetailURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("backups", "detail")
}

func updateURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("backups", id)
}

func restoreURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("backups", id, "restore")
}

func exportURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("backups", id, "export_record")
}

func importURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("backups", "import_record")
}
