package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func createURL(client *golangsdk.ServiceClient, server_id string) string {
	return client.ServiceURL("servers", server_id, "tags")
}

func getURL(client *golangsdk.ServiceClient, server_id string) string {
	return client.ServiceURL("servers", server_id, "tags")
}

func deleteURL(client *golangsdk.ServiceClient, server_id string) string {
	return client.ServiceURL("servers", server_id, "tags")
}
