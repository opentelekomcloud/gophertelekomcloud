package others

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Get maintain windows
func ListMaintenanceWindows(client *golangsdk.ServiceClient) (r GetResult3) {
	// remove projectid from endpoint
	_, r.Err = client.Get(strings.Replace(client.ServiceURL("instances/maintain-windows"), "/"+client.ProjectID, "", -1), &r.Body, nil)
	return
}
