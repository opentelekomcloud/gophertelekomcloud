package services

import "github.com/opentelekomcloud/gophertelekomcloud"

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("os-services")
}
