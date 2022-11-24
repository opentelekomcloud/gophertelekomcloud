package others

import (
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListAvailableZones(client *golangsdk.ServiceClient) (*AvailableZonesResponse, error) {
	// remove projectId from endpoint
	raw, err := client.Get(strings.Replace(client.ServiceURL("availableZones"), "/"+client.ProjectID, "", -1), nil, nil)
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
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	// Port number of the AZ.
	Port string `json:"port"`
	// An indicator of whether there are available Redis 3.0 resources in the AZ.
	// true: There are available resources in the AZ.
	// false: There are no available resources in the AZ.
	ResourceAvailability string `json:"resource_availability"`
	// An indicator of whether there are available Redis 4.0 and 5.0 resources in the AZ.
	ResourceAvailabilityDcs2 string `json:"resource_availability_dcs2"`
}
