package servergroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create requests the creation of a new Server Group.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToServerGroupCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("os-server-groups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
