package peerings

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, peeringID string) (*Peering, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("vpc", "peerings", peeringID).Build()
	if err != nil {
		return nil, err
	}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		Peering Peering `json:"peering"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.Peering, nil
}
