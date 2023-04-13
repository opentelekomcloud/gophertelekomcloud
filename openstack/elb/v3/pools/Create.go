package pools

import "github.com/opentelekomcloud/gophertelekomcloud"

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToPoolCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the common options' struct used in this package's Create
// operation.
type CreateOpts struct {
	// The algorithm used to distribute load between the members of the pool.
	LBMethod string `json:"lb_algorithm" required:"true"`

	// The protocol used by the pool members, you can use either
	// ProtocolTCP, ProtocolHTTP, or ProtocolHTTPS.
	Protocol string `json:"protocol" required:"true"`

	// The Loadbalancer on which the members of the pool will be associated with.
	// Note: one of LoadbalancerID or ListenerID must be provided.
	LoadbalancerID string `json:"loadbalancer_id,omitempty"`

	// The Listener on which the members of the pool will be associated with.
	// Note: one of LoadbalancerID or ListenerID must be provided.
	ListenerID string `json:"listener_id,omitempty"`

	// ProjectID is the UUID of the project who owns the Pool.
	// Only administrative users can specify a project UUID other than their own.
	ProjectID string `json:"project_id,omitempty"`

	// Name of the pool.
	Name string `json:"name,omitempty"`

	// Human-readable description for the pool.
	Description string `json:"description,omitempty"`

	// Persistence is the session persistence of the pool.
	// Omit this field to prevent session persistence.
	Persistence *SessionPersistence `json:"session_persistence,omitempty"`

	SlowStart *SlowStart `json:"slow_start,omitempty"`

	// The administrative state of the Pool. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// Specifies whether to enable deletion protection for the pool.
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

// ToPoolCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToPoolCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "pool")
}

// Create accepts a CreateOpts struct and uses the values to create a new
// load balancer pool.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPoolCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("pools"), b, &r.Body, nil)
	return
}
