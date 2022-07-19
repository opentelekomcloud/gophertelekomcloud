package policies

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	rootPath            = "roles"
	openstackPolicyPath = "OS-ROLE"
)

func listURL(client *golangsdk.ServiceClient) string {
	url := client.ServiceURL(openstackPolicyPath, rootPath)
	return strings.Replace(url, "/v3/", "/v3.0/", 1)
}

func getURL(client *golangsdk.ServiceClient, roleId string) string {
	url := client.ServiceURL(openstackPolicyPath, rootPath, roleId)
	return strings.Replace(url, "/v3/", "/v3.0/", 1)
}

func createURL(client *golangsdk.ServiceClient) string {
	url := client.ServiceURL(openstackPolicyPath, rootPath)
	return strings.Replace(url, "/v3/", "/v3.0/", 1)
}

func updateURL(client *golangsdk.ServiceClient, roleId string) string {
	url := client.ServiceURL(openstackPolicyPath, rootPath, roleId)
	return strings.Replace(url, "/v3/", "/v3.0/", 1)
}

func deleteURL(client *golangsdk.ServiceClient, roleId string) string {
	url := client.ServiceURL(openstackPolicyPath, rootPath, roleId)
	return strings.Replace(url, "/v3/", "/v3.0/", 1)
}
