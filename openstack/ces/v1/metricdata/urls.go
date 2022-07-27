package metricdata

import "github.com/opentelekomcloud/gophertelekomcloud"

// POST /V1.0/{project_id}/batch-query-metric-data
func batchQueryMetricDataURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("batch-query-metric-data")
}

// GET /V1.0/{project_id}/metric-data
// POST /V1.0/{project_id}/metric-data
func metricDataURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("metric-data")
}

// GET /V1.0/{project_id}/event-data
func eventDataURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("event-data")
}
