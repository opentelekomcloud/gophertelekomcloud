package publicips

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func CreateURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("publicips")
}

func DeleteURL(c *golangsdk.ServiceClient, publicipId string) string {
	return c.ServiceURL("publicips", publicipId)
}

func GetURL(c *golangsdk.ServiceClient, publicipId string) string {
	return c.ServiceURL("publicips", publicipId)
}

func ListURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("publicips")
}

func UpdateURL(c *golangsdk.ServiceClient, publicipId string) string {
	return c.ServiceURL("publicips", publicipId)
}
