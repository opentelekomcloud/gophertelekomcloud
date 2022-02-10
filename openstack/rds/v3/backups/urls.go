package backups

import "github.com/opentelekomcloud/gophertelekomcloud"

func baseURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("backups")
}

func backupURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("backups", id)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("instances", id, "backups/policy")
}

func restoreURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("instances", "recovery")
}
