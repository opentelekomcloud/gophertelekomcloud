package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	rootPath            = "roles"
	openstackPolicyPath = "OS-ROLE"
)

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(openstackPolicyPath, rootPath)
}

func getURL(client *golangsdk.ServiceClient, roleId string) string {
	return client.ServiceURL(openstackPolicyPath, rootPath, roleId)
}

func createURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(openstackPolicyPath, rootPath)
}

func updateURL(client *golangsdk.ServiceClient, roleId string) string {
	return client.ServiceURL(openstackPolicyPath, rootPath, roleId)
}

func deleteURL(client *golangsdk.ServiceClient, roleId string) string {
	return client.ServiceURL(openstackPolicyPath, rootPath, roleId)
}
