package others

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Get maintain windows
func ListMaintenanceWindows(client *golangsdk.ServiceClient) (r GetResult3) {
	// remove projectid from endpoint
	raw, err := client.Get(strings.Replace(client.ServiceURL("instances/maintain-windows"), "/"+client.ProjectID, "", -1), nil, nil)
	return
}
