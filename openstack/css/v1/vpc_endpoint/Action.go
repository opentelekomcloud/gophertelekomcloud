package vpc_endpoint

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ActionOpts struct {
	// Specifies whether to accept or reject a VPC endpoint for a VPC endpoint service.
	// receive: means to accept the VPC endpoint.
	// reject: means to reject the VPC endpoint.
	Action string `json:"action" required:"true"`
	// Lists VPC endpoint IDs.
	// Each request accepts or rejects only one VPC endpoint
	Endpoints []string `json:"endpointIdList" required:"true"`
}

// Accept or Reject VPC Endpoint Connection
func Action(client *golangsdk.ServiceClient, clusterID string, opts ActionOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	url := client.ServiceURL("clusters", clusterID, "vpcepservice", "connections")

	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
