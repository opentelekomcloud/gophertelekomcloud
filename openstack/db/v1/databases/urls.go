package databases

import "github.com/opentelekomcloud/gophertelekomcloud"

func baseURL(c *golangsdk.ServiceClient, instanceID string) string {
	return c.ServiceURL("instances", instanceID, "databases")
}

func dbURL(c *golangsdk.ServiceClient, instanceID, dbName string) string {
	return c.ServiceURL("instances", instanceID, "databases", dbName)
}
