package pools

import "github.com/opentelekomcloud/gophertelekomcloud"

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToPoolUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the common options' struct used in this package's Update
// operation.
type UpdateOpts struct {
	// Name of the pool.
	Name *string `json:"name,omitempty"`

	// Human-readable description for the pool.
	Description *string `json:"description,omitempty"`

	// The algorithm used to distribute load between the members of the pool. The
	// current specification supports LBMethodRoundRobin, LBMethodLeastConnections
	// and LBMethodSourceIp as valid values for this attribute.
	LBMethod string `json:"lb_algorithm,omitempty"`

	// Specifies whether to enable sticky sessions.
	Persistence *SessionPersistence `json:"session_persistence,omitempty"`

	// The administrative state of the Pool. The value can only be updated to true.
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// Specifies whether to enable slow start.
	// This parameter is unsupported. Please do not use it.
	SlowStart *SlowStart `json:"slow_start,omitempty"`

	// Specifies whether to enable deletion protection for the load balancer.
	DeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`

	// Specifies the ID of the VPC where the backend server group works.
	VpcId string `json:"vpc_id,omitempty"`

	// Specifies the type of the backend server group.
	// Values:
	// instance: Any type of backend servers can be added. vpc_id is mandatory.
	// ip: Only cross-VPC backend servers can be added. vpc_id cannot be specified.
	// "": Any type of backend servers can be added.
	Type string `json:"type,omitempty"`
}

// ToPoolUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToPoolUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "pool")
}

// Update allows pools to be updated.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPoolUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(client.ServiceURL("pools", id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
