package defsecrules

import "github.com/opentelekomcloud/gophertelekomcloud"

const rulepath = "os-security-group-default-rules"

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(rulepath, id)
}

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rulepath)
}
