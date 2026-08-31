package vpcs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name                string `json:"name,omitempty"`
	Description         string `json:"description,omitempty"`
	CIDR                string `json:"cidr,omitempty"`
	EnterpriseProjectID string `json:"enterprise_project_id,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Vpc, error) {
	b, err := build.RequestBody(opts, "vpc")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL(client.ProjectID, "vpcs"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res struct {
		VPC Vpc `json:"vpc"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.VPC, err
}
