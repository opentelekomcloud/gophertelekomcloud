package usage

import "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "os-simple-tenant-usage"

func allTenantsURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(resourcePath)
}

func getTenantURL(client *golangsdk.ServiceClient, tenantID string) string {
	return client.ServiceURL(resourcePath, tenantID)
}
