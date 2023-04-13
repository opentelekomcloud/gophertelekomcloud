package members

import "github.com/opentelekomcloud/gophertelekomcloud"

// UpdateOptsBuilder allows extensions to add additional parameters to the
// List request.
type UpdateOptsBuilder interface {
	ToMemberUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the common options' struct used in this package's Update
// operation.
type UpdateOpts struct {
	// Name of the Member.
	Name *string `json:"name,omitempty"`

	// A positive integer value that indicates the relative portion of traffic
	// that this member should receive from the pool. For example, a member with
	// a weight of 10 receives five times as much traffic as a member with a
	// weight of 2.
	Weight *int `json:"weight,omitempty"`

	// The administrative state of the Pool. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
}

// ToMemberUpdateMap builds a request body from UpdateOptsBuilder.
func (opts UpdateOpts) ToMemberUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "member")
}

// Update allows Member to be updated.
func Update(client *golangsdk.ServiceClient, poolID string, memberID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToMemberUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(client.ServiceURL("pools", poolID, "members", memberID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	return
}
