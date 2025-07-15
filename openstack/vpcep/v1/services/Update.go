package services

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Specifies whether connection approval is required.
	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`
	// Specifies the name of the VPC endpoint service.
	// The name can contain a maximum of 16 characters, including letters, digits, underscores (_), and hyphens (-).
	ServiceName string `json:"service_name,omitempty"`
	// Lists the port mappings opened to the VPC endpoint service.
	Ports []PortMapping `json:"ports,omitempty"`
	// Specifies the ID for identifying the backend resource of the VPC endpoint service.
	// The ID is in UUID format. The values are as follows:
	PortID string `json:"port_id,omitempty"`
	// Specifies whether the client IP address and port number or marker_id
	// information is transmitted to the server.
	TcpProxy string `json:"tcp_proxy,omitempty"`
	// Supplementary information about the VPC endpoint service.
	// The description can contain a maximum of 128 characters and cannot contain left angle brackets (<) or right angle brackets (>).
	Description string `json:"description,omitempty"`
	VIPPortID   string `json:"vip_port_id,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Service, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("vpc-endpoint-services", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}

	var res Service

	err = extract.Into(raw.Body, &res)
	return &res, err
}
