package services

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ActionOpts struct {
	// Specifies whether to accept or reject a VPC endpoint for a VPC endpoint service.
	// receive: means to accept the VPC endpoint.
	// reject: means to reject the VPC endpoint.
	Action string `json:"action" required:"true"`
	// Lists VPC endpoint IDs.
	// Each request accepts or rejects only one VPC endpoint
	Endpoints []string `json:"endpoints" required:"true"`
}

func Action(client *golangsdk.ServiceClient, id string, opts ActionOpts) ([]Connection, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("vpc-endpoint-services", id, "connections", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}
	var res []Connection
	err = extract.IntoSlicePtr(raw.Body, &res, "connections")
	return res, err
}
