package limits

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

const resourcePath = "limits"

func getURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(resourcePath)
}
