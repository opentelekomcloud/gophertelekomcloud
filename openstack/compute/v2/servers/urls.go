package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath     = "servers"
	nicPath      = "os-interface"
	detailPath   = "detail"
	actionPath   = "action"
	metadataPath = "metadata"
	ipsPath      = "ips"
	passwordPath = "os-server-password"
)

func createURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rootPath)
}

func listURL(client *golangsdk.ServiceClient) string {
	return createURL(client)
}

func listDetailURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rootPath, detailPath)
}

func deleteURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rootPath, id)
}

func getURL(client *golangsdk.ServiceClient, id string) string {
	return deleteURL(client, id)
}

func updateURL(client *golangsdk.ServiceClient, id string) string {
	return deleteURL(client, id)
}

func actionURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rootPath, id, actionPath)
}

func metadatumURL(client *golangsdk.ServiceClient, id, key string) string {
	return client.ServiceURL(rootPath, id, metadataPath, key)
}

func metadataURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rootPath, id, metadataPath)
}

func listAddressesURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rootPath, id, ipsPath)
}

func listAddressesByNetworkURL(client *golangsdk.ServiceClient, id, network string) string {
	return client.ServiceURL(rootPath, id, ipsPath, network)
}

func passwordURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rootPath, id, passwordPath)
}

func getNICManagementURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rootPath, id, nicPath)
}
