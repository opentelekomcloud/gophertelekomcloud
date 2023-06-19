package volumetransfers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// AcceptOpts contains options for a Volume transfer accept reqeust.
type AcceptOpts struct {
	// Specifies the authentication key of the disk transfer.
	// Specifies the authentication key returned during the disk transfer creation.
	AuthKey string `json:"auth_key" required:"true"`
}

// Accept will accept a volume transfer request based on the values in AcceptOpts.
func Accept(client *golangsdk.ServiceClient, id string, opts AcceptOpts) (*Transfer, error) {
	b, err := build.RequestBody(opts, "accept")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/os-volume-transfer/{transfer_id}/accept
	raw, err := client.Post(client.ServiceURL("os-volume-transfer", id, "accept"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return extra(err, raw)
}
