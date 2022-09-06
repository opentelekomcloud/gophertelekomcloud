package servergroups

import "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "os-server-groups"

func resourceURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(resourcePath)
}

func listURL(client *golangsdk.ServiceClient) string {
	return resourceURL(client)
}

func createURL(client *golangsdk.ServiceClient) string {
	return resourceURL(client)
}

func getURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(resourcePath, id)
}

func deleteURL(client *golangsdk.ServiceClient, id string) string {
	return getURL(client, id)
}
