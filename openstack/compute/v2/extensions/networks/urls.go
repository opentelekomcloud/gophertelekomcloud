package networks

import "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "os-networks"

func resourceURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(resourcePath)
}

func listURL(client *golangsdk.ServiceClient) string {
	return resourceURL(client)
}

func getURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(resourcePath, id)
}
