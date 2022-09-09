package availabilityzones

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List will return the existing availability zones.
func List(client *golangsdk.ServiceClient) ([]AvailabilityZone, error) {
	raw, err := client.Get(client.ServiceURL("os-availability-zone"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []AvailabilityZone
	err = extract.IntoSlicePtr(raw.Body, &res, "availabilityZoneInfo")
	return res, err
}

// ListDetail will return the existing availability zones with detailed information.
func ListDetail(client *golangsdk.ServiceClient) ([]AvailabilityZone, error) {
	raw, err := client.Get(client.ServiceURL("os-availability-zone", "detail"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []AvailabilityZone
	err = extract.IntoSlicePtr(raw.Body, &res, "availabilityZoneInfo")
	return res, err
}
