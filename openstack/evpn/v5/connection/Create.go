package connection

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	// Specifies the name of a VPN connection.
	// If this parameter is not specified, a name in the format of vpn-**** is automatically generated, for example, vpn-13be.
	// The value is a string of 1 to 64 characters,
	// which can contain digits, letters, underscores (_), hyphens (-), and periods (.).
	Name string `json:"name,omitempty"`
	// Specifies a VPN gateway ID.
	// The value is a UUID containing 36 characters.
	VgwId string `json:"vgw_id" required:"true"`
	// When network_type of the VPN gateway is set to public, set vgw_ip to the EIP IDs of the VPN gateway.
	// When network_type of the VPN gateway is set to private, set vgw_ip to the private IP addresses of the VPN gateway.
	// The value is a UUID containing 36 characters or an IPv4 address in dotted decimal notation (for example, 192.168.45.7).
	VgwIp string `json:"vgw_ip" required:"true"`
	// Specifies the connection mode.
	// Value range:
	// policy: policy-based mode
	// static: static routing mode
	// bgp: BGP routing mode
	// policy-template: policy template mode
	// The default value is static.
	Style string `json:"style,omitempty"`
	// Specifies a customer gateway ID.
	// The value is a UUID containing 36 characters.
	CgwId string `json:"cgw_id" required:"true"`
	// Specifies a customer subnet.
	// This parameter is not required when the association mode of the VPN gateway
	// is set to er and style is set to policy or bgp. This parameter is mandatory in other scenarios.
	// Reserved VPC CIDR blocks such as 100.64.0.0/10 cannot be used as customer subnets.
	// A maximum of 50 customer subnets can be configured for each VPN connection.
	PeerSubnets []string `json:"peer_subnets,omitempty"`
	// Specifies the tunnel interface address configured on the VPN gateway in route-based mode, for example, 169.254.76.1/30.
	// Constraints:
	// The first 16 bits must be 169.254, and the value cannot be 169.254.195.xxx.
	// The mask length must be 30, and the address must be in the same CIDR block as the value of tunnel_peer_address.
	// The address needs to be a host address in a CIDR block.
	TunnelLocalAddress string `json:"tunnel_local_address,omitempty"`
	// Specifies the tunnel interface address configured on the customer gateway device in route-based mode, for example, 169.254.76.2/30.
	// Constraints:
	// The first 16 bits must be 169.254, and the value cannot be 169.254.195.xxx.
	// The mask length must be 30, and the address must be in the same CIDR block as the value of tunnel_local_address.
	// The address needs to be a host address in a CIDR block.
	TunnelPeerAddress string `json:"tunnel_peer_address,omitempty"`
	// Specifies whether to enable the network quality analysis (NQA) function.
	// The value can be true or false.
	// The default value is false.
	// Set this parameter only when style is set to static.
	EnableNqa *bool `json:"enable_nqa,omitempty"`
	// Specifies a pre-shared key.
	// The value is a string of 8 to 128 characters, which must contain at least three types of the
	// following: uppercase letters, lowercase letters, digits, and special characters ~!@#$%^()-_+={ },./:;.
	Psk string `json:"psk" required:"true"`
	// Specifies policy rules.
	// A maximum of five policy rules can be specified. Set this parameter only when style is set to policy.
	PolicyRules []PolicyRules `json:"policy_rules,omitempty"`
	// Specifies the Internet Key Exchange (IKE) policy object.
	IkePolicy *IkePolicy `json:"ikepolicy,omitempty"`
	// Specifies the Internet Protocol Security (IPsec) policy object.
	IpSecPolicy *IpSecPolicy `json:"ipsecpolicy,omitempty"`
	// This parameter is optional when you create a connection for a VPN gateway in active-active mode.
	// When you create a connection for a VPN gateway in active-standby mode,
	// master indicates the active connection, and slave indicates the standby connection.
	// The default value is master.
	HaRole string `json:"ha_role,omitempty"`
	// Specifies a tag list.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

type IkePolicy struct {
	// Specifies the IKE version.
	// Value range:
	// v1 and v2
	// Default value:
	// v2
	IkeVersion string `json:"ike_version,omitempty"`
	// Specifies the negotiation mode.
	// Value range:
	// main: ensures high security during negotiation.
	// aggressive: ensures fast negotiation and a high negotiation success rate.
	// The default value is main.
	// This parameter is mandatory only when the IKE version is v1.
	PhaseOneNegotiationMode string `json:"phase1_negotiation_mode,omitempty"`
	// Specifies an authentication algorithm.
	// Value range:
	// sha2-512, sha2-384, sha2-256, sha1, md5
	// Exercise caution when using sha1 and md5 as they have low security.
	// Default value:
	// sha2-256
	AuthenticationAlgorithm string `json:"authentication_algorithm,omitempty"`
	// Specifies an encryption algorithm.
	// Value range:
	// aes-256-gcm-16, aes-128-gcm-16, aes-256, aes-192, aes-128, 3des
	// Exercise caution when using 3des as it has low security.
	// Default value:
	// aes-128
	EncryptionAlgorithm string `json:"encryption_algorithm,omitempty"`
	// Specifies the DH group used for key exchange in phase 1.
	// The value can be group1, group2, group5, group14, group15, group16, group19, group20, or group21.
	// Exercise caution when using group1, group2, group5, or group14 as they have low security.
	// The default value is group15.
	DhGroup string `json:"dh_group,omitempty"`
	// Specifies the authentication method used during IKE negotiation.
	// Value range:
	// pre-share: pre-shared key
	// Default value: pre-share
	AuthenticationMethod string `json:"authentication_method,omitempty"`
	// Specifies the security association (SA) lifetime.
	// When the lifetime expires, an IKE SA is automatically updated.
	// The value ranges from 60 to 604800, in seconds.
	// The default value is 86400.
	LifetimeSeconds *int `json:"lifetime_seconds,omitempty"`
	// Specifies the local ID type.
	// Value range:
	// ip
	// fqdn (currently not supported)
	// The default value is ip.
	LocalIdType string `json:"local_id_type,omitempty"`
	// Specifies the local ID.
	// The value can contain a maximum of 255 case-sensitive characters,
	// including letters, digits, and special characters (excluding & < > [ ] \).
	// Spaces are not supported. Set this parameter when local_id_type is set to fqdn.
	// The value must be the same as that of peer_id on the peer device.
	LocalId string `json:"local_id,omitempty"`
	// Specifies the peer ID type.
	// Value range:
	// ip
	// fqdn (currently not supported)
	// The default value is ip.
	PeerIdType string `json:"peer_id_type,omitempty"`
	// Specifies the peer ID.
	// The value can contain a maximum of 255 case-sensitive characters,
	// including letters, digits, and special characters (excluding & < > [ ] \).
	// Spaces are not supported. Set this parameter when peer_id_type is set to fqdn.
	// The value must be the same as that of local_id on the peer device.
	PeerId string `json:"peer_id,omitempty"`
	// Specifies the dead peer detection (DPD) object.
	Dpd *Dpd `json:"dpd,omitempty"`
}

type Dpd struct {
	// Specifies the interval for retransmitting DPD packets.
	// The value ranges from 2 to 60, in seconds.
	// The default value is 15.
	Timeout *int `json:"timeout,omitempty"`
	// Specifies the DPD idle timeout period.
	// The value ranges from 10 to 3600, in seconds.
	// The default value is 30.
	Interval *int `json:"interval,omitempty"`
	// Specifies the format of DPD packets.
	// Value range:
	// seq-hash-notify: indicates that the payload of DPD packets is in the sequence of hash-notify.
	// seq-notify-hash: indicates that the payload of DPD packets is in the sequence of notify-hash.
	// The default value is seq-hash-notify.
	Msg string `json:"msg,omitempty"`
}

type IpSecPolicy struct {
	// Specifies an authentication algorithm.
	// Value range:
	// sha2-512, sha2-384, sha2-256, sha1, md5
	// Exercise caution when using sha1 and md5 as they have low security.
	// Default value:
	// sha2-256
	AuthenticationAlgorithm string `json:"authentication_algorithm,omitempty"`
	// Specifies an encryption algorithm.
	// Value range:
	// aes-256-gcm-16, aes-128-gcm-16, aes-256, aes-192, aes-128, 3des
	// Exercise caution when using 3des as it has low security.
	// Default value:
	// aes-128
	EncryptionAlgorithm string `json:"encryption_algorithm,omitempty"`
	// Specifies the DH key group used by Perfect Forward Secrecy (PFS).
	// The value can be group1, group2, group5, group14, group15, group16, group19, group20, group21, or disable.
	// Exercise caution when using group1, group2, group5, or group14 as they have low security.
	// The default value is group15.
	Pfs string `json:"pfs,omitempty"`
	// Specifies the transfer protocol.
	// Value range:
	// esp: encapsulating security payload protocol
	// The default value is esp.
	TransformProtocol string `json:"transform_protocol,omitempty"`
	// Specifies the lifetime of a tunnel established over an IPsec connection.
	// The value ranges from 30 to 604800, in seconds.
	// The default value is 3600.
	LifetimeSeconds *int `json:"lifetime_seconds,omitempty"`
	// Specifies the packet encapsulation mode.
	// Value range:
	// tunnel: encapsulates packets in tunnel mode.
	// The default value is tunnel.
	EncapsulationMode string `json:"encapsulation_mode,omitempty"`
}

type PolicyRules struct {
	// Specifies a rule ID, which is used to identify the sequence in which the rule is configured.
	// You are advised not to set this parameter.
	// The value ranges from 0 to 50.
	// The value of rule_index in each policy rule must be unique.
	// The value of rule_index in ResponseVpnConnection may be different from the value of this parameter.
	// This is because if multiple destination CIDR blocks are specified, the VPN service generates a rule for each destination CIDR block.
	RuleIndex int `json:"rule_index,omitempty"`
	// Specifies a source CIDR block.
	// The value of source in each policy rule must be unique.
	Source string `json:"source,omitempty"`
	// Specifies a destination CIDR block.
	// For example, a destination CIDR block can be 192.168.52.0/24.
	// A maximum of 50 destination CIDR blocks can be configured in each policy rule.
	Destination []string `json:"destination,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Connection, error) {
	b, err := build.RequestBody(opts, "vpn_connection")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("vpn-connection"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202, 201},
	})
	if err != nil {
		return nil, err
	}

	var res Connection
	return &res, extract.IntoStructPtr(raw.Body, &res, "vpn_connection")
}

type Connection struct {
	// Specifies a VPN connection ID.
	ID string `json:"id"`
	// Specifies a VPN connection name. If no VPN connection name is specified, the system automatically generates one.
	Name string `json:"name"`
	// Specifies a VPN gateway ID.
	VgwId string `json:"vgw_id"`
	// Specifies an EIP ID or private IP address of the VPN gateway.
	VgwIp string `json:"vgw_ip"`
	// Specifies the connection mode.
	Style string `json:"style"`
	// Specifies a customer gateway ID.
	CgwId string `json:"cgw_id"`
	// Specifies a customer subnet.
	PeerSubnets []string `json:"peer_subnets"`
	// Specifies the tunnel interface address configured on the VPN gateway in route-based mode.
	TunnelLocalAddress string `json:"tunnel_local_address"`
	// Specifies the tunnel interface address configured on the customer gateway device in route-based mode.
	TunnelPeerAddress string `json:"tunnel_peer_address"`
	// Specifies whether NQA is enabled. This parameter is returned only when style is STATIC.
	EnableNqa bool `json:"enable_nqa"`
	// Specifies policy rules, which are returned only when style is set to POLICY.
	PolicyRules []PolicyRules `json:"policy_rules"`
	// Specifies the IKE policy object.
	IkePolicy IkePolicy `json:"ikepolicy"`
	// Specifies the IPsec policy object.
	IpSecPolicy IpSecPolicy `json:"ipsecpolicy"`
	// Specifies the ID of a VPN connection monitor.
	ConnectionMonitorId string `json:"connection_monitor_id"`
	// Specifies the time when the VPN connection is created.
	CreatedAt string `json:"created_at"`
	// Specifies the last update time.
	UpdatedAt string `json:"updated_at"`
	// For a VPN gateway in active-standby mode, master indicates the active connection,
	// and slave indicates the standby connection.
	// For a VPN gateway in active-active mode, the value of ha_role can only be master.
	HaRole string `json:"ha_role"`
	// Specifies a tag list.
	Tags []tags.ResourceTag `json:"tags"`
	// Specifies an EIP ID or private IP address of the VPN gateway.
	EipId string `json:"eip_id"`
	// Specifies the connection mode.
	// Value range:
	// POLICY: policy-based mode
	// ROUTE: routing mode
	Type string `json:"type"`
	// Specifies the routing mode.
	// Value range:
	// static: static routing mode
	// bgp: BGP routing mode
	RouteMode string `json:"route_mode"`
	// Specifies the status of the VPN connection.
	// Value range:
	// ERROR: abnormal
	// ACTIVE: normal
	// DOWN: not connected
	// PENDING_CREATE: creating
	// PENDING_UPDATE: updating
	// PENDING_DELETE: deleting
	// FREEZED: frozen
	// UNKNOWN: unknown
	Status string `json:"status"`
}
