package vpc_endpoint

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type EnableOpts struct {
	// Indicates whether to enable the internal DNS name .
	EndpointWithDnsName *bool `json:"endpointWithDnsName"`
}

// EnableVpcEndpoint function is used to Enable VCP endpoint of the cluster.
func Enable(client *golangsdk.ServiceClient, clusterID string, opts EnableOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	url := client.ServiceURL("clusters", clusterID, "vpcepservice", "open")
	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return err
}
