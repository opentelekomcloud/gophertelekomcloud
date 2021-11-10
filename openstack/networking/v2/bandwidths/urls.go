package bandwidths

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	resourcePath = "bandwidths"
	insertPath   = "insert"
	removePath   = "remove"
)

func rootURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(client.ProjectID, resourcePath)
}

func resourceURL(client *golangsdk.ServiceClient, ID string) string {
	return client.ServiceURL(client.ProjectID, resourcePath, ID)
}

func insertURL(client *golangsdk.ServiceClient, ID string) string {
	return client.ServiceURL(client.ProjectID, resourcePath, ID, insertPath)
}

func removeURL(client *golangsdk.ServiceClient, ID string) string {
	return client.ServiceURL(client.ProjectID, resourcePath, ID, removePath)
}
