package keypairs

import "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "os-keypairs"

func resourceURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(resourcePath)
}

func listURL(client *golangsdk.ServiceClient) string {
	return resourceURL(client)
}

func createURL(client *golangsdk.ServiceClient) string {
	return resourceURL(client)
}

func getURL(client *golangsdk.ServiceClient, name string) string {
	return client.ServiceURL(resourcePath, name)
}

func deleteURL(client *golangsdk.ServiceClient, name string) string {
	return getURL(client, name)
}
