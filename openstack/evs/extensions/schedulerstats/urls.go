package schedulerstats

import "github.com/opentelekomcloud/gophertelekomcloud"

func storagePoolsListURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "get_pools")
}
