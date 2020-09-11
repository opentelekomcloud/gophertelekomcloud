package whitelists

import "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "instance"

// resourceURL will build the url of put and get request url
// url: client.Endpoint/instance/{instance_id}/whitelist
func resourceURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(resourcePath, id, "whitelist")
}
