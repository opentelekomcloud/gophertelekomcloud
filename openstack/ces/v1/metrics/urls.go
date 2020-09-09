package metrics

import "github.com/opentelekomcloud/gophertelekomcloud"

func getMetricsURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("metrics")
}
