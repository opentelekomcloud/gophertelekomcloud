package volumeattach

import "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "os-volume_attachments"

func resourceURL(client *golangsdk.ServiceClient, serverID string) string {
	return client.ServiceURL("servers", serverID, resourcePath)
}

func listURL(client *golangsdk.ServiceClient, serverID string) string {
	return resourceURL(client, serverID)
}

func createURL(client *golangsdk.ServiceClient, serverID string) string {
	return resourceURL(client, serverID)
}

func getURL(client *golangsdk.ServiceClient, serverID, aID string) string {
	return client.ServiceURL("servers", serverID, resourcePath, aID)
}

func deleteURL(client *golangsdk.ServiceClient, serverID, aID string) string {
	return getURL(client, serverID, aID)
}
