package checkpoint

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const basePath = "checkpoints"

func rootURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(basePath)
}

func checkpointURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(basePath, id)
}
