package alarms

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func alarmsURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("alarms")
}
