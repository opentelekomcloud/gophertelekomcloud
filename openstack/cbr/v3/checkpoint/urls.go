package checkpoint

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const baseURL = "checkpoints"

func rootUrl(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(baseURL)
}

func checkpointUrl(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(baseURL, id)
}
