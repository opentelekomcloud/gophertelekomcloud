package tracker

import "github.com/opentelekomcloud/gophertelekomcloud"

const rootPath = "tracker"
const trackerName = "system"

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath)
}

func resourceURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, trackerName)
}

func deleteURL(c *golangsdk.ServiceClient, trackerName string) string {
	deletePath := rootPath + "?tracker_name=" + trackerName
	return c.ServiceURL(deletePath)
}
