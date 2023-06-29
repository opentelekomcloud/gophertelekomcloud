package pools

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Pool attributes you want to see returned. SortKey allows you to
// sort by a particular Pool attribute. SortDir sets the direction, and is
// either `asc` or `desc`. Marker and Limit are used for pagination.
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
	Limit *int `q:"limit"`
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
	// Provides supplementary information about the backend server group.
	//
	// Multiple descriptions can be queried in the format of description=xxx&description=xxx.
	Description []string `q:"description"`
	// Specifies the administrative status of the backend server group.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `q:"admin_state_up"`
	// Specifies the ID of the health check configured for the backend server group.
	//
	// Multiple IDs can be queried in the format of healthmonitor_id=xxx&healthmonitor_id=xxx.
	HealthmonitorId []string `q:"healthmonitor_id"`
	// Specifies the ID of the backend server group.
	//
	// Multiple IDs can be queried in the format of id=xxx&id=xxx.
	Id []string `q:"id"`
	// Specifies the backend server group name.
	//
	// Multiple names can be queried in the format of name=xxx&name=xxx.
	Name []string `q:"name"`
	// Specifies the ID of the load balancer with which the backend server group is associated.
	//
	// Multiple IDs can be queried in the format of loadbalancer_id=xxx&loadbalancer_id=xxx.
	LoadbalancerId []string `q:"loadbalancer_id"`
	// Specifies the protocol used by the backend server group to receive requests from the load balancer. The value can be TCP, UDP, HTTP, HTTPS, or QUIC.
	//
	// Multiple protocols can be queried in the format of protocol=xxx&protocol=xxx.
	//
	// QUIC protocol is not supported in eu-nl region.
	Protocol []string `q:"protocol"`
	// Specifies the load balancing algorithm used by the load balancer to route requests to backend servers in the associated backend server group.
	//
	// The value can be one of the following:
	//
	// ROUND_ROBIN: weighted round robin
	//
	// LEAST_CONNECTIONS: weighted least connections
	//
	// SOURCE_IP: source IP hash
	//
	// QUIC_CID: connection ID
	//
	// Multiple algorithms can be queried in the format of lb_algorithm=xxx&lb_algorithm=xxx.
	//
	// QUIC_CID is not supported in eu-nl region.
	LbAlgorithm []string `q:"lb_algorithm"`
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
	// Specifies the IP address version supported by the backend server group.
	//
	// Multiple versions can be queried in the format of ip_version=xxx&ip_version=xxx.
	IpVersion []string `q:"ip_version"`
	// Specifies the private IP address bound to the backend server. This is a query parameter and will not be included in the response.
	//
	// Multiple IP addresses can be queried in the format of member_address=xxx&member_address=xxx.
	MemberAddress []string `q:"member_address"`
	// Specifies the ID of the cloud server that serves as a backend server. This parameter is used only as a query condition and is not included in the response.
	//
	// Multiple IDs can be queried in the format of member_device_id=xxx&member_device_id=xxx.
	MemberDeviceId []string `q:"member_device_id"`
	// Specifies whether to enable removal protection on backend servers.
	//
	// true: Enable removal protection.
	//
	// false: Disable removal protection.
	//
	// All backend servers will be queried if this parameter is not passed.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	MemberDeletionProtectionEnable *bool `q:"member_deletion_protection_enable"`
	// Specifies the IDs of the associated listeners, including the listeners associated through forwarding policies.
	//
	// Multiple IDs can be queried in the format of listener_id=xxx&listener_id=xxx.
	ListenerId []string `q:"listener_id"`
	// Specifies the backend server ID. This parameter is used only as a query condition and is not included in the response. Multiple IDs can be queried in the format of member_instance_id=xxx&member_instance_id=xxx.
	MemberInstanceId []string `q:"member_instance_id"`
	// Specifies the ID of the VPC where the backend server group works.
	VpcId []string `q:"vpc_id"`
	// Specifies the type of the backend server group.
	//
	// Values:
	//
	// instance: Any type of backend servers can be added. vpc_id is mandatory.
	//
	// ip: Only IP as Backend servers can be added. vpc_id cannot be specified.
	//
	// "": Any type of backend servers can be added.
	Type []string `q:"type"`
}

// List returns a Pager which allows you to iterate over a collection of
// pools. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those pools that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("pools")+query.String(), func(r pagination.PageResult) pagination.Page {
		return PoolPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}

// PoolPage is the page returned by a pager when traversing over a
// collection of pools.
type PoolPage struct {
	pagination.PageWithInfo
}

// IsEmpty checks whether a PoolPage struct is empty.
func (r PoolPage) IsEmpty() (bool, error) {
	is, err := ExtractPools(r)
	return len(is) == 0, err
}

// ExtractPools accepts a Page struct, specifically a PoolPage struct,
// and extracts the elements into a slice of Pool structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractPools(r pagination.Page) ([]Pool, error) {
	var res []Pool
	err := extract.IntoSlicePtr(r.(PoolPage).BodyReader(), &res, "pools")
	return res, err
}
