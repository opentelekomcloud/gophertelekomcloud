package vpcs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type AddExtendCidrOption struct {
	// Secondary CIDR blocks that can be added to VPCs
	// The value cannot contain the following:
	// 100.64.0.0/10
	// 214.0.0.0/7
	// 198.18.0.0/15
	// 169.254.0.0/16
	// 0.0.0.0/8
	// 127.0.0.0/8
	// 240.0.0.0/4
	// 172.31.0.0/16
	// 192.168.0.0/16
	// Currently, only one secondary CIDR block can be added to each VPC.
	ExtendCidrs []string `json:"extend_cidrs"`
}

type CidrOpts struct {
	// Whether to only check the request.
	// Value range:
	// true: Only the check request will be sent and no secondary CIDR block will be added.
	// Check items include mandatory parameters, request format, and constraints.
	// If the check fails, an error will be returned. If the check succeeds, response code 202 will be returned.
	// false (default value): A request will be sent and a secondary CIDR block will be added.
	DryRun *bool `json:"dry_run,omitempty"`
	// Request body for adding a secondary CIDR block.
	Vpc *AddExtendCidrOption `json:"vpc"`
}

// AddSecondaryCidr is used to add a secondary CIDR block to a VPC.
func AddSecondaryCidr(client *golangsdk.ServiceClient, id string, opts CidrOpts) (*Vpc, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/vpc/vpcs/{vpc_id}/add-extend-cidr
	raw, err := client.Put(client.ServiceURL("vpcs", id, "add-extend-cidr"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Vpc
	err = extract.IntoStructPtr(raw.Body, &res, "vpc")
	return &res, err
}
