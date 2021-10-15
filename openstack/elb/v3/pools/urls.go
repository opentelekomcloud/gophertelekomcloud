package pools

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const (
	resourcePath = "pools"
)

func rootURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(resourcePath)
}

func resourceURL(client *golangsdk.ServiceClient, poolID string) string {
	return client.ServiceURL(resourcePath, poolID)
}
