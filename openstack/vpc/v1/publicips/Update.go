package publicips

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	PortId string `json:"port_id,omitempty"`
	// IPVersion is not documented as an update request parameter in the
	// reviewed OTC documentation.
	IPVersion int    `json:"ip_version,omitempty"`
	Alias     string `json:"alias,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*PublicIPUpdateResp, error) {
	b, err := build.RequestBody(opts, "publicip")
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL(client.ProjectID, "publicips", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res struct {
		PublicIP PublicIPUpdateResp `json:"publicip"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.PublicIP, err
}
