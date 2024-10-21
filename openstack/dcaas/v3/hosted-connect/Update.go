package hosted_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Provides supplementary information about the connection.
	Description string `json:"description,omitempty"`
	// Specifies the connection name.
	Name string `json:"name,omitempty"`
	// Specifies the bandwidth of the hosted connection, in Mbit/s.
	Bandwidth int `json:"bandwidth,omitempty"`
	// Specifies the location of the on-premises facility
	// at the other end of the connection, specific to the street or data center name.
	PeerLocation string `json:"peer_location,omitempty"`
}

// Update is an operation which modifies the attributes of the specified hosted connect
func Update(c *golangsdk.ServiceClient, id string, opts UpdateOpts) (*HostedConnect, error) {
	// PUT /v3/{project_id}/dcaas/hosted-connects/{hosted_connect_id}
	b, err := build.RequestBody(opts, "hosted_connect")
	if err != nil {
		return nil, err
	}

	raw, err := c.Put(c.ServiceURL("dcaas", "hosted-connects", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})

	var res HostedConnect

	err = extract.IntoStructPtr(raw.Body, &res, "hosted_connect")
	return &res, err
}
