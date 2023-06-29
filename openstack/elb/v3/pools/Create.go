package pools

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateOpts is the common options' struct used in this package's Create operation.
type CreateOpts struct {
	// Specifies the administrative status of the backend server group. The value can only be updated to true.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// Provides supplementary information about the backend server group.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Description string `json:"description,omitempty"`
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
	// Note:
	//
	// If the value is SOURCE_IP, the weight parameter will not take effect for backend servers.
	//
	// QUIC_CID is supported only when the protocol of the backend server group is QUIC.
	//
	// QUIC_CID is not supported in eu-nl region.
	LbAlgorithm string `json:"lb_algorithm" required:"true"`
	// Specifies the ID of the listener with which the backend server group is associated. Specify either listener_id or loadbalancer_id, or both of them.
	//
	// Specify either listener_id or loadbalancer_id for shared loadbalancer.
	//
	// Minimum: 1
	//
	// Maximum: 36
	ListenerId string `json:"listener_id,omitempty"`
	// Specifies the ID of the load balancer with which the backend server group is associated. Specify either listener_id or loadbalancer_id, or both of them.
	//
	// Minimum: 1
	//
	// Maximum: 36
	LoadbalancerId string `json:"loadbalancer_id,omitempty"`
	// Specifies the backend server group name.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Name string `json:"name,omitempty"`
	// Specifies the project ID of the backend server group.
	//
	// Minimum: 32
	//
	// Maximum: 32
	ProjectId string `json:"project_id,omitempty"`
	// Specifies the protocol used by the backend server group to receive requests. The value can be TCP, UDP, HTTP, HTTPS, or QUIC.
	//
	// Note:
	//
	// If the listener's protocol is UDP, the protocol of the backend server group must be UDP or QUIC.
	//
	// If the listener's protocol is TCP, the protocol of the backend server group must be TCP.
	//
	// If the listener's protocol is HTTP, the protocol of the backend server group must be HTTP.
	//
	// If the listener's protocol is HTTPS, the protocol of the backend server group can be HTTP or HTTPS.
	//
	// If the listener's protocol is TERMINATED_HTTPS, the protocol of the backend server group must be HTTP.
	//
	// If the backend server group protocol is QUIC, sticky session must be enabled with type set to SOURCE_IP.
	//
	// QUIC protocol is not supported in eu-nl region.
	//
	// Minimum: 1
	//
	// Maximum: 255
	Protocol string `json:"protocol" required:"true"`
	// Specifies the sticky session.
	SessionPersistence SessionPersistence `json:"session_persistence,omitempty"`
	// Specifies slow start details. After you enable slow start, new backend servers added to the backend server group are warmed up, and the number of requests they can receive increases linearly during the configured slow start duration.
	//
	// This parameter can be used when the protocol of the backend server group is HTTP or HTTPS. An error will be returned if the protocol is not HTTP or HTTPS. This parameter is not available in eu-nl region. Please do not use it.
	SlowStart SlowStart `json:"slow_start,omitempty"`
	// Specifies whether to enable removal protection for the load balancer.
	//
	// true: Enable removal protection.
	//
	// false (default): Disable removal protection.
	//
	// Disable removal protection for all your resources before deleting your account.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`
	// Specifies the ID of the VPC where the backend server group works.
	//
	// Note:
	//
	// The backend server group must be associated with the VPC.
	//
	// Only backend servers in the VPC or IP as Backend servers can be added.
	//
	// type must be set to instance.
	//
	// If vpc_id is not specified: vpc_id is determined by the VPC where the backend server works.
	//
	// Minimum: 0
	//
	// Maximum: 36
	VpcId string `json:"vpc_id,omitempty"`
	// Specifies the type of the backend server group.
	//
	// Values:
	//
	// instance: Any type of backend servers can be added. vpc_id is mandatory.
	//
	// ip: Only IP as Backend servers can be added. vpc_id cannot be specified.
	//
	// "": Any type of backend servers can be added.
	//
	// Note:
	//
	// If this parameter is not passed, any type of backend servers can be added. type will be returned as an empty string.
	//
	// Specify one of listener_id, loadbalancer_id, or type. Shared load balancers can only can added to the backend server group with loadbalancer_id or listener_id specified.
	//
	// Minimum: 0
	//
	// Maximum: 36
	Type string `json:"type,omitempty"`
}

// SessionPersistence represents the session persistence feature of the load
// balancing service. It attempts to force connections or requests in the same
// session to be processed by the same member as long as it is active. Three
// types of persistence are supported:
type SessionPersistence struct {
	// Specifies the cookie name. The value can contain only letters, digits, hyphens (-), underscores (_), and periods (.). Note: This parameter will take effect only when type is set to APP_COOKIE. Otherwise, an error will be returned.
	CookieName string `json:"cookie_name,omitempty"`
	// Specifies the sticky session type. The value can be SOURCE_IP, HTTP_COOKIE, or APP_COOKIE.Note:
	//
	// If the protocol of the backend server group is TCP or UDP, only SOURCE_IP takes effect.
	//
	// For dedicated load balancers, if the protocol of the backend server group is HTTP or HTTPS, the value can only be HTTP_COOKIE.
	//
	// If the backend server group protocol is QUIC, sticky session must be enabled with type set to SOURCE_IP.
	//
	// QUIC protocol is not supported in eu-nl region.
	Type string `json:"type" required:"true"`
	// Specifies the stickiness duration, in minutes. This parameter will not take effect when type is set to APP_COOKIE.
	//
	// If the protocol of the backend server group is TCP or UDP, the value ranges from 1 to 60, and the default value is 1.
	//
	// If the protocol of the backend server group is HTTP or HTTPS, the value ranges from 1 to 1440, and the default value is 1440.
	PersistenceTimeout *int `json:"persistence_timeout,omitempty"`
}

type SlowStart struct {
	// Specifies whether to enable slow start.
	//
	// true: Enable slow start.
	//
	// false: Disable slow start.
	//
	// Default: false
	Enable *bool `json:"enable,omitempty"`
	// Specifies the slow start duration, in seconds.
	//
	// The value ranges from 30 to 1200, and the default value is 30.
	//
	// Minimum: 30
	//
	// Maximum: 1200
	//
	// Default: 30
	Duration *int `json:"duration,omitempty"`
}

// Create accepts a CreateOpts struct and uses the values to create a new
// load balancer pool.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Pool, error) {
	b, err := build.RequestBody(opts, "pool")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("pools"), b, nil, nil)
	return extra(err, raw)
}
