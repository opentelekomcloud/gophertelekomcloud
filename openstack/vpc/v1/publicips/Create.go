package publicips

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type PublicIPRequest struct {
	Type string `json:"type" required:"true"`
	// IpAddress is not documented in the reviewed OTC documentation.
	IpAddress string `json:"ip_address,omitempty"`
	IPVersion int    `json:"ip_version,omitempty"`
	Alias     string `json:"alias,omitempty"`
}

type BandWidth struct {
	Name       string `json:"name" required:"true"`
	Size       int    `json:"size" required:"true"`
	ID         string `json:"id,omitempty"`
	ShareType  string `json:"share_type" required:"true"`
	ChargeMode string `json:"charge_mode,omitempty"`
}

type CreateOpts struct {
	Publicip  PublicIPRequest `json:"publicip" required:"true"`
	Bandwidth BandWidth       `json:"bandwidth" required:"true"`
	// EnterpriseProjectId is not documented as a create request parameter in
	// the reviewed OTC documentation.
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*PublicIPCreateResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL(client.ProjectID, "publicips"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res struct {
		PublicIP PublicIPCreateResp `json:"publicip"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.PublicIP, err
}
