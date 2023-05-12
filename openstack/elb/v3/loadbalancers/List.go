package loadbalancers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Specifies the load balancer ID.
	//
	// Multiple IDs can be queried in the format of id=xxx&id=xxx.
	ID []string `q:"id"`
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
	VpcID []string `q:"vpc_id"`
	// Specifies the ID of the port bound to the private IPv4 address of the load balancer.
	//
	// Multiple IDs can be queried in the format of vip_port_id=xxx&vip_port_id=xxx.
	VipPortID []string `q:"vip_port_id"`
	// Specifies the virtual IP address bound to the load balancer.
	//
	// Multiple virtual IP addresses can be queried in the format of vip_address=xxx&vip_address=xxx.
	VipAddress []string `q:"vip_address"`
	// Specifies the ID of the IPv4 subnet where the load balancer resides.
	//
	// Multiple IDs can be queried in the format of vip_subnet_cidr_id=xxx&vip_subnet_cidr_id=xxx.
	VipSubnetCidrID []string `q:"vip_subnet_cidr_id"`
	// Specifies the ID of a flavor at Layer 4.
	//
	// Multiple IDs can be queried in the format of l4_flavor_id=xxx&l4_flavor_id=xxx.
	L4FlavorID []string `q:"l4_flavor_id"`
	// Specifies the ID of the elastic flavor at Layer 4, which is reserved for now.
	//
	// Multiple flavors can be queried in the format of l4_scale_flavor_id=xxx&l4_scale_flavor_id=xxx.
	//
	// This parameter is unsupported. Please do not use it.
	L4ScaleFlavorID []string `q:"l4_scale_flavor_id"`
	// Specifies the list of AZs where the load balancer is created.
	//
	// Multiple AZs can be queried in the format of availability_zone_list=xxx&availability_zone_list=xxx.
	AvailabilityZoneList []string `q:"availability_zone_list"`
	// Specifies the ID of a flavor at Layer 7.
	//
	// Multiple flavors can be queried in the format of l7_flavor_id=xxx&l7_flavor_id=xxx.
	L7FlavorID []string `q:"l7_flavor_id"`
	// Specifies the ID of the elastic flavor at Layer 7. Multiple flavors can be queried in the format of l7_scale_flavor_id=xxx&l7_scale_flavor_id=xxx. This parameter is unsupported. Please do not use it.
	L7ScaleFlavorID []string `q:"l7_scale_flavor_id"`
	// Specifies the number of records on each page.
	//
	// Minimum: 0
	//
	// Maximum: 2000
	//
	// Default: 2000
	Limit int `q:"limit"`
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
	//
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
}

// ToLoadbalancerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToLoadbalancerListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("loadbalancers")
	if opts != nil {
		query, err := opts.ToLoadbalancerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return LoadbalancerPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}
