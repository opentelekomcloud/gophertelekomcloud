package backups

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const rootURL = "backups"

func singleURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rootURL, id)
}

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rootURL)
}

func restoreUrl(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rootURL, id, "restore")
}
