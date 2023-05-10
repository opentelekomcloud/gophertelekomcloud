package loadbalancers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// Get retrieves a particular Loadbalancer based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (*LoadBalancer, error) {
	// GET /v3/{project_id}/elb/loadbalancers/{loadbalancer_id}
	raw, err := client.Get(client.ServiceURL("loadbalancers", id), nil, nil)
	return extra(err, raw)
}

// LoadBalancer is the primary load balancing configuration object that
// specifies the virtual IP address on which client traffic is received, as well
// as other details such as the load balancing method to be use, protocol, etc.
type LoadBalancer struct {
	// The unique ID for the LoadBalancer.
	ID string `json:"id"`
	// Provides supplementary information about the load balancer.
	//
	// Minimum: 1
	//
	// Maximum: 255
	Description string `json:"description"`
	// Specifies the provisioning status of the load balancer. The value can be one of the following:
	//
	// ACTIVE: The load balancer is successfully provisioned.
	//
	// PENDING_DELETE: The load balancer is being deleted.
	ProvisioningStatus string `json:"provisioning_status"`
	// Specifies the administrative status of the load balancer. The value can only be true.
	AdminStateUp bool `json:"admin_state_up"`
	// Specifies the provider of the load balancer. The value can only be vlb.
	Provider string `json:"provider"`
	// Lists the IDs of backend server groups associated with the load balancer.
	Pools []structs.ResourceRef `json:"pools"`
	// Lists the IDs of listeners added to the load balancer.
	Listeners []structs.ResourceRef `json:"listeners"`
	// Specifies the operating status of the load balancer. The value can only be ONLINE, indicating that the load balancer is running normally.
	OperatingStatus string `json:"operating_status"`
	// Specifies the private IPv4 address bound to the load balancer.
	VipAddress string `json:"vip_address"`
	// Specifies the ID of the IPv4 subnet where the load balancer resides.
	VipSubnetCidrID string `json:"vip_subnet_cidr_id"`
	// Specifies the load balancer name.
	Name string `json:"name"`
	// Owner of the LoadBalancer.
	ProjectID string `json:"project_id"`
	// Specifies the ID of the port bound to the private IPv4 address of the load balancer.
	//
	// The security group associated with the port will not take effect.
	VipPortID string `json:"vip_port_id"`
	// Lists the tags added to the load balancer.
	Tags []tags.ResourceTag `json:"tags"`
	// Specifies whether the load balancer is a dedicated load balancer.
	//
	// true (default): The load balancer is a dedicated load balancer.
	//
	// false: The load balancer is a shared load balancer.
	Guaranteed bool `json:"guaranteed"`
	// Specifies the ID of the VPC where the load balancer resides.
	VpcID string `json:"vpc_id"`
	// Specifies the EIP bound to the load balancer. Only one EIP can be bound to a load balancer.
	//
	// This parameter has the same meaning as publicips.
	Eips []EipInfo `json:"eips"`
	// Specifies the IPv6 address bound to the load balancer.
	//
	// IPv6 is unsupported. Please do not use this parameter.
	IpV6VipAddress string `json:"ipv6_vip_address"`
	// Specifies the ID of the IPv6 subnet where the load balancer resides.
	//
	// IPv6 is unsupported. Please do not use this parameter.
	IpV6VipSubnetID string `json:"ipv6_vip_virsubnet_id"`
	// Specifies the ID of the port bound to the IPv6 address of the load balancer.
	//
	// IPv6 is unsupported. Please do not use this parameter.
	IpV6VipPortID string `json:"ipv6_vip_port_id"`
	// Specifies the list of AZs where the load balancer is created.
	AvailabilityZoneList []string `json:"availability_zone_list"`
	// Specifies the ID of a flavor at Layer 4.
	//
	// Minimum: 1
	//
	// Maximum: 255
	L4FlavorID string `json:"l4_flavor_id"`
	// Specifies the ID of the reserved flavor at Layer 4.
	//
	// This parameter is unsupported. Please do not use it.
	//
	// Minimum: 1
	//
	// Maximum: 255
	L4ScaleFlavorID string `json:"l4_scale_flavor_id"`
	// Specifies the ID of a flavor at Layer 7.
	//
	// Minimum: 1
	//
	// Maximum: 255
	L7FlavorID string `json:"l7_flavor_id"`
	// Specifies the ID of the reserved flavor at Layer 7.
	//
	// This parameter is unsupported. Please do not use it.
	//
	// Minimum: 1
	//
	// Maximum: 255
	L7ScaleFlavorID string `json:"l7_scale_flavor_id"`
	// Specifies the EIP bound to the load balancer. Only one EIP can be bound to a load balancer.
	//
	// This parameter has the same meaning as eips.
	PublicIps []PublicIpInfo `json:"publicips"`
	// Lists the IDs of subnets on the downstream plane.
	ElbSubnetIDs []string `json:"elb_virsubnet_ids"`
	// Specifies the type of the subnet on the downstream plane.
	//
	// ipv4: IPv4 subnet
	//
	// dualstack: subnet that supports IPv4/IPv6 dual stack
	//
	// "dualstack" is not supported.
	ElbSubnetType string `json:"elb_virsubnet_type"`
	// Specifies whether to enable cross-VPC backend.
	//
	// If you enable this function, you can add servers in a peer VPC connected through a VPC peering connection, or in an on-premises data center at the other end of a Direct Connect or VPN connection, by using their IP addresses.
	//
	// This function is supported only by dedicated load balancers.
	//
	// The value can be true (enable cross-VPC backend) or false (disable cross-VPC backend).
	//
	// The value can only be update to true. This parameter is not available in eu-nl region. Please do not use it.
	IpTargetEnable bool `json:"ip_target_enable"`
	// Specifies the scenario where the load balancer is frozen. Multiple values are separated using commas (,).
	//
	// POLICE: The load balancer is frozen due to security reasons.
	//
	// ILLEGAL: The load balancer is frozen due to violation of laws and regulations.
	//
	// VERIFY: Your account has not completed real-name authentication.
	//
	// PARTNER: The load balancer is frozen by the partner.
	//
	// ARREAR: Your account is in arrears.
	//
	// This parameter is unsupported. Please do not use it.
	FrozenScene string `json:"frozen_scene"`
	// Specifies the ID of the bandwidth used by an IPv6 address. This parameter is available only when you create or update a load balancer with a public IPv6 address. If you use a new IPv6 address and specify a shared bandwidth, the IPv6 address will be added to the shared bandwidth.
	//
	// IPv6 is unsupported. Please do not use this parameter.
	IpV6Bandwidth BandwidthRef `json:"ipv6_bandwidth"`
	// Specifies the time when the load balancer was created, in the format of yyyy-MM-dd''T''HH:mm:ss''Z''.
	CreatedAt string `json:"created_at"`
	// Specifies the time when the load balancer was updated, in the format of yyyy-MM-dd''T''HH:mm:ss''Z''.
	UpdatedAt string `json:"updated_at"`
	// Specifies whether deletion protection is enabled.
	//
	// false: Deletion protection is not enabled.
	//
	// true: Deletion protection is enabled.
	//
	// Note
	//
	// Disable deletion protection for all your resources before deleting your account.
	//
	// This parameter is returned only when deletion protection is enabled at the site.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	DeletionProtectionEnable bool `json:"deletion_protection_enable"`
	// Specifies the AZ group to which the load balancer belongs.
	PublicBorderGroup string `json:"public_border_group"`
}
