package pools

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts is the common options' struct used in this package's Update
// operation.
type UpdateOpts struct {
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
	LbAlgorithm string `json:"lb_algorithm,omitempty"`
	// Specifies the backend server group name.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Name string `json:"name,omitempty"`
	// Specifies the sticky session.
	SessionPersistence SessionPersistence `json:"session_persistence,omitempty"`
	// Specifies slow start details. After you enable slow start, new backend servers added to the backend server group are warmed up, and the number of requests they can receive increases linearly during the configured slow start duration.
	//
	// This parameter can be used when the protocol of the backend server group is HTTP or HTTPS. An error will be returned if the protocol is not HTTP or HTTPS.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	SlowStart SlowStart `json:"slow_start,omitempty"`
	// Specifies whether to enable removal protection for the load balancer.
	//
	// true: Enable removal protection.
	//
	// false: Disable removal protection.
	//
	// Disable removal protection for all your resources before deleting your account.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`
	// Specifies the ID of the VPC where the backend server group works.
	//
	// This parameter can be updated only when vpc_id is left blank.
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
	// Note: This parameter can be updated only when type is left blank.
	Type string `json:"type,omitempty"`
}

// Update allows pools to be updated.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Pool, error) {
	b, err := build.RequestBody(opts, "pool")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("pools", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
