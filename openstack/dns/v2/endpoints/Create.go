package endpoints

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts contains the parameters of the Create function.
type CreateOpts struct {
	// Endpoint name.
	// It can contain 1 to 64 characters. Only letters, digits, underscores (_), hyphens (-), and periods (.) are allowed.
	Name string `json:"name" required:"true"`
	// Endpoint direction. The value can be:
	//   inbound: indicates an inbound endpoint.
	//   outbound: indicates an outbound endpoint.
	Direction string `json:"direction" required:"true"`
	// Region of the subnet. Example: eu-de
	Region string `json:"region" required:"true"`
	// IP address and subnet of the endpoint.
	// You need to add at least two IP addresses and can add a maximum of six IP addresses.
	IPAddresses []IPAddressInfo `json:"ipaddresses" required:"true"`
}

// IPAddressInfo is an IP address of the endpoint.
type IPAddressInfo struct {
	// Subnet ID.
	SubnetID string `json:"subnet_id" required:"true"`
	// Custom IP address, which must fall into the subnet.
	IP string `json:"ip,omitempty"`
}

// Create is used to create an endpoint.
// Send POST /v2.1/endpoints
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Endpoint, error) {
	// build.RequestBody does not validate tags of slice elements.
	for _, ip := range opts.IPAddresses {
		if err := build.ValidateTags(ip); err != nil {
			return nil, err
		}
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(v21URL(client, "endpoints"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res Endpoint
	err = extract.IntoStructPtr(raw.Body, &res, "endpoint")
	return &res, err
}
