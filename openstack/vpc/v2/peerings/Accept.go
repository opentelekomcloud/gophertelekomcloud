package peerings

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Accept(client *golangsdk.ServiceClient, peeringID string) (*Peering, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("vpc", "peerings", peeringID, "accept").Build()
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL(url.String()), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res Peering
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
