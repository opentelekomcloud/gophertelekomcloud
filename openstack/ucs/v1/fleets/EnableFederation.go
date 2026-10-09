package fleets

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type EnableFederationOpts struct {
	RetryJoinAll *bool `q:"retryjoinall"`
}

func EnableFederation(client *golangsdk.ServiceClient, id string, opts EnableFederationOpts) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("clustergroups", id, "federations").WithQueryParams(&opts).Build()
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL(url.String()), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return err
}
