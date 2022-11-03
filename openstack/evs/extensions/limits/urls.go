package limits

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

const resourcePath = "limits"

func getURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}
