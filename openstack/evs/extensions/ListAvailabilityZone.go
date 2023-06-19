package extensions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListAvailabilityZone will return the existing availability zones.
func ListAvailabilityZone(client *golangsdk.ServiceClient) ([]AvailabilityZone, error) {
	// GET /v3/{project_id}/os-availability-zone
	raw, err := client.Get(client.ServiceURL("os-availability-zone"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []AvailabilityZone
	err = extract.IntoSlicePtr(raw.Body, &res, "availabilityZoneInfo")
	return res, err
}

// ZoneState represents the current state of the availability zone.
type ZoneState struct {
	// Returns true if the availability zone is available
	Available bool `json:"available"`
}

// AvailabilityZone contains all the information associated with an OpenStack
// AvailabilityZone.
type AvailabilityZone struct {
	// The availability zone name
	ZoneName  string    `json:"zoneName"`
	ZoneState ZoneState `json:"zoneState"`
}
