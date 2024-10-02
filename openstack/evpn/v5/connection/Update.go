package connection

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Specifies the ID of a VPN connection id.
	ConnectionID string `json:"-"`
	// Specifies the name of a VPN connection.
	Name string `json:"name,omitempty"`
	// Specifies a customer gateway ID.
	CgwId string `json:"cgw_id,omitempty"`
	// Specifies a customer subnet.
	PeerSubnets []string `json:"peer_subnets,omitempty"`
	// Specifies the tunnel interface address configured on the VPN gateway in route-based mode, for example, 169.254.76.1/30.
	TunnelLocalAddress string `json:"tunnel_local_address,omitempty"`
	// Specifies the tunnel interface address configured on the customer gateway device in route-based mode, for example, 169.254.76.1/30.
	TunnelPeerAddress string `json:"tunnel_peer_address,omitempty"`
	// Specifies a pre-shared key. When the IKE version is v2 and only this parameter is modified, the modification does not take effect.
	Psk string `json:"psk,omitempty"`
	// Specifies policy rules, which are returned only when style is set to POLICY.
	PolicyRules []PolicyRules `json:"policy_rules,omitempty"`
	// Specifies the IKE policy object.
	IkePolicy *IkePolicy `json:"ikepolicy,omitempty"`
	// Specifies the IPsec policy object.
	IpSecPolicy *IpSecPolicy `json:"ipsecpolicy,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*Connection, error) {
	b, err := build.RequestBody(opts, "vpn_connection")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("vpn-connection", opts.ConnectionID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Connection
	return &res, extract.IntoStructPtr(raw.Body, &res, "vpn_connection")
}
