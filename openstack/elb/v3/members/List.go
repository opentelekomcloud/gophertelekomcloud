package members

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections
// through the API. Filtering is achieved by passing in struct field values
// that map to the Member attributes you want to see returned. SortKey allows
// you to sort by a particular Member attribute. SortDir sets the direction,
// and is either `asc' or `desc'. Marker and Limit are used for pagination.
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
	// Specifies the backend server name.
	//
	// Multiple names can be queried in the format of name=xxx&name=xxx.
	Name []string `q:"name"`
	// Specifies the weight of the backend server. Requests are routed to backend servers in the same backend server group based on their weights.
	//
	// The value ranges from 0 to 100. The larger the weight is, the higher proportion of requests the backend server receives. If the weight is set to 0, the backend server will not accept new requests.
	//
	// Multiple weights can be queried in the format of weight=xxx&weight=xxx.
	Weight []string `q:"weight"`
	// Specifies the administrative status of the backend server. The value can be true or false.
	//
	// Although this parameter can be used in the APIs for creating and updating backend servers, its actual value depends on whether cloud servers exist. If cloud servers exist, the value is true. Otherwise, the value is false.
	AdminStateUp *bool `q:"admin_state_up"`
	// Specifies the ID of the IPv4 or IPv6 subnet where the backend server resides.
	//
	// Multiple IDs can be queried in the format of subnet_cidr_id=xxx&subnet_cidr_id=xxx.
	//
	// IPv6 is unsupported. Please do not set the value to the ID of an IPv6 subnet.
	SubnetCidrId []string `q:"subnet_cidr_id"`
	// Specifies the IP address bound to the backend server.
	//
	// Multiple IP addresses can be queried in the format of address=xxx&address=xxx.
	//
	// IPv6 is unsupported. Please do not set the value to an IPv6 address.
	Address []string `q:"address"`
	// Specifies the port used by the backend server to receive requests.
	//
	// Multiple ports can be queried in the format of protocol_port=xxx&protocol_port=xxx.
	ProtocolPort []string `q:"protocol_port"`
	// Specifies the backend server ID.
	//
	// Multiple IDs can be queried in the format of id=xxx&id=xxx.
	Id []string `q:"id"`
	// Specifies the health status of the backend server. The value can be one of the following:
	//
	// ONLINE: The backend server is running normally.
	//
	// NO_MONITOR: No health check is configured for the backend server group to which the backend server belongs.
	//
	// OFFLINE: The cloud server used as the backend server is stopped or does not exist.
	//
	// Multiple operating statuses can be queried in the format of operating_status=xxx&operating_status=xxx.
	OperatingStatus []string `q:"operating_status"`
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
	// Specifies the IP version supported by the backend server. The value can be v4 (IPv4) or v6 (IPv6).
	IpVersion []string `q:"ip_version"`
	// Specifies the type of the backend server. Values:
	//
	// ip: cross-VPC backend servers
	//
	// instance: ECSs used as backend servers Multiple values can be queried in the format of member_type=xxx&member_type=xxx.
	MemberType []string `q:"member_type"`
	// Specifies the ID of the instance associated with the backend server. If this parameter is left blank, the backend server is not an ECS. It may be an IP address.
	//
	// Multiple instance id can be queried in the format of instance_id=xxx&instance_id=xxx.
	InstanceId []string `q:"instance_id"`
}

// List returns a Pager which allows you to iterate over a collection of
// members. It accepts a ListOptsBuilder, which allows you to filter and
// sort the returned collection for greater efficiency.
//
// Default policy settings return only those members that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(client *golangsdk.ServiceClient, poolID string, opts ListOpts) pagination.Pager {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	// GET /v3/{project_id}/elb/pools/{pool_id}/members
	return pagination.NewPager(client, client.ServiceURL("pools", poolID, "members")+query.String(), func(r pagination.PageResult) pagination.Page {
		return MemberPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}

// MemberPage is the page returned by a pager when traversing over a
// collection of Members in a Pool.
type MemberPage struct {
	pagination.PageWithInfo
}

func (p MemberPage) IsEmpty() (bool, error) {
	l, err := ExtractMembers(p)
	return len(l) == 0, err
}

// ExtractMembers accepts a Page struct, specifically a MemberPage struct,
// and extracts the elements into a slice of Members structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractMembers(r pagination.Page) ([]Member, error) {
	var res []Member
	err := extract.IntoSlicePtr(r.(MemberPage).BodyReader(), &res, "members")
	return res, err
}
