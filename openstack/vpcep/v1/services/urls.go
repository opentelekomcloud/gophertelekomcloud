package services

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

const basePath = "vpc-endpoint-services"

func baseURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(basePath)
}

func publicURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(basePath, "public")
}

func resourceURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(basePath, id)
}
