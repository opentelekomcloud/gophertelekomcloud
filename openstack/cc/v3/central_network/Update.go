package central_network

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Central network ID.
	CentralNetworkId string `json:"-" required:"true"`
	// Central network name.
	Name string `json:"name,omitempty"`
	// Resource description. Angle brackets (<>) are not allowed.
	Description *string `json:"description,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*CentralNetwork, error) {
	b, err := build.RequestBody(opts, "central_network")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL(client.DomainID, "gcn", "central-networks", opts.CentralNetworkId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	return extra(raw)
}
