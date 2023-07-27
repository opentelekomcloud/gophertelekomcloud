package loadbalancers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts is the common options' struct used in this package's Update
// operation.
type UpdateOpts struct {
	// Specifies the load balancer name.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Name string `json:"name,omitempty"`
	// Specifies the administrative status of the load balancer. The value can only be true.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// Provides supplementary information about the load balancer.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Description string `json:"description,omitempty"`
	// Specifies the ID of the IPv6 subnet where the load balancer resides. You can query parameter id in the response by calling the API (GET https://{VPC_Endpoint}/v1/{project_id}/subnets).
	//
	// The IPv6 subnet can be updated using ipv6_vip_virsubnet_id, and the private IPv6 address of the load balancer will be changed accordingly.
	//
	// Note:
	//
	// This parameter will be passed only when IPv6 is enabled for the subnet. The subnet specified by ipv6_vip_virsubnet_id must be in the VPC specified by vpc_id.
	//
	// This parameter can be updated only when guaranteed is set to true.
	//
	// The value will become null if the IPv6 address is unbound from the load balancer.
	//
	// The IPv4 subnet will not change, if IPv6 subet is updated. This parameter is unsupported. Please do not use it.
	Ipv6VipVirsubnetId string `json:"ipv6_vip_virsubnet_id,omitempty"`
	// Specifies the ID of the IPv4 subnet where the load balancer resides. You can query parameter neutron_subnet_id in the response by calling the API (GET https://{VPC_Endpoint}/v1/{project_id}/subnets).
	//
	// The IPv4 subnet can be updated using vip_subnet_cidr_id, and the private IPv4 address of the load balancer will be changed accordingly. Note:
	//
	// If vip_address is also specified, the IP address specified by vip_address must be in the subnet specified by vip_subnet_cidr_id and will be used as the private IPv4 address of the load balancer.
	//
	// The IPv4 subnet must be in the VPC where the load balancer resides.
	//
	// This parameter can be updated only when guaranteed is set to true.
	//
	// The value will become null if the private IPv4 address is unbound from the load balancer.
	//
	// The IPv6 subnet will not change, if IPv4 subet is updated.
	//
	// Minimum: 1
	//
	// Maximum: 36
	VipSubnetCidrId string `json:"vip_subnet_cidr_id,omitempty"`
	// Specifies the private IPv4 address bound to the load balancer. The IP address must be from the IPv4 subnet where the load balancer resides and should not be occupied by other services.
	//
	// vip_address can be updated only when guaranteed is set to true.
	//
	// Minimum: 1
	//
	// Maximum: 36
	VipAddress string `json:"vip_address,omitempty"`
	// Specifies the ID of a flavor at Layer 4.
	//
	// Note:
	//
	// This parameter can be updated only when guaranteed is set to true.
	//
	// The value cannot be changed from null to a specific value, or in the other way around.
	//
	// If you change the flavor, you can select only a higher or lower one. If you select a lower one, part of persistent connections will be interrupted.
	//
	// If autoscaling.enable is set to true, updating this parameter will not take effect.
	//
	// Minimum: 1
	//
	// Maximum: 255
	L4FlavorId string `json:"l4_flavor_id,omitempty"`
	// Specifies the ID of a flavor at Layer 7.
	//
	// Note:
	//
	// This parameter can be updated only when guaranteed is set to true.
	//
	// The value cannot be changed from null to a specific value, or in the other way around.
	//
	// If you change the flavor, you can select only a higher or lower one. If you select a lower one, part of persistent connections will be interrupted.
	//
	// If autoscaling.enable is set to true, updating this parameter will not take effect.
	//
	// Minimum: 1
	//
	// Maximum: 36
	L7FlavorId string `json:"l7_flavor_id,omitempty"`
	// Specifies the ID of the bandwidth used by an IPv6 address. This parameter is available only when you create or update a load balancer with a public IPv6 address. If you use a new IPv6 address and specify a shared bandwidth, the IPv6 address will be added to the shared bandwidth.
	//
	// IPv6 is unsupported. Please do not use this parameter.
	Ipv6Bandwidth BandwidthRef `json:"ipv6_bandwidth,omitempty"`
	// Specifies whether to enable cross-VPC backend.
	//
	// If you enable this function, you can add servers in a peer VPC connected through a VPC peering connection, or in an on-premises data center at the other end of a Direct Connect or VPN connection, by using their IP addresses.
	//
	// This function is supported only by dedicated load balancers.
	//
	// The value can be true (enable cross-VPC backend) or false (disable cross-VPC backend).
	//
	// The value can only be update to true.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`
	// Specifies the IDs of subnets on the downstream plane. You can query parameter neutron_network_id in the response by calling the API (GET https://{VPC_Endpoint}/v1/{project_id}/subnets).
	//
	// If the IDs of the subnets required by the load balancer are specified in elb_virsubnet_ids, the subnets will still be bound to the load balancer.
	//
	// If the IDs of the subnets required by the load balancer are not specified in elb_virsubnet_ids, the subnets will be unbound from the load balancers. Do not unbound the subnets that have been used by the load balancer. Otherwise, an error will be returned.
	//
	// If the IDs of the subnets are specified in elb_virsubnet_ids, but not on the downstream plane, a new load balancer will be bound to the downstream plane.
	//
	// Note:
	//
	// All subnets belong to the same VPC where the load balancer resides.
	//
	// Edge subnets are not supported.
	//
	// Minimum: 1
	//
	// Maximum: 64
	ElbVirSubnetIds []string `json:"elb_virsubnet_ids,omitempty"`
	// Specifies whether to enable deletion protection for the load balancer.
	//
	// true: Enable deletion protection.
	//
	// false: Disable deletion protection.
	//
	// Disable deletion protection for all your resources before deleting your account.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`
}

// Update is an operation which modifies the attributes of the specified LoadBalancer.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*LoadBalancer, error) {
	b, err := build.RequestBody(opts, "loadbalancer")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/elb/loadbalancers/{loadbalancer_id}
	raw, err := client.Put(client.ServiceURL("loadbalancers", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return extra(err, raw)
}
