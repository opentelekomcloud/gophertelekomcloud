package configurations

import "github.com/opentelekomcloud/gophertelekomcloud"

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("configurations")
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("configurations", id)
}

func applyURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("configurations", id, "apply")
}

func instanceConfigURL(c *golangsdk.ServiceClient, instanceID string) string {
	return c.ServiceURL("instances", instanceID, "configurations")
}
