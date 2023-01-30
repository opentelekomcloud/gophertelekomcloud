package others

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListMaintenanceWindows(client *golangsdk.ServiceClient) ([]MaintainWindow, error) {
	// remove projectId from endpoint
	raw, err := client.Get(strings.Replace(client.ServiceURL("instances/maintain-windows"), "/"+client.ProjectID, "", -1), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []MaintainWindow
	err = extract.IntoSlicePtr(raw.Body, &res, "maintain_windows")
	return res, err
}

type MaintainWindow struct {
	// Start time of the maintenance time window.
	ID int `json:"seq"`
	// Start time of the maintenance time window.
	Begin string `json:"begin"`
	// End time of the maintenance time window.
	End string `json:"end"`
	// An indicator of whether the maintenance time window is set to the default time segment.
	Default bool `json:"default"`
}
