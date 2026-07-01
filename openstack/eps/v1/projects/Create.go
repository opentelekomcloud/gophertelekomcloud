package projects

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name        string `json:"name" required:"true"`
	Description string `json:"description,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*EnterpriseProject, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("enterprise-projects"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}

	var res EnterpriseProject
	err = extract.IntoStructPtr(raw.Body, &res, "enterprise_project")
	return &res, err
}
