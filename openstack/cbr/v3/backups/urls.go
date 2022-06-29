package backups

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const rootPath = "backups"

func singleURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rootPath, id)
}

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rootPath)
}

func restoreURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rootPath, id, "restore")
}
