package availablezones

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// getURL will build the get url of get function
func getURL(client *golangsdk.ServiceClient) string {
	url := strings.Split(client.Endpoint, "/v2/")[0]
	return url + "/v2/available-zones"
}
