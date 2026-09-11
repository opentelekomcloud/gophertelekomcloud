package peerings

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name           string  `json:"name" required:"true"`
	Description    string  `json:"description,omitempty"`
	RequestVpcInfo VpcInfo `json:"request_vpc_info" required:"true"`
	AcceptVpcInfo  VpcInfo `json:"accept_vpc_info" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Peering, error) {
	b, err := build.RequestBody(opts, "peering")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("vpc", "peerings"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
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
