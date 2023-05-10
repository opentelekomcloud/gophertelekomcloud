package loadbalancers

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateOpts is the common options' struct used in this package's Create operation.
type CreateOpts struct {
	// Specifies the load balancer name.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Name string `json:"name,omitempty"`
	// Provides supplementary information about the load balancer.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Description string `json:"description,omitempty"`
	// Specifies the private IPv4 address bound to the load balancer. The IP address must be from the IPv4 subnet where the load balancer resides and should not be occupied by other services.
	//
	// Note:
	//
	// vip_subnet_cidr_id is also required if vip_address is passed.
	//
	// If only vip_subnet_cidr_id is passed, the system will automatically assign a private IPv4 address to the load balancer.
	//
	// If both vip_address and vip_subnet_cidr_id are not passed, no private IPv4 address will be assigned, and the value of vip_address will be null.
	VipAddress string `json:"vip_address,omitempty"`
	// Specifies the ID of the IPv4 subnet where the load balancer resides. This parameter is mandatory if you need to create a load balancer with a private IPv4 address.
	//
	// You can query parameter neutron_subnet_id in the response by calling the API (GET https://{VPC_Endpoint}/v1/{project_id}/subnets).
	//
	// Note:
	//
	// vpc_id, vip_subnet_cidr_id and ipv6_vip_virsubnet_id cannot be left blank at the same time. The subnet specified by vip_subnet_cidr_id and the subnet specified by ipv6_vip_virsubnet_id must be in the VPC specified by vpc_id.
	//
	// The subnet specified by vip_subnet_cidr_id must be in the VPC specified by vpc_id if both vpc_id and vip_subnet_cidr_id are passed.
	//
	// Minimum: 1
	//
	// Maximum: 36
	VipSubnetCidrID string `json:"vip_subnet_cidr_id,omitempty"`
	// Specifies the ID of the IPv6 subnet where the load balancer resides. You can query id in the response by calling the API (GET https://{VPC_Endpoint}/v1/{project_id}/subnets).
	//
	// Note:
	//
	// vpc_id, vip_subnet_cidr_id and ipv6_vip_virsubnet_id cannot be left blank at the same time. The subnet specified by vip_subnet_cidr_id and the subnet specified by ipv6_vip_virsubnet_id must be in the VPC specified by vpc_id.
	//
	// IPv6 must have been enabled for the IPv6 subnet where the load balancer resides.
	//
	// IPv6 is unsupported. Please do not use this parameter.
	IpV6VipSubnetID string `json:"ipv6_vip_virsubnet_id,omitempty"`
	// Specifies the provider of the load balancer. The value can only be vlb.
	//
	// Minimum: 1
	//
	// Maximum: 255
	Provider string `json:"provider,omitempty"`
	// Specifies the ID of a flavor at Layer 4.
	//
	// Note:
	//
	// If neither l4_flavor_id nor l7_flavor_id is specified, the default flavor is used. The default flavor varies depending on the sites.
	//
	// Minimum: 1
	//
	// Maximum: 36
	L4Flavor string `json:"l4_flavor_id,omitempty"`
	// Specifies whether the load balancer is a dedicated load balancer.
	//
	// true (default): The load balancer is a dedicated load balancer.
	//
	// false: The load balancer is a shared load balancer.
	//
	// Currently, the value can only be true. If the value is set to false, 400 Bad Request will be returned.
	Guaranteed *bool `json:"guaranteed,omitempty"`
	// Specifies the ID of the VPC where the load balancer resides. You can query parameter id in the response by calling the API (GET https://{VPC_Endpoint}/v1/{project_id}/vpcs).
	//
	// vpc_id, vip_subnet_cidr_id and ipv6_vip_virsubnet_id cannot be left blank at the same time. The subnet specified by vip_subnet_cidr_id and the subnet specified by ipv6_vip_virsubnet_id must be in the VPC specified by vpc_id.
	VpcID string `json:"vpc_id,omitempty"`
	// Specifies the list of AZs where the load balancer can be created. You can query the AZs by calling the API (GET https://{ELB_Endpoint}/v3/{project_id}/elb/availability-zones). Select one or more AZs in the same set.
	AvailabilityZoneList []string `json:"availability_zone_list" required:"true"`
	// Lists the tags added to the load balancer.
	//
	// Example: "tags":[{"key":"my_tag","value":"my_tag_value"}]
	Tags []tags.ResourceTag `json:"tags,omitempty"`
	// Specifies the administrative status of the load balancer. The value can only be true (default).
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// Specifies the ID of a flavor at Layer 7.
	//
	// Note:
	//
	// If neither l4_flavor_id nor l7_flavor_id is specified, the default flavor is used. The default flavor varies depending on the sites.
	//
	// Minimum: 1
	//
	// Maximum: 36
	L7Flavor string `json:"l7_flavor_id,omitempty"`
	// Specifies the ID of the bandwidth used by an IPv6 address. This parameter is available only when you create or update a load balancer with a public IPv6 address. If you use a new IPv6 address and specify a shared bandwidth, the IPv6 address will be added to the shared bandwidth.
	//
	// IPv6 is unsupported. Please do not use this parameter.
	IPV6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`
	// Specifies the ID of the EIP the system will automatically assign and bind to the load balancer during load balancer creation. Only the first EIP will be bound to the load balancer although multiple EIP IDs can be set.
	PublicIpIDs []string `json:"publicip_ids,omitempty"`
	// Specifies the new EIP that will be bound to the load balancer.
	PublicIp *PublicIp `json:"publicip,omitempty"`
	//
	// Specifies the IDs of subnets on the downstream plane. You can query parameter neutron_network_id in the response by calling the API (GET https://{VPC_Endpoint}/v1/{project_id}/subnets).
	//
	// If this parameter is not specified, select subnets as follows:
	//
	// If IPv6 is enabled for a load balancer, the ID of subnet specified in ipv6_vip_virsubnet_id will be used.
	//
	// If IPv4 is enabled for a load balancer, the ID of subnet specified in vip_subnet_cidr_id will be used.
	//
	// If only pubilc network is available for a load balancer, the ID of any subnet in the VPC where the load balancer resides will be used. Subnets with more IP addresses are preferred.
	//
	// If there is more than one subnet, the first subnet in the list will be used.
	//
	// The subnets must be in the VPC where the load balancer resides.
	//
	// IPv6 is unsupported.
	ElbSubnetIDs []string `json:"elb_virsubnet_ids,omitempty"`
	// Specifies whether to enable cross-VPC backend.
	//
	// If you enable this function, you can add servers in a peer VPC connected through a VPC peering connection, or in an on-premises data center at the other end of a Direct Connect or VPN connection, by using their IP addresses.
	//
	// This function is supported only by dedicated load balancers.
	//
	// The value can be true (enable cross-VPC backend) or false (disable cross-VPC backend).
	//
	// The value can only be update to true. This parameter is not available in eu-nl region. Please do not use it.
	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`
	// Specifies whether to enable deletion protection for the load balancer.
	//
	// true: Enable deletion protection.
	//
	// false (default): Disable deletion protection.
	//
	// Note
	//
	// Disable deletion protection for all your resources before deleting your account.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`
}

// Create is an operation which provisions a new loadbalancer based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*LoadBalancer, error) {
	b, err := build.RequestBody(opts, "loadbalancer")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/elb/loadbalancers
	raw, err := client.Post(client.ServiceURL("loadbalancers"), b, nil, nil)
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*LoadBalancer, error) {
	if err != nil {
		return nil, err
	}

	var res LoadBalancer
	err = extract.IntoStructPtr(raw.Body, &res, "loadbalancer")
	return &res, err
}
