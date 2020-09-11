package users

import "github.com/opentelekomcloud/gophertelekomcloud"

func baseURL(c *golangsdk.ServiceClient, instanceID string) string {
	return c.ServiceURL("instances", instanceID, "users")
}

func userURL(c *golangsdk.ServiceClient, instanceID, userName string) string {
	return c.ServiceURL("instances", instanceID, "users", userName)
}
