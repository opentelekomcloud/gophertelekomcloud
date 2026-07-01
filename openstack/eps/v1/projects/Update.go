package projects

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name        string `json:"name" required:"true"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*EnterpriseProject, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("enterprise-projects", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res EnterpriseProject
	err = extract.IntoStructPtr(raw.Body, &res, "enterprise_project")
	return &res, err
}
