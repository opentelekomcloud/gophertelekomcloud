package loadbalancers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Specifies the ID of the last record on the previous page.
	//
	// Note:
	//
	// This parameter must be used together with limit.
	//
	// If this parameter is not specified, the first page will be queried.
	//
	// This parameter cannot be left blank or set to an invalid ID.
	Marker string `q:"marker"`
	// Specifies the number of records on each page.
	//
	// Minimum: 0
	//
	// Maximum: 2000
	//
	// Default: 2000
	Limit int `q:"limit"`
	// Specifies whether to use reverse query. Values:
	//
	// true: Query the previous page.
	//
	// false (default): Query the next page.
	//
	// Note:
	//
	// This parameter must be used together with limit.
	//
	// If page_reverse is set to true and you want to query the previous page, set the value of marker to the value of previous_marker.
	PageReverse *bool `q:"page_reverse"`
	// Specifies the load balancer ID.
	//
	// Multiple IDs can be queried in the format of id=xxx&id=xxx.
	Id []string `q:"id"`
	// Specifies the load balancer name.
	//
	// Multiple names can be queried in the format of name=xxx&name=xxx.
	Name []string `q:"name"`
	// Provides supplementary information about the load balancer.
	//
	// Multiple descriptions can be queried in the format of description=xxx&description=xxx.
	Description []string `q:"description"`
	// Specifies the administrative status of the load balancer.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `q:"admin_state_up"`
	// Specifies the provisioning status of the load balancer.
	//
	// ACTIVE: The load balancer is successfully provisioned.
	//
	// PENDING_DELETE: The load balancer is being deleted.
	//
	// Multiple provisioning statuses can be queried in the format of provisioning_status=xxx&provisioning_status=xxx.
	ProvisioningStatus []string `q:"provisioning_status"`
	// Specifies the operating status of the load balancer.
	//
	// ONLINE: The load balancer is working normally.
	//
	// FROZEN: The load balancer has been frozen.
	//
	// Multiple operating statuses can be queried in the format of operating_status=xxx&operating_status=xxx.
	OperatingStatus []string `q:"operating_status"`
	// Specifies whether the load balancer is a dedicated load balancer.
	//
	// false: The load balancer is a shared load balancer.
	//
	// true: The load balancer is a dedicated load balancer.
	Guaranteed *bool `q:"guaranteed"`
	// Specifies the ID of the VPC where the load balancer resides.
	//
	// Multiple IDs can be queried in the format of vpc_id=xxx&vpc_id=xxx.
	VpcId []string `q:"vpc_id"`
	// Specifies the ID of the port bound to the private IPv4 address of the load balancer.
	//
	// Multiple IDs can be queried in the format of vip_port_id=xxx&vip_port_id=xxx.
	VipPortId []string `q:"vip_port_id"`
	// Specifies the virtual IP address bound to the load balancer.
	//
	// Multiple virtual IP addresses can be queried in the format of vip_address=xxx&vip_address=xxx.
	VipAddress []string `q:"vip_address"`
	// Specifies the ID of the IPv4 subnet where the load balancer resides.
	//
	// Multiple IDs can be queried in the format of vip_subnet_cidr_id=xxx&vip_subnet_cidr_id=xxx.
	VipSubnetCidrId []string `q:"vip_subnet_cidr_id"`
	// Specifies the ID of the port bound to the IPv6 address of the load balancer.
	//
	// Multiple ports can be queried in the format of ipv6_vip_port_id=xxx&ipv6_vip_port_id=xxx.
	//
	// IPv6 is unsupported. Please do not use this parameter.
	Ipv6VipPortId []string `q:"ipv6_vip_port_id"`
	// Specifies the IPv6 address bound to the load balancer.
	//
	// Multiple IPv6 addresses can be queried in the format of ipv6_vip_address=xxx&ipv6_vip_address=xxx.
	//
	// IPv6 is unsupported. Please do not use this parameter.
	Ipv6VipAddress []string `q:"ipv6_vip_address"`
	// Specifies the ID of the IPv6 subnet where the load balancer resides.
	//
	// Multiple IDs can be queried in the format of ipv6_vip_virsubnet_id=xxx&ipv6_vip_virsubnet_id=xxx.
	//
	// IPv6 is unsupported. Please do not use this parameter.
	Ipv6VipVirSubnetId []string `q:"ipv6_vip_virsubnet_id"`
	// Specifies the IPv4 EIP bound to the load balancer. The following is an example: "eips": [ { "eip_id": "e9b72a9d-4275-455e-a724-853504e4d9c6", "eip_address": "88.88.14.122", "ip_version": 4 } ]
	//
	// Multiple EIPs can be queried.
	//
	// If eip_id is used as the query condition, the format is eips=eip_id=xxx&eips=eip_id=xxx.
	//
	// If eip_address is used as the query condition, the format is eips=eip_address=xxx&eips=eip_address=xxx.
	//
	// If ip_version is used as the query condition, the format is eips=ip_version=xxx&eips=ip_version=xxx.
	//
	// Note that this parameter has the same meaning as publicips.
	Eips []string `q:"eips"`
	// Specifies the IPv4 EIP bound to the load balancer. The following is an example: "publicips": [ { "publicip_id": "e9b72a9d-4275-455e-a724-853504e4d9c6", "publicip_address": "88.88.14.122", "ip_version": 4 } ]
	//
	// Multiple EIPs can be queried.
	//
	// If publicip_id is used as the query condition, the format is publicips=publicip_id=xxx&publicips=publicip_id=xxx.
	//
	// If publicip_address is used as the query condition, the format is *publicips=publicip_address=xxx&publicips=publicip_address=xxx.
	//
	// If publicip_address is used as the query condition, the format is publicips=ip_version=xxx&publicips=ip_version=xxx.
	//
	// Note that this parameter has the same meaning as eips.
	PublicIps []string `q:"publicips"`
	// Specifies the list of AZs where the load balancer is created.
	//
	// Multiple AZs can be queried in the format of availability_zone_list=xxx&availability_zone_list=xxx.
	AvailabilityZoneList []string `q:"availability_zone_list"`
	// Specifies the ID of a flavor at Layer 4.
	//
	// Multiple IDs can be queried in the format of l4_flavor_id=xxx&l4_flavor_id=xxx.
	L4FlavorId []string `q:"l4_flavor_id"`
	// Specifies the ID of the elastic flavor at Layer 4, which is reserved for now.
	//
	// Multiple flavors can be queried in the format of l4_scale_flavor_id=xxx&l4_scale_flavor_id=xxx.
	//
	// This parameter is unsupported. Please do not use it.
	L4ScaleFlavorId []string `q:"l4_scale_flavor_id"`
	// Specifies the ID of a flavor at Layer 7.
	//
	// Multiple flavors can be queried in the format of l7_flavor_id=xxx&l7_flavor_id=xxx.
	L7FlavorId []string `q:"l7_flavor_id"`
	// Specifies the ID of the elastic flavor at Layer 7. Multiple flavors can be queried in the format of l7_scale_flavor_id=xxx&l7_scale_flavor_id=xxx. This parameter is unsupported. Please do not use it.
	L7ScaleFlavorId []string `q:"l7_scale_flavor_id"`
	// Provides resource billing information.
	//
	// Multiple values can be queried in the format of billing_info=xxx&billing_info=xxx.
	//
	// This parameter is unsupported. Please do not use it.
	BillingInfo []string `q:"billing_info"`
	// Specifies the ID of the cloud server that is associated with the load balancer as a backend server. This is a query parameter and will not be included in the response.
	//
	// Multiple IDs can be queried in the format of member_device_id=xxx&member_device_id=xxx.
	MemberDeviceId []string `q:"member_device_id"`
	// Specifies the private IP address of the cloud server that is associated with the load balancer as a backend server. This is a query parameter and will not be included in the response.
	//
	// Multiple private IP addresses can be queried in the format of member_address=xxx&member_address=xxx.
	MemberAddress []string `q:"member_address"`
	// Specifies the enterprise project ID.
	//
	// If this parameter is not passed, resources in the default enterprise project are queried, and authentication is performed based on the default enterprise project.
	//
	// If this parameter is passed, its value can be the ID of an existing enterprise project (resources in the specific enterprise project are required) or all_granted_eps (resources in all enterprise projects are queried).
	//
	// Multiple IDs can be queried in the format of enterprise_project_id=xxx&enterprise_project_id=xxx.
	//
	// This parameter is unsupported. Please do not use it.
	EnterpriseProjectId []string `q:"enterprise_project_id"`
	// Specifies the IP version. The value can be 4 (IPv4) or 6 (IPv6).
	//
	// Multiple versions can be queried in the format of ip_version=xxx&ip_version=xxx.
	//
	// IPv6 is unsupported. The value cannot be 6.
	IpVersion []string `q:"ip_version"`
	// Specifies whether to enable deletion protection.
	//
	// true: Enable deletion protection.
	//
	// false: Disable deletion protection.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	DeletionProtectionEnable *bool `q:"deletion_protection_enable"`
	// Specifies the type of the subnet on the downstream plane.
	//
	// ipv4: IPv4 subnet
	//
	// dualstack: subnet that supports IPv4/IPv6 dual stack
	//
	// Multiple values query can be queried in the format of elb_virsubnet_type=ipv4&elb_virsubnet_type=dualstack.
	//
	// "dualstack" is not supported.
	ElbVirSubnetType []string `q:"elb_virsubnet_type"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	// GET /v3/{project_id}/elb/loadbalancers
	return pagination.NewPager(client, client.ServiceURL("loadbalancers")+query.String(), func(r pagination.PageResult) pagination.Page {
		return LoadBalancerPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}

// LoadBalancerPage is the page returned by a pager when traversing over a
// collection of loadbalancer.
type LoadBalancerPage struct {
	pagination.PageWithInfo
}

// IsEmpty checks whether a FlavorsPage struct is empty.
func (r LoadBalancerPage) IsEmpty() (bool, error) {
	is, err := ExtractLoadBalancers(r)
	return len(is) == 0, err
}

// ExtractLoadBalancers accepts a Page struct, specifically a LoadBalancerPage struct,
// and extracts the elements into a slice of loadbalancer structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractLoadBalancers(r pagination.Page) ([]LoadBalancer, error) {
	var res []LoadBalancer
	err := extract.IntoSlicePtr(r.(LoadBalancerPage).BodyReader(), &res, "loadbalancers")
	return res, err
}
