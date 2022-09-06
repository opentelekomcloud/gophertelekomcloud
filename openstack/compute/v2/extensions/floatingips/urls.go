package floatingips

import "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "os-floating-ips"

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

func serverURL(client *golangsdk.ServiceClient, serverID string) string {
	return client.ServiceURL("servers/" + serverID + "/action")
}

func associateURL(client *golangsdk.ServiceClient, serverID string) string {
	return serverURL(client, serverID)
}

func disassociateURL(client *golangsdk.ServiceClient, serverID string) string {
	return serverURL(client, serverID)
}
