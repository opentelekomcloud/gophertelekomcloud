package nameservers

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func baseURL(c *golangsdk.ServiceClient, zoneID string) string {
	return c.ServiceURL("zones", zoneID, "nameservers")
}
