package bandwidths

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Specifies the bandwidth name. The value can contain 1 to 64
	// characters, including letters, digits, underscores (_), hyphens (-),
	// and periods (.).
	Name string `json:"name" required:"true"`

	// Specifies the bandwidth size. The shared bandwidth has a minimum
	// limit, which may vary depending on sites. The default minimum value
	// is 5 Mbit/s. The value ranges from 1 Mbit/s to 1000 Mbit/s by
	// default.
	Size int `json:"size" required:"true"`

	// Specifies the enterprise project ID.
	//
	// This parameter is unsupported by OTC; do not use it.
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`

	// Specifies whether it is in a central site or an edge site. This
	// resource can only be associated with an EIP of the same region.
	PublicBorderGroup string `json:"public_border_group,omitempty"`
}

// Create assigns a new shared bandwidth.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Bandwidth, error) {
	b, err := build.RequestBody(opts, "bandwidth")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL(client.ProjectID, "bandwidths"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res struct {
		Bandwidth Bandwidth `json:"bandwidth"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.Bandwidth, err
}
