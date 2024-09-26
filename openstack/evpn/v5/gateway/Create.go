package gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	// Specifies the name of a VPN gateway.
	// The value is a string of 1 to 64 characters, which can contain digits, letters, underscores (_), hyphens (-), and periods (.).
	// If this parameter is not specified, a name in the format of vpngw-**** is automatically generated, for example, vpngw-a45b.
	Name string `json:"name,omitempty"`
	// Specifies the network type of the VPN gateway. A public VPN gateway (public) uses EIPs to connect to a customer gateway. A private VPN gateway (private) uses private IP addresses in a VPC to connect to a customer gateway.
	// The value can be public or private.
	// The default value is public.
	NetworkType string `json:"network_type,omitempty"`
	// Specifies the association mode.
	// The value can be vpc or er.
	// The default value is vpc.
	AttachmentType string `json:"attachment_type,omitempty"`
	// Specifies the ID of the enterprise router instance to which the VPN gateway connects.
	// The value is a UUID containing 36 characters.
	// This parameter is mandatory when attachment_type is set to er, and cannot be configured when attachment_type is set to vpc.
	ErId string `json:"er_id,omitempty"`
	// When attachment_type is set to vpc, vpc_id specifies the ID of the service VPC associated with the VPN gateway.
	// When attachment_type is set to er, vpc_id specifies the ID of the access VPC used by the VPN gateway. In this case, any VPC ID can be used.
	VpcId string `json:"vpc_id,omitempty"`
	// Specifies a local subnet. This subnet is a cloud-side subnet that needs to communicate with an on-premises customer subnet through a VPN.
	// Set this parameter only when attachment_type is set to vpc.
	LocalSubnets []string `json:"local_subnets,omitempty"`
	// Specifies the ID of the VPC subnet used by the VPN gateway.
	ConnectSubnet string `json:"connect_subnet,omitempty"`
	// Specifies the BGP AS number of the VPN gateway.
	// The value ranges from 1 to 4294967295.
	// The default value is 64512.
	BgpAsn int `json:"bgp_asn,omitempty"`
	// Specifies the specifications of the VPN gateway.
	// For the value range, see the Specification parameter on the page for creating a VPN gateway on the VPN console.
	// Value range:
	// V1G
	// V300
	// Basic
	// Professional1
	// Professional2
	// Professional1-NonFixedIP (not yet supported)
	// Professional2-NonFixedIP (not yet supported)
	// This parameter cannot be set to Basic when network_type is private or when attachment_type is er.
	// The default value is Professional1.
	Flavor string `json:"flavor,omitempty"`
	// Specifies the AZ where the VPN gateway is to be deployed. If this parameter is not specified,
	// an AZ is automatically selected for the VPN gateway. You can obtain the AZ list by referring to Querying the AZs of VPN Gateways.
	AvailabilityZoneIds []string `json:"availability_zone_ids,omitempty"`
	// Specifies the first EIP of the VPN gateway using the active-active mode or the active EIP of the VPN gateway using the active-standby mode.
	// Set this parameter only when network_type is set to public.
	Eip1 *Eip `json:"eip1,omitempty"`
	// Specifies the second EIP of the VPN gateway using the active-active mode or the standby EIP of the VPN gateway using the active-standby mode.
	// Set this parameter only when network_type is set to public.
	Eip2 *Eip `json:"eip2,omitempty"`
	// Specifies the ID of the access VPC used by the VPN gateway.
	// This parameter is optional when attachment_type is set to vpc.
	// If both access_vpc_id and vpc_id are set, both of them take effect.
	// When attachment_type is set to er, set either vpc_id or access_vpc_id.
	// Setting access_vpc_id is recommended. If both access_vpc_id and vpc_id are set, only access_vpc_id takes effect.
	// By default, the value is the same as the value of vpc_id.
	AccessVpcId string `json:"access_vpc_id,omitempty"`
	// Specifies the ID of the subnet in the access VPC used by the VPN gateway.
	// This parameter is optional when attachment_type is set to vpc. If both access_subnet_id and connect_subnet are set and their values are the same,
	// ensure that the subnet has at least four available IP addresses.
	// If both access_subnet_id and connect_subnet are set and their values are different,
	// ensure that each subnet has at least two available IP addresses.
	// When attachment_type is set to er, set either access_subnet_id or connect_subnet.
	// Setting access_subnet_id is recommended. If both access_subnet_id and connect_subnet are set,
	// only access_subnet_id takes effect. Ensure that the subnet has at least two available IP addresses.
	// By default, the value is the same as the value of connect_subnet.
	AccessSubnetId string `json:"access_subnet_id,omitempty"`
	// Specifies the HA mode of the gateway. The value can be active-active or active-standby.
	// Value range: active-active, active-standby
	// Default value: active-active
	HaMode string `json:"ha_mode,omitempty"`
	// Specifies private IP address 1 of a private VPN gateway.
	// Set this parameter if a private VPN gateway needs to use specified IP addresses.
	// In active/standby gateway mode, the specified IP address is the active IP address.
	// In active-active gateway mode, the specified IP address is active IP address 1.
	// Value range: allocatable IP addresses in the access subnet
	// This parameter must be specified together with access_private_ip_2, and the two parameters must have different values.
	AccessPrivateIp1 string `json:"access_private_ip_1,omitempty"`
	// Specifies private IP address 2 of a private VPN gateway.
	// Set this parameter if a private VPN gateway needs to use specified IP addresses.
	// In active/standby gateway mode, the specified IP address is the standby IP address.
	// In active-active gateway mode, the specified IP address is active IP address 2.
	// Value range: allocatable IP addresses in the access subnet
	// This parameter must be specified together with access_private_ip_1, and the two parameters must have different values.
	AccessPrivateIp2 string `json:"access_private_ip_2,omitempty"`
	// Specifies a tag list.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

type Eip struct {
	// Specifies an EIP ID.
	ID string `json:"id,omitempty"`
	// Specifies the EIP type.
	Type string `json:"type,omitempty"`
	// Specifies the bandwidth billing mode of an EIP.
	// Value range:
	// traffic: billed by traffic
	ChargeMode string `json:"charge_mode,omitempty"`
	// Specifies the bandwidth (Mbit/s) of an EIP. The maximum EIP bandwidth varies according to regions and depends on the EIP service.
	// You can submit a service ticket to increase the maximum EIP bandwidth under your account.
	BandwidthSize int `json:"bandwidth_size,omitempty"`
	// Specifies the bandwidth name of an EIP.
	BandwidthName string `json:"bandwidth_name,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Gateway, error) {
	b, err := build.RequestBody(opts, "vpn_gateway")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("vpn-gateways"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res Gateway
	return &res, extract.IntoStructPtr(raw.Body, &res, "vpn_gateway")
}

type Gateway struct {
	// Specifies a VPN gateway ID.
	ID string `json:"id"`
	// Specifies a VPN gateway name. If no VPN gateway name is specified, the system automatically generates one.
	Name string `json:"name"`
	// Specifies the network type of the VPN gateway.
	NetworkType string `json:"network_type"`
	// Specifies the association mode.
	AttachmentType string `json:"attachment_type"`
	// Specifies the ID of the enterprise router instance to which the VPN gateway connects.
	// This parameter is available only when attachment_type is set to er.
	ErId string `json:"er_id"`
	// When attachment_type is set to vpc, vpc_id specifies the ID of the service VPC associated with the VPN gateway.
	VpcId string `json:"vpc_id"`
	// Specifies a local subnet.
	LocalSubnets []string `json:"local_subnets"`
	// Specifies the ID of the VPC subnet used by the VPN gateway.
	ConnectSubnet string `json:"connect_subnet"`
	// Specifies the BGP AS number of the VPN gateway.
	BgpAsn int `json:"bgp_asn"`
	// Specifies the specification of the VPN gateway.
	Flavor string `json:"flavor"`
	// Specifies the maximum number of VPN connections supported for the VPN gateway.
	ConnectionNumber int `json:"connection_number"`
	// Specifies the number of VPN connections that have been used by the VPN gateway.
	UsedConnectionNumber int `json:"used_connection_number"`
	// Specifies the number of VPN connection groups that have been used by the VPN gateway.
	// A connection group consists of two connections between a customer gateway and a VPN gateway.
	// By default, 10 VPN connection groups are included free of charge with the purchase of a VPN gateway.
	UsedConnectionGroup int `json:"used_connection_group"`
	// Specifies the ID of the access VPC used by the VPN gateway.
	AccessVpcId string `json:"access_vpc_id"`
	// Specifies the ID of the subnet in the access VPC used by the VPN gateway.
	AccessSubnetId string `json:"access_subnet_id"`
	// Specifies the HA mode of the gateway. The value can be active-active or active-standby.
	HaMode string `json:"ha_mode"`
	// Specifies a policy template.
	// This parameter is returned only for a VPN gateway that supports access via non-fixed IP addresses.
	PolicyTemplate *PolicyTemplate `json:"policy_template"`
	// Specifies a tag list.
	Tags []tags.ResourceTag `json:"tags"`
	// Specifies the status of the VPN gateway.
	Status string `json:"status"`
	// Specifies the first EIP of the VPN gateway using the active-active mode
	// or the active EIP of the VPN gateway using the active-standby mode.
	Eip1 EipResp `json:"eip1"`
	// Specifies the second EIP of the VPN gateway using the active-active mode
	// or the standby EIP of the VPN gateway using the active-standby mode.
	Eip2 EipResp `json:"eip2"`
	// Specifies the time when the VPN gateway is created.
	CreatedAt string `json:"created_at"`
	// Specifies the last update time.
	UpdatedAt string `json:"updated_at"`
	// Specifies whether a frozen VPN gateway can be deleted.
	// The value 1 indicates that a frozen gateway can be deleted.
	// The value 2 indicates that a frozen gateway cannot be deleted.
	LockStatus int `json:"lock_status"`
	// Specifies a private IP address used by the VPN gateway to connect to
	// a customer gateway when the network type is private network.
	AccessPrivateIp1 string `json:"access_private_ip_1"`
	// Specifies a private IP address used by the VPN gateway to connect to
	// a customer gateway when the network type is private network.
	AccessPrivateIp2 string `json:"access_private_ip_2"`
}

type PolicyTemplate struct {
	// Specifies the IKE policy object.
	IkePolicy *IkePolicyResp `json:"ike_policy,omitempty"`
	// Specifies the IPsec policy object.
	IpsecPolicy *IpsecPolicyResp `json:"ipsec_policy,omitempty"`
}

type IkePolicyResp struct {
	// Specifies an encryption algorithm.
	EncryptionAlgorithm string `json:"encryption_algorithm,omitempty"`
	// Specifies the DH group used for key exchange in phase 1.
	DhGroup string `json:"dh_group,omitempty"`
	// Specifies an authentication algorithm.
	AuthenticationAlgorithm string `json:"authentication_algorithm,omitempty"`
	// Specifies the SA lifetime. When the lifetime expires, an IKE SA is automatically updated.
	LifetimeSeconds string `json:"lifetime_seconds,omitempty"`
}

type IpsecPolicyResp struct {
	// Specifies an encryption algorithm.
	EncryptionAlgorithm string `json:"encryption_algorithm,omitempty"`
	// Specifies the DH key group used by PFS.
	Pfs string `json:"pfs,omitempty"`
	// Specifies an authentication algorithm.
	AuthenticationAlgorithm string `json:"authentication_algorithm,omitempty"`
	// Specifies the lifetime of a tunnel established over an IPsec connection.
	LifetimeSeconds string `json:"lifetime_seconds,omitempty"`
}

type EipResp struct {
	// Specifies an EIP ID.
	ID string `json:"id"`
	// Specifies the EIP version.
	IpVersion int `json:"ip_version"`
	// Specifies the EIP type.
	Type string `json:"type"`
	// Specifies an EIP, that is, a public IPv4 address.
	IpAddress string `json:"ip_address"`
	// Specifies the bandwidth billing mode of an EIP.
	ChargeMode string `json:"charge_mode"`
	// Specifies the bandwidth ID of an EIP.
	BandwidthId string `json:"bandwidth_id"`
	// Specifies the bandwidth (Mbit/s) of an EIP.
	BandwidthSize int `json:"bandwidth_size"`
	// Specifies the bandwidth name of an EIP.
	BandwidthName string `json:"bandwidth_name"`
	// Specifies the type of EIP bandwidth.
	ShareType string `json:"share_type"`
	// Specifies the EIP type.
	NetworkType string `json:"network_type"`
}
