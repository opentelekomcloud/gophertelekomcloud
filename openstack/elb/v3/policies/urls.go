package policies

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	basePath = "l7policies"
)

func baseURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(basePath)
}

func resourceURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(basePath, id)
}
