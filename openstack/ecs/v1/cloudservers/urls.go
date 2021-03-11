package cloudservers

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath = "cloudservers"
	jobsPath = "jobs"
)

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath)
}

func deleteURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, "delete")
}

func getURL(c *golangsdk.ServiceClient, serverID string) string {
	return c.ServiceURL(rootPath, serverID)
}

func jobURL(c *golangsdk.ServiceClient, jobID string) string {
	return c.ServiceURL(jobsPath, jobID)
}
