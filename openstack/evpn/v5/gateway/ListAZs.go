package gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListAZs(client *golangsdk.ServiceClient) (*AvailabilityZones, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("vpn-gateways", "availability-zones").Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res AvailabilityZones
	err = extract.IntoStructPtr(raw.Body, &res, "availability_zones")
	return &res, err
}

type AvailabilityZones struct {
	// Indicates that the specification of VPN gateways is Basic.
	Basic *VpnGatewayAvailabilityZones `json:"basic"`
	// Indicates that the specification of VPN gateways is Professional1.
	Professional1 *VpnGatewayAvailabilityZones `json:"professional1"`
	// Indicates that the specification of VPN gateways is Professional1-NonFixedIP.
	Professional1NonFixedIP *VpnGatewayAvailabilityZones `json:"Professional1-NonFixedIP"`
	// Indicates that the specification of VPN gateways is Professional2.
	Professional2 *VpnGatewayAvailabilityZones `json:"professional2"`
	// Indicates that the specification of VPN gateways is Professional2-NonFixedIP.
	Professional2NonFixedIP *VpnGatewayAvailabilityZones `json:"Professional2-NonFixedIP"`
	Gm                      *VpnGatewayAvailabilityZones `json:"gm"`
}

type VpnGatewayAvailabilityZones struct {
	// Specifies the list of AZs for VPN gateways associated with VPCs.
	Vpc []string `json:"vpc"`
	// Specifies the list of AZs for VPN gateways associated with enterprise routers.
	Er []string `json:"er"`
}
