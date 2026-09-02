package bandwidths

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Specifies the bandwidth name. The value is a string of 1 to 64
	// characters that can contain letters, digits, underscores (_), and
	// hyphens (-). Either Name or Size must be specified.
	Name string `json:"name,omitempty"`

	// Specifies the bandwidth size in Mbit/s. Either Name or Size must be
	// specified.
	Size int `json:"size,omitempty"`
}

// Update modifies the name and/or size of a bandwidth.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*BandWidth, error) {
	b, err := build.RequestBody(opts, "bandwidth")
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL(client.ProjectID, "bandwidths", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res struct {
		Bandwidth BandWidth `json:"bandwidth"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.Bandwidth, err
}
