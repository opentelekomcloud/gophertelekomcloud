package availability_zones

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Enterprise router ID
	ID string `q:"instance_id"`
	// Bandwidth size, in Mbit/s
	BandwidthSize int `q:"bandwidth_size"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]AZsResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("enterprise-router", "availability-zones").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []AZsResponse
	err = extract.IntoSlicePtr(raw.Body, &res, "availability_zones")
	return res, err
}

type AZsResponse struct {
	// AZ code
	Code string `json:"code"`
	// Whether the AZ is available. Value options: available and unavailable
	State string `json:"state"`
}
