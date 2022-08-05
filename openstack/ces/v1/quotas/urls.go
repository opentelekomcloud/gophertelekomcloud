package quotas

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// GET /V1.0/{project_id}/quotas
func quotasURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("quotas")
}
