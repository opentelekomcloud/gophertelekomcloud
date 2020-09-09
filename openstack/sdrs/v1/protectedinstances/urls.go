package protectedinstances

import "github.com/opentelekomcloud/gophertelekomcloud"

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("protected-instances")
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("protected-instances", id)
}
