package pools

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
)

// Get retrieves a particular pool based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (*Pool, error) {
	raw, err := client.Get(client.ServiceURL("pools", id), nil, nil)
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*Pool, error) {
	if err != nil {
		return nil, err
	}

	var res Pool
	err = extract.IntoStructPtr(raw.Body, &res, "pool")
	return &res, err
}

// Pool represents a logical set of devices, such as web servers, that you
// group together to receive and process traffic. The load balancing function
// chooses a Member of the Pool according to the configured load balancing
// method to handle the new requests or connections received on the VIP address.
type Pool struct {
	// Specifies the administrative status of the backend server group. The value can only be true.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `json:"admin_state_up"`
	// Provides supplementary information about the backend server group.
	Description string `json:"description"`
	// Specifies the ID of the health check configured for the backend server group.
	HealthmonitorId string `json:"healthmonitor_id"`
	// Specifies the backend server group ID.
	Id string `json:"id"`
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
	LbAlgorithm string `json:"lb_algorithm"`
	// Specifies the IDs of the listeners with which the backend server group is associated.
	Listeners []structs.ResourceRef `json:"listeners"`
	// Specifies the IDs of the load balancers with which the backend server group is associated.
	Loadbalancers []structs.ResourceRef `json:"loadbalancers"`
	// Specifies the IDs of the backend servers in the backend server group.
	Members []structs.ResourceRef `json:"members"`
	// Specifies the backend server group name.
	Name string `json:"name"`
	// Specifies the project ID.
	ProjectId string `json:"project_id"`
	// Specifies the protocol used by the backend server group to receive requests. The value can be TCP, UDP, HTTP, HTTPS, or QUIC.
	//
	// If the listener's protocol is UDP, the protocol of the backend server group must be UDP.
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
	Protocol string `json:"protocol"`
	// Specifies the sticky session.
	SessionPersistence SessionPersistence `json:"session_persistence"`
	// Specifies the IP address version supported by the backend server group.
	//
	// IPv6 is unsupported. Only v4 will be returned.
	IpVersion string `json:"ip_version"`
	// Specifies slow start details. After you enable slow start, new backend servers added to the backend server group are warmed up, and the number of requests they can receive increases linearly during the configured slow start duration.
	//
	// This parameter can be used when the protocol of the backend server group is HTTP or HTTPS. An error will be returned if the protocol is not HTTP or HTTPS. This parameter is not available in eu-nl region. Please do not use it.
	SlowStart SlowStart `json:"slow_start"`
	// Specifies whether to enable removal protection.
	//
	// true: Enable removal protection.
	//
	// false: Disable removal protection.
	//
	// Disable removal protection for all your resources before deleting your account.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable"`
	// Specifies the time when a backend server group was created. The format is yyyy-MM-dd'T'HH:mm:ss'Z' (UTC time).
	//
	// This is a new field in this version, and it will not be returned for resources associated with existing dedicated load balancers and for resources associated with existing and new shared load balancers.
	CreatedAt string `json:"created_at"`
	// Specifies the time when when a backend server group was updated. The format is yyyy-MM-dd'T'HH:mm:ss'Z' (UTC time).
	//
	// This is a new field in this version, and it will not be returned for resources associated with existing dedicated load balancers and for resources associated with existing and new shared load balancers.
	UpdatedAt string `json:"updated_at"`
	// Specifies the ID of the VPC where the backend server group works.
	VpcId string `json:"vpc_id"`
	// Specifies the type of the backend server group.
	//
	// Values:
	//
	// instance: Any type of backend servers can be added. vpc_id is mandatory.
	//
	// ip: Only IP as Backend servers can be added. vpc_id cannot be specified.
	//
	// "": Any type of backend servers can be added.
	Type string `json:"type"`
}
