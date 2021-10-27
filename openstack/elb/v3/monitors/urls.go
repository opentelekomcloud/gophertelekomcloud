package monitors

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const (
	resourcePath = "healthmonitors"
)

func rootURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(resourcePath)
}

func resourceURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(resourcePath, id)
}
