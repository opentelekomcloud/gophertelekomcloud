package direct_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// description of the direct connect
	Description string `json:"description,omitempty"`
	// name of the direct connect
	Name string `json:"name,omitempty"`
	// bandwidth of the direct connect in Mbps
	Bandwidth int `json:"bandwidth,omitempty"`
	// provider_status specifies the status of the carrier's leased line.
	// The value can be ACTIVE or DOWN.
	ProviderStatus string `json:"provider_status,omitempty"`
}

// Update is an operation which modifies the attributes of the specified direct connect
func Update(c *golangsdk.ServiceClient, id string, opts UpdateOpts) (err error) {
	b, err := build.RequestBody(opts, "direct_connect")
	if err != nil {
		return
	}

	_, err = c.Put(c.ServiceURL("dcaas", "direct-connects", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})

	return
}
