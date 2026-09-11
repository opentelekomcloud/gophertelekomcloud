package peerings

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func Update(client *golangsdk.ServiceClient, peeringID string, opts UpdateOpts) (*Peering, error) {
	if opts.Name == nil && opts.Description == nil {
		return nil, fmt.Errorf("at least one of name or description must be specified")
	}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("vpc", "peerings", peeringID).Build()
	if err != nil {
		return nil, err
	}
	b, err := build.RequestBody(opts, "peering")
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
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
