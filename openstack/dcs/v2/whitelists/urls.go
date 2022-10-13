package whitelists

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

const resourcePath = "instance"

func whitelistUrl(client *golangsdk.ServiceClient, id string) string {
	url := client.ServiceURL(resourcePath, id, "whitelist")
	return strings.Replace(url, "v1.0", "v2", 1)
}
