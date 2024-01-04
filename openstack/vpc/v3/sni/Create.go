package sni

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateSubNetworkInterfaceOption struct {
	// Virtual subnet ID
	// The value must be in standard UUID format.
	SubnetId string `json:"virsubnet_id"`
	// VLAN ID of the supplementary network interface
	// The value can be from 1 to 4094.
	// Each supplementary network interface of an elastic network interface has a unique VLAN ID.
	VlanId string `json:"vlan_id,omitempty"`
	// ID of the elastic network interface
	// The value must be in standard UUID format.
	// The value must be an existing port ID.
	PortId string `json:"parent_id"`
	// Description of the supplementary network interface
	// The value can contain no more than 255 characters and cannot contain angle brackets (< or >).
	Description string `json:"description,omitempty"`
	// Whether to enable IPv6 for the supplementary network interface.
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`
	// Private IPv4 address of the supplementary network interface
	// The value must be within the virtual subnet. If this parameter is left blank, an IP address will be randomly assigned.
	PrivateIpAddress string `json:"private_ip_address,omitempty"`
	// IPv6 address of the supplementary network interface
	// If this parameter is left blank, an IP address will be randomly assigned.
	Ipv6IpAddress string `json:"ipv6_ip_address,omitempty"`
	// Security group IDs
	// Example: "security_groups": ["a0608cbf-d047-4f54-8b28-cd7b59853fff"]
	// The default value is the default security group.
	SecurityGroups []string `json:"security_groups,omitempty"`
	// Project ID of the supplementary network interface
	// The value must be in standard UUID format.
	// Only administrators have permissions to specify project IDs.
	ProjectId string `json:"project_id,omitempty"`
}

type CreateOpts struct {
	// Whether to only check the request.
	// Value range:
	// true: Only the check request will be sent and no secondary CIDR block will be added.
	// Check items include mandatory parameters, request format, and constraints.
	// If the check fails, an error will be returned. If the check succeeds, response code 202 will be returned.
	// false (default value): A request will be sent and a secondary CIDR block will be added.
	DryRun *bool `json:"dry_run,omitempty"`
	// Request body for adding a secondary CIDR block.
	Sni *CreateSubNetworkInterfaceOption `json:"sub_network_interface"`
}

// Create is used to create a supplementary network interface.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*SubNetworkInterface, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/vpc/sub-network-interfaces
	raw, err := client.Post(client.ServiceURL("sub-network-interfaces"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res SubNetworkInterface
	err = extract.IntoStructPtr(raw.Body, &res, "sub_network_interface")
	return &res, err
}
