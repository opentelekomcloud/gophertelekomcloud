package secgroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// GroupOpts is the underlying struct responsible for creating or updating
// security groups. It therefore represents the mutable attributes of a security group.
type GroupOpts struct {
	// the name of your security group.
	Name string `json:"name" required:"true"`
	// the description of your security group.
	Description string `json:"description" required:"true"`
}

// Create will create a new security group.
func Create(client *golangsdk.ServiceClient, opts GroupOpts) (*SecurityGroup, error) {
	b, err := build.RequestBody(opts, "security_group")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("os-security-groups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
