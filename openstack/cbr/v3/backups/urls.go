package backups

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const rootURL = "backups"

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rootURL)
}

func singleURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rootURL, id)
}

func restoreURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rootURL, id, "restore")
}
