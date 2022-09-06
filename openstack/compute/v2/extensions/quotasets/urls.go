package quotasets

import "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "os-quota-sets"

func getURL(client *golangsdk.ServiceClient, tenantID string) string {
	return client.ServiceURL(resourcePath, tenantID)
}

func getDetailURL(client *golangsdk.ServiceClient, tenantID string) string {
	return client.ServiceURL(resourcePath, tenantID, "detail")
}

func updateURL(client *golangsdk.ServiceClient, tenantID string) string {
	return getURL(client, tenantID)
}

func deleteURL(client *golangsdk.ServiceClient, tenantID string) string {
	return getURL(client, tenantID)
}
