package others

import (
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListAvailableZones(client *golangsdk.ServiceClient) (*AvailableZonesResponse, error) {
	// remove projectId from endpoint
	raw, err := client.Get(strings.Replace(client.ServiceURL("available-zones"), "/"+client.ProjectID, "", -1), nil, nil)
	if err != nil {
		return nil, err
	}

	var res AvailableZonesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type AvailableZonesResponse struct {
	RegionID       string          `json:"regionId"`
	AvailableZones []AvailableZone `json:"available_zones"`
}

type AvailableZone struct {
	ID                   string `json:"id"`
	Code                 string `json:"code"`
	Name                 string `json:"name"`
	Port                 string `json:"port"`
	ResourceAvailability string `json:"resource_availability"`
}
