package servergroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToServerGroupCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies Server Group creation parameters.
type CreateOpts struct {
	// Name is the name of the server group
	Name string `json:"name" required:"true"`

	// Policies are the server group policies
	Policies []string `json:"policies" required:"true"`
}

// ToServerGroupCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToServerGroupCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "server_group")
}
