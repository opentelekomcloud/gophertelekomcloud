package vpcs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// RemoveSecondaryCidr is used to add a secondary CIDR block to a VPC.
func RemoveSecondaryCidr(client *golangsdk.ServiceClient, id string, opts CidrOpts) (*Vpc, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/vpc/vpcs/{vpc_id}/remove-extend-cidr
	raw, err := client.Put(client.ServiceURL("vpcs", id, "remove-extend-cidr"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Vpc
	err = extract.IntoStructPtr(raw.Body, &res, "vpc")
	return &res, err
}
