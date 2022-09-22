package others

import (
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Get available zones
func ListAvailableZones(client *golangsdk.ServiceClient) (r GetResult1) {
	// remove projectid from endpoint
	raw, err := client.Get(strings.Replace(client.ServiceURL("availableZones"), "/"+client.ProjectID, "", -1), nil, nil)
	return
}