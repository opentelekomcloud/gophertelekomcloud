package certificates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Domain      string `json:"domain,omitempty"`
	PrivateKey  string `json:"private_key,omitempty"`
	Certificate string `json:"certificate" required:"true"`
	ProjectID   string `json:"project_id,omitempty"`

	// These fields are documented but marked unsupported by the reviewed OTC API.
	AdminStateUp        *bool  `json:"admin_state_up,omitempty"`
	EnterpriseProjectID string `json:"enterprise_project_id,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Certificate, error) {
	b, err := build.RequestBody(opts, "certificate")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("certificates"), b, nil, &golangsdk.RequestOpts{OkCodes: []int{201}})
	if err != nil {
		return nil, err
	}
	var res struct {
		Certificate Certificate `json:"certificate"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.Certificate, nil
}
