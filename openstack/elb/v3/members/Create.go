package members

import "github.com/opentelekomcloud/gophertelekomcloud"

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToMemberCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the common options' struct used in this package's CreateMember
// operation.
type CreateOpts struct {
	// The IP address of the member to receive traffic from the load balancer.
	Address string `json:"address" required:"true"`

	// The port on which to listen for client traffic.
	ProtocolPort int `json:"protocol_port" required:"true"`

	// Name of the Member.
	Name string `json:"name,omitempty"`

	// ProjectID is the UUID of the project who owns the Member.
	// Only administrative users can specify a project UUID other than their own.
	ProjectID string `json:"project_id,omitempty"`

	// Specifies the weight of the backend server.
	//
	// Requests are routed to backend servers in the same backend server group based on their weights.
	//
	// If the weight is 0, the backend server will not accept new requests.
	//
	// This parameter is invalid when lb_algorithm is set to SOURCE_IP for the backend server group that contains the backend server.
	Weight *int `json:"weight,omitempty"`

	// If you omit this parameter, LBaaS uses the vip_subnet_id parameter value
	// for the subnet UUID.
	SubnetID string `json:"subnet_cidr_id,omitempty"`

	// The administrative state of the Pool. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
}

// ToMemberCreateMap builds a request body from CreateOptsBuilder.
func (opts CreateOpts) ToMemberCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "member")
}

// Create will create and associate a Member with a particular Pool.
func Create(client *golangsdk.ServiceClient, poolID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToMemberCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("pools", poolID, "members"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}
