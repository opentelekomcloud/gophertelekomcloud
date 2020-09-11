package credentials

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath        = "OS-CREDENTIAL"
	credentialsPath = "credentials"
)

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rootPath, credentialsPath)
}

func getURL(client *golangsdk.ServiceClient, credID string) string {
	return client.ServiceURL(rootPath, credentialsPath, credID)
}

func createURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rootPath, credentialsPath)
}

func updateURL(client *golangsdk.ServiceClient, credID string) string {
	return client.ServiceURL(rootPath, credentialsPath, credID)
}

func deleteURL(client *golangsdk.ServiceClient, credID string) string {
	return client.ServiceURL(rootPath, credentialsPath, credID)
}
