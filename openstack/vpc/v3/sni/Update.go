package sni

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateSubNetworkInterfaceOption struct {
	// Description of the supplementary network interface
	// The value can contain no more than 255 characters and cannot contain angle brackets (< or >).
	Description string `json:"description,omitempty"`
	// Security group IDs
	// Example: "security_groups": ["a0608cbf-d047-4f54-8b28-cd7b59853fff"]
	// The default value is the default security group.
	SecurityGroups []string `json:"security_groups,omitempty"`
}

type UpdateOpts struct {
	// Whether to only check the request.
	// Value range:
	// true: Only the check request will be sent and no secondary CIDR block will be added.
	// Check items include mandatory parameters, request format, and constraints.
	// If the check fails, an error will be returned. If the check succeeds, response code 202 will be returned.
	// false (default value): A request will be sent and a secondary CIDR block will be added.
	DryRun *bool `json:"dry_run,omitempty"`
	// Request body for adding a secondary CIDR block.
	Sni *UpdateSubNetworkInterfaceOption `json:"sub_network_interface"`
}

// Update is used to update a supplementary network interface.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*SubNetworkInterface, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/vpc/sub-network-interfaces/{sub_network_interface_id}
	raw, err := client.Put(client.ServiceURL("sub-network-interfaces"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res SubNetworkInterface
	err = extract.IntoStructPtr(raw.Body, &res, "sub_network_interface")
	return &res, err
}
