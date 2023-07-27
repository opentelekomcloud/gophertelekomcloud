package members

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateOpts is the common options' struct used in this package's CreateMember operation.
type CreateOpts struct {
	// Specifies the private IP address bound to the backend server.
	//
	// If subnet_cidr_id is left blank, cross-VPC backend is enabled. In this case, the IP address must be an IPv4 address.
	//
	// If subnet_cidr_id is not left blank, the IP address can be IPv4 or IPv6. It must be in the subnet specified by subnet_cidr_id and can only be bound to the primary NIC of the backend server.
	//
	// IPv6 is unsupported. Please do not enter an IPv6 address.
	//
	// Minimum: 1
	//
	// Maximum: 64
	Address string `json:"address" required:"true"`
	// Specifies the administrative status of the backend server. The value can be true or false.
	//
	// Although this parameter can be used in the APIs for creating and updating backend servers, its actual value depends on whether cloud servers exist. If cloud servers exist, the value is true. Otherwise, the value is false.
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// Specifies the backend server name.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Name string `json:"name,omitempty"`
	// Specifies the project ID.
	//
	// Minimum: 1
	//
	// Maximum: 32
	ProjectId string `json:"project_id,omitempty"`
	// Specifies the port used by the backend server to receive requests.
	//
	// Minimum: 1
	//
	// Maximum: 65535
	ProtocolPort *int `json:"protocol_port" required:"true"`
	// Specifies the ID of the IPv4 or IPv6 subnet where the backend server resides.
	//
	// Note:
	//
	// The IPv4 or IPv6 subnet must be in the same VPC as the subnet of the load balancer.
	//
	// If this parameter is not passed, cross-VPC backend has been enabled for the load balancer. In this case, cross-VPC backend servers must use private IPv4 addresses, and the protocol of the backend server group must be TCP, HTTP, or HTTPS.
	//
	// IPv6 is unsupported. Please do not set the value to the ID of an IPv6 subnet.
	//
	// Minimum: 1
	//
	// Maximum: 36
	SubnetCidrId string `json:"subnet_cidr_id,omitempty"`
	// Specifies the weight of the backend server. Requests are routed to backend servers in the same backend server group based on their weights.
	//
	// The value ranges from 0 to 100, and the default value is 1. The larger the weight is, the higher proportion of requests the backend server receives. If the weight is set to 0, the backend server will not accept new requests.
	//
	// If lb_algorithm is set to SOURCE_IP, this parameter will not take effect.
	Weight *int `json:"weight,omitempty"`
}

// Create will create and associate a Member with a particular Pool.
func Create(client *golangsdk.ServiceClient, poolID string, opts CreateOpts) (*Member, error) {
	b, err := build.RequestBody(opts, "member")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/elb/pools/{pool_id}/members
	raw, err := client.Post(client.ServiceURL("pools", poolID, "members"), b, nil, nil)
	return extra(err, raw)
}
