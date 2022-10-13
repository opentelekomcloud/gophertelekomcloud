package secgroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// Update will modify the mutable properties of a security group, notably its name and description.
func Update(client *golangsdk.ServiceClient, id string, opts GroupOpts) (*SecurityGroup, error) {
	b, err := build.RequestBody(opts, "security_group")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("os-security-groups", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
