package secgroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create will create a new security group.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToSecGroupCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(rootURL(client), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
