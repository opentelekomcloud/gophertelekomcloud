package zones

import "github.com/opentelekomcloud/gophertelekomcloud"

const rootPath = "zones"

func baseURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath)
}

func zoneURL(c *golangsdk.ServiceClient, zoneID string) string {
	return c.ServiceURL(rootPath, zoneID)
}

func associateURL(client *golangsdk.ServiceClient, zoneID string) string {
	return client.ServiceURL(rootPath, zoneID, "associaterouter")
}

func disassociateURL(client *golangsdk.ServiceClient, zoneID string) string {
	return client.ServiceURL(rootPath, zoneID, "disassociaterouter")
}
