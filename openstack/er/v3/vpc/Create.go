package vpc

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	RouterID            string `json:"-" required:"true"`
	VpcId               string `json:"vpc_id" required:"true"`
	Name                string `json:"name" required:"true"`
	SubnetId            string `json:"virsubnet_id" required:"true"`
	Description         string `json:"description,omitempty"`
	AutoCreateVpcRoutes bool   `json:"auto_create_vpc_routes,omitempty"`
	// Feature not available
	// Ipv6Enable          bool               `json:"ipv6_enable,omitempty"`
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*VpcAttachmentDetails, error) {
	b, err := build.RequestBody(opts, "vpc_attachment")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("enterprise-router", opts.RouterID, "vpc-attachments"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res VpcAttachmentDetails
	return &res, extract.IntoStructPtr(raw.Body, &res, "vpc_attachment")
}

type VpcAttachmentDetails struct {
	ID                  string             `json:"id"`
	Name                string             `json:"name"`
	VpcId               string             `json:"vpc_id"`
	SubnetId            string             `json:"virsubnet_id"`
	AutoCreateVpcRoutes bool               `json:"auto_create_vpc_routes"`
	State               string             `json:"state"`
	CreatedAt           string             `json:"created_at"`
	UpdatedAt           string             `json:"updated_at"`
	Tags                []tags.ResourceTag `json:"tags"`
	Description         string             `json:"description"`
	ProjectId           string             `json:"project_id"`
	VpcProjectId        string             `json:"vpc_project_id"`
	Ipv6Enable          bool               `json:"ipv6_enable"`
}
