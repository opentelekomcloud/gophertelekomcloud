package endpoints

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const basePath = "vpc-endpoints"

func baseURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(basePath)
}

func resourceURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(basePath, id)
}
