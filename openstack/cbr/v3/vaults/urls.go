package vaults

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const baseURL = "vaults"

func rootURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(baseURL)
}

func vaultURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(baseURL, id)
}
