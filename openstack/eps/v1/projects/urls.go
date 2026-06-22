package projects

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("enterprise-projects")
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("enterprise-projects", id)
}
