package maintainwindows

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get maintain windows
func Get(client *golangsdk.ServiceClient) (*GetResponse, error) {
	url := strings.Replace(client.ServiceURL("instances", "maintain-windows"), "/"+client.ProjectID, "", -1)
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}
	var s GetResponse
	err = extract.Into(raw.Body, &s)
	return &s, err
}

// GetResponse response
type GetResponse struct {
	MaintainWindows []MaintainWindow `json:"maintain_windows"`
}

// MaintainWindow for dms
type MaintainWindow struct {
	ID      int    `json:"seq"`
	Begin   string `json:"begin"`
	End     string `json:"end"`
	Default bool   `json:"default"`
}
