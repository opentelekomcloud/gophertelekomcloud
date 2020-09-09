package stackevents

import "github.com/opentelekomcloud/gophertelekomcloud"

func listURL(c *golangsdk.ServiceClient, stackName, stackID string) string {
	return c.ServiceURL("stacks", stackName, stackID, "events")
}
