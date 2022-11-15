package users

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	rootPath          = "users"
	openstackUserPath = "OS-USER"
	groupsPath        = "groups"
	projectPath       = "projects"
	welcomePath       = "welcome"
)

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rootPath)
}

func getURL(client *golangsdk.ServiceClient, userID string) string {
	return client.ServiceURL(rootPath, userID)
}

func createURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rootPath)
}

func updateURL(client *golangsdk.ServiceClient, userID string) string {
	return client.ServiceURL(rootPath, userID)
}

func updateExtendedURL(client *golangsdk.ServiceClient, userID string) string {
	url := client.ServiceURL(openstackUserPath, rootPath, userID)
	return v30(url)
}

func welcomeExtendedURL(client *golangsdk.ServiceClient, userID string) string {
	url := client.ServiceURL(openstackUserPath, rootPath, userID, welcomePath)
	return v30(url)
}

func deleteURL(client *golangsdk.ServiceClient, userID string) string {
	return client.ServiceURL(rootPath, userID)
}

func listGroupsURL(client *golangsdk.ServiceClient, userID string) string {
	return client.ServiceURL(rootPath, userID, groupsPath)
}

func listProjectsURL(client *golangsdk.ServiceClient, userID string) string {
	return client.ServiceURL(rootPath, userID, projectPath)
}

func listInGroupURL(client *golangsdk.ServiceClient, groupID string) string {
	return client.ServiceURL(groupsPath, groupID, rootPath)
}

func membershipURL(client *golangsdk.ServiceClient, groupID string, userID string) string {
	return client.ServiceURL(groupsPath, groupID, rootPath, userID)
}

func v30(url string) string {
	return strings.Replace(url, "/v3/", "/v3.0/", 1)
}
