package metrics

import "github.com/opentelekomcloud/gophertelekomcloud"

// GET /V1.0/{project_id}/metrics
func metricsURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("metrics")
}
