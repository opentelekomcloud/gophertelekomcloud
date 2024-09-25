package availablezones

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get available zones
func Get(client *golangsdk.ServiceClient) (*GetResponse, error) {
	raw, err := client.Get(getURL(client), nil, nil)
	if err != nil {
		return nil, err
	}
	var res GetResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// GetResponse response
type GetResponse struct {
	RegionID       string          `json:"regionId"`
	AvailableZones []AvailableZone `json:"available_zones"`
}

// AvailableZone for dms
type AvailableZone struct {
	ID                   string `json:"id"`
	Code                 string `json:"code"`
	Name                 string `json:"name"`
	Port                 string `json:"port"`
	ResourceAvailability string `json:"resource_availability"`
}

// getURL will build the get url of get function
func getURL(client *golangsdk.ServiceClient) string {
	url := strings.Split(client.Endpoint, "/v2/")[0]
	return url + "/v2/available-zones"
}
