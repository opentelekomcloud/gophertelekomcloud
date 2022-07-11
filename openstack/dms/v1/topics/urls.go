package topics

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "instances"

func createURL(client *golangsdk.ServiceClient, instanceId string) string {
	return client.ServiceURL(resourcePath, instanceId, "topics")
}

func getURL(client *golangsdk.ServiceClient, instanceId string) string {
	return client.ServiceURL(resourcePath, instanceId, "topics")
}

func deleteURL(client *golangsdk.ServiceClient, instanceId string) string {
	return client.ServiceURL(resourcePath, instanceId, "topics", "delete")
}
