package gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Specifies the ID of a VPN gateway instance.
	GatewayID string `json:"-"`
	// Specifies the name of a VPN gateway.
	// The value is a string of 1 to 64 characters,
	// which can contain digits, letters, underscores (_), hyphens (-), and periods (.).
	Name string `json:"name,omitempty"`
	// Specifies a local subnet. This subnet is a cloud-side subnet that needs to communicate with
	// an on-premises network through a VPN. For example, a local subnet can be 192.168.52.0/24.
	LocalSubnets []string `json:"local_subnets,omitempty"`
	// Whether to enable Default Route Table Propagation.
	// Specifies the first EIP of the VPN gateway using the active-active mode or the active EIP of the VPN gateway using the active-standby mode.
	// Set this parameter only when network_type is set to public.
	Eip1 *Eip `json:"eip1,omitempty"`
	// Specifies the second EIP of the VPN gateway using the active-active mode or the standby EIP of the VPN gateway using the active-standby mode.
	// Set this parameter only when network_type is set to public.
	Eip2 *Eip `json:"eip2,omitempty"`
	// Specifies a policy template.
	// This parameter is returned only for a VPN gateway that supports access via non-fixed IP addresses.
	PolicyTemplate *PolicyTemplate `json:"policy_template,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*Gateway, error) {
	b, err := build.RequestBody(opts, "instance")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("vpn-gateways", opts.GatewayID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Gateway
	return &res, extract.Into(raw.Body, &res)
}
