package task

import "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "objectstorage/task"

func rootURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(resourcePath)
}

func resourceURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(resourcePath, id)
}
