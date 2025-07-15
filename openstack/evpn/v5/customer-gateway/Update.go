package customer_gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Specifies the ID of a VPN gateway instance.
	GatewayID string `json:"-"`
	// Specifies a gateway name.
	// The value is a string of 1 to 64 characters, which can contain digits, letters, underscores (_), hyphens (-), and periods (.).
	Name string `json:"name,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*CustomerGateway, error) {
	b, err := build.RequestBody(opts, "customer_gateway")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("customer-gateways", opts.GatewayID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CustomerGateway
	return &res, extract.IntoStructPtr(raw.Body, &res, "customer_gateway")
}
