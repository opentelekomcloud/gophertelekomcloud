package tasks

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const rootURL = "operation-logs"

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rootURL)
}

func singleURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rootURL, id)
}
