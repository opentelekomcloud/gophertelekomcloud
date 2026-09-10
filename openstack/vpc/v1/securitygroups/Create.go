package securitygroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name                string `json:"name" required:"true"`
	VpcID               string `json:"vpc_id,omitempty"`
	EnterpriseProjectID string `json:"enterprise_project_id,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*SecurityGroup, error) {
	b, err := build.RequestBody(opts, "security_group")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL(client.ProjectID, "security-groups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res struct {
		SecurityGroup SecurityGroup `json:"security_group"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.SecurityGroup, nil
}
