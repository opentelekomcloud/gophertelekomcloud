package instances

import "github.com/opentelekomcloud/gophertelekomcloud"

// endpoint/instances
const resourcePath = "instances"

// updateURL will build the url of update
func updateURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func changePasswordURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(resourcePath, id, "password")
}

func crossVpcURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(resourcePath, id, "crossvpc/modify")
}
