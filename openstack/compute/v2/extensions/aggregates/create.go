package aggregates

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create makes a request against the API to create an aggregate.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (r CreateResult) {
	b, err := opts.ToAggregatesCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("os-aggregates"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
