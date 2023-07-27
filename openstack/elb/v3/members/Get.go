package members

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
)

// Get retrieves a particular Pool Member based on its unique ID.
func Get(client *golangsdk.ServiceClient, poolID string, memberID string) (*Member, error) {
	// GET /v3/{project_id}/elb/pools/{pool_id}/members/{member_id}
	raw, err := client.Get(client.ServiceURL("pools", poolID, "members", memberID), nil, nil)
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*Member, error) {
	if err != nil {
		return nil, err
	}

	var res Member
	err = extract.IntoStructPtr(raw.Body, &res, "member")
	return &res, err
}

// Member represents the application running on a backend server.
type Member struct {
	// Specifies the backend server ID.
	//
	// Note:
	//
	// The value of this parameter is not the ID of the server but an ID automatically generated for the backend server that has already associated with the load balancer.
	Id string `json:"id"`
	// Specifies the backend server name.
	Name string `json:"name"`
	// Specifies the project ID of the backend server.
	ProjectId string `json:"project_id"`
	// Specifies the ID of the backend server group to which the backend server belongs.
	//
	// This parameter is unsupported. Please do not use it.
	PoolId string `json:"pool_id"`
	// Specifies the administrative status of the backend server. The value can be true or false.
	//
	// Although this parameter can be used in the APIs for creating and updating backend servers, its actual value depends on whether cloud servers exist. If cloud servers exist, the value is true. Otherwise, the value is false.
	AdminStateUp *bool `json:"admin_state_up"`
	// Specifies the ID of the IPv4 or IPv6 subnet where the backend server resides.
	//
	// This parameter can be left blank, indicating that cross-VPC backend has been enabled for the load balancer. In this case, IP addresses of these servers must be IPv4 addresses, and the protocol of the backend server group must be TCP, HTTP, or HTTPS.
	//
	// The IPv4 or IPv6 subnet must be in the same VPC as the subnet of the load balancer.
	//
	// IPv6 is unsupported. Please do not set the value to the ID of an IPv6 subnet.
	SubnetCidrId string `json:"subnet_cidr_id"`
	// Specifies the port used by the backend server to receive requests.
	//
	// Minimum: 1
	//
	// Maximum: 65535
	ProtocolPort *int `json:"protocol_port"`
	// Specifies the weight of the backend server. Requests are routed to backend servers in the same backend server group based on their weights.
	//
	// The value ranges from 0 to 100, and the default value is 1. The larger the weight is, the higher proportion of requests the backend server receives. If the weight is set to 0, the backend server will not accept new requests.
	//
	// If lb_algorithm is set to SOURCE_IP, this parameter will not take effect.
	//
	// Minimum: 0
	//
	// Maximum: 100
	Weight *int `json:"weight"`
	// Specifies the private IP address bound to the backend server.
	//
	// If subnet_cidr_id is left blank, cross-VPC backend is enabled. In this case, the IP address must be an IPv4 address.
	//
	// If subnet_cidr_id is not left blank, the IP address can be IPv4 or IPv6. It must be in the subnet specified by subnet_cidr_id and can only be bound to the primary NIC of the backend server.
	//
	// IPv6 is unsupported. Please do not enter an IPv6 address.
	Address string `json:"address"`
	// Specifies the IP version supported by the backend server. The value can be v4 (IPv4) or v6 (IPv6), depending on the value of address returned by the system.
	IpVersion string `json:"ip_version"`
	// Specifies the health status of the backend server if listener_id under status is not specified. The value can be one of the following:
	//
	// ONLINE: The backend server is running normally.
	//
	// NO_MONITOR: No health check is configured for the backend server group to which the backend server belongs.
	//
	// OFFLINE: The cloud server used as the backend server is stopped or does not exist.
	OperatingStatus string `json:"operating_status"`
	// Specifies the health status of the backend server if listener_id is specified.
	Status []MemberStatus `json:"status"`
	// Specifies the ID of the load balancer with which the backend server is associated.
	//
	// This parameter is unsupported. Please do not use it.
	LoadBalancerId string `json:"loadbalancer_id"`
	// Specifies the IDs of the load balancers associated with the backend server.
	//
	// This parameter is unsupported. Please do not use it.
	LoadBalancers []structs.ResourceRef `json:"loadbalancers"`
	// Specifies the time when a backend server was added. The format is yyyy-MM-dd'T'HH:mm:ss'Z' (UTC time).
	//
	// This is a new field in this version, and it will not be returned for resources associated with existing dedicated load balancers and for resources associated with existing and new shared load balancers.
	CreatedAt string `json:"created_at"`
	// Specifies the time when a backend server was updated. The format is yyyy-MM-dd'T'HH:mm:ss'Z' (UTC time).
	//
	// This is a new field in this version, and it will not be returned for resources associated with existing dedicated load balancers and for resources associated with existing and new shared load balancers.
	UpdatedAt string `json:"updated_at"`
	// Specifies the type of the backend server. Values:
	//
	// ip: cross-VPC backend servers
	//
	// instance: ECSs used as backend servers
	MemberType string `json:"member_type"`
	// Specifies the ID of the ECS used as the backend server. If this parameter is left blank, the backend server is not an ECS. For example, it may be an IP address.
	InstanceId string `json:"instance_id"`
}

type MemberStatus struct {
	// Specifies the listener ID.
	ListenerId string `json:"listener_id"`
	// Specifies the health status of the backend server. The value can be one of the following:
	//
	// ONLINE: The backend server is running normally.
	//
	// NO_MONITOR: No health check is configured for the backend server group to which the backend server belongs.
	//
	// OFFLINE: The cloud server used as the backend server is stopped or does not exist.
	OperatingStatus string `json:"operating_status"`
}
