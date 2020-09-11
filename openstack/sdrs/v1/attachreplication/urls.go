package attachreplication

import "github.com/opentelekomcloud/gophertelekomcloud"

func createURL(c *golangsdk.ServiceClient, instanceID string) string {
	return c.ServiceURL("protected-instances", instanceID, "attachreplication")
}

func deleteURL(c *golangsdk.ServiceClient, instanceID string, replicationID string) string {
	return c.ServiceURL("protected-instances", instanceID, "detachreplication", replicationID)
}
