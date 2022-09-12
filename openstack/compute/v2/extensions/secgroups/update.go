package secgroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Update will modify the mutable properties of a security group, notably its name and description.
func Update(client *golangsdk.ServiceClient, id string, opts GroupOpts) (*SecurityGroup, error) {
	b, err := golangsdk.BuildRequestBody(opts, "security_group")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("os-security-groups", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
