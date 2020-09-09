package block_devices

import "github.com/opentelekomcloud/gophertelekomcloud"

func getURL(c *golangsdk.ServiceClient, server_id string, volume_id string) string {
	return c.ServiceURL("cloudservers", server_id, "block_device", volume_id)
}
