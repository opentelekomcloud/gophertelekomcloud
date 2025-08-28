package transitip

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateTransitIpOpts struct {
	// Specifies the subnet ID of the current project.
	VirSubnetID string `json:"virsubnet_id" required:"true"`
	// Specifies the transit IP address.
	IpAddress string `json:"ip_address,omitempty"`
	// Specifies the ID of the enterprise project that is associated with the transit IP address when the transit IP address is being assigned.
	// Default value: 0.
	EnterpriseProjectID string `json:"enterprise_project_id,omitempty"`
	// Specifies the tag list.
	Tags []TagOptions `json:"tags,omitempty"`
}

type TagOptions struct {
	// Specifies the tag key. A key can contain up to 128 Unicode characters.
	// Key cannot be left blank.
	Key string `json:"key" required:"true"`
	// Specifies the tag value. Each value can contain up to 255 Unicode characters.
	Value string `json:"value" required:"true"`
}

// This function is used to assign a transit IP address.
func Create(client *golangsdk.ServiceClient, opts CreateTransitIpOpts) (*TransitIPCommonResponse, error) {
	b, err := build.RequestBody(opts, "transit_ip")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/private-nat/transit-ips
	raw, err := client.Post(client.ServiceURL("private-nat", "transit-ips"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res TransitIPCommonResponse
	return &res, extract.Into(raw.Body, &res)
}
