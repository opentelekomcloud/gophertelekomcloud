package projects

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*EnterpriseProject, error) {
	raw, err := client.Get(client.ServiceURL("enterprise-projects", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res EnterpriseProject
	err = extract.IntoStructPtr(raw.Body, &res, "enterprise_project")
	return &res, err
}
