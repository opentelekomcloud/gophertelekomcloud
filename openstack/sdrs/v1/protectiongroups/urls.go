package protectiongroups

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("server-groups", id)
}
