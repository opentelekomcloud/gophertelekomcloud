package servergroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateOpts specifies Server Group creation parameters.
type CreateOpts struct {
	// Name is the name of the server group
	Name string `json:"name" required:"true"`
	// Policies are the server group policies
	Policies []string `json:"policies" required:"true"`
}

// Create requests the creation of a new Server Group.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*ServerGroup, error) {
	b, err := golangsdk.BuildRequestBody(opts, "server_group")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("os-server-groups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
