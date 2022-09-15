package clusters

import "github.com/opentelekomcloud/gophertelekomcloud"

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToClusterCreateMap()
	if err != nil {
		return nil, err
	}

	raw, err = client.Post(client.ServiceURL("clusters"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
