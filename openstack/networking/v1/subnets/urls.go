package subnets

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	resourcePath = "subnets"
	rootPath     = "vpcs"
)

func rootURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(client.ProjectID, resourcePath)
}

func resourceURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(client.ProjectID, resourcePath, id)
}

func updateURL(client *golangsdk.ServiceClient, vpcID, id string) string {
	return client.ServiceURL(client.ProjectID, rootPath, vpcID, resourcePath, id)
}
