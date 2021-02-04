package eiptags

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "publicips"
const tags = "tags"
const action = "action"

func rootURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(client.ProjectID, resourcePath, id, tags)
}

func deleteURL(client *golangsdk.ServiceClient, id string, key string) string {
	return client.ServiceURL(client.ProjectID, resourcePath, id, tags, key)
}

func actionURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(client.ProjectID, resourcePath, id, tags, action)
}

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(client.ProjectID, resourcePath, tags)
}
