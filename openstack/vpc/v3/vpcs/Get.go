package vpcs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// Get retrieves a particular VPC based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (*Vpc, error) {
	// GET /v3/{project_id}/vpc/vpcs/{vpc_id}
	raw, err := client.Get(client.ServiceURL("vpcs", id), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res Vpc
	err = extract.IntoStructPtr(raw.Body, &res, "vpc")
	return &res, err
}

type CloudResource struct {
	// Resource type
	ResourceType string `json:"resource_type"`
	// Number of resources
	ResourceCount int `json:"resource_count"`
}

type Vpc struct {
	// VPC ID, which uniquely identifies the VPC
	// The value is in UUID format with hyphens (-).
	Id string `json:"id"`
	// VPC name
	// The value can contain no more than 64 characters, including letters,
	// digits, underscores (_), hyphens (-), and periods (.).
	Name string `json:"name"`
	// Provides supplementary information about the VPC.
	// The value can contain no more than 255 characters and cannot contain angle brackets (< or >).
	Description string `json:"description"`
	// Available VPC CIDR blocks
	// Value range:
	// 10.0.0.0/8-10.255.255.240/28
	// 172.16.0.0/12-172.31.255.240/28
	// 192.168.0.0/16-192.168.255.240/28
	// If cidr is not specified, the default value is "".
	// The value must be in IPv4 CIDR format, for example, 192.168.0.0/16.
	Cidr string `json:"cidr"`
	// Secondary CIDR blocks of VPCs
	// Currently, only IPv4 CIDR blocks are supported.
	SecondaryCidrs []string `json:"extend_cidrs"`
	// VPC status
	// Value range:
	// PENDING: The VPC is being created.
	// ACTIVE: The VPC is created successfully.
	Status string `json:"status"`
	// ID of the project to which the VPC belongs
	ProjectId string `json:"project_id"`
	// Time when the VPC is created
	// UTC time in the format of yyyy-MM-ddTHH:mmss
	CreatedAt string `json:"created_at"`
	// Time when the VPC is updated
	// UTC time in the format of yyyy-MM-ddTHH:mmss
	UpdatedAt string `json:"updated_at"`
	// Type and number of resources associated with the VPC
	// Currently, only route tables and subnets of the VPC are returned.
	// The number of virsubnets is the total number of IPv4 and IPv6 subnets.
	CloudResources []CloudResource `json:"cloud_resources"`
	// VPC tags. For details, see the tag objects.
	// Value range: 0 to 10 tag key-value pairs
	Tags []tags.ResourceTag `json:"tags"`
}
