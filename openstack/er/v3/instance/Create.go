package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	tag "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	// Enterprise router name. The value can contain 1 to 64 characters, including letters, digits,
	// underscores (_), hyphens (-), and periods (.).
	Name string `json:"name" required:"true"`
	// Supplementary information about an enterprise router
	Description string `json:"description,omitempty"`
	// Enterprise router BGP ASN. Specify a dedicated ASN in the range of 64512-65534 or 4200000000-4294967294.
	// ASN can only be set during enterprise router creation.
	Asn float64 `json:"asn" required:"true"`
	// Default: postPaid
	ChargeMode string `json:"charge_mode,omitempty"`
	// Tag information
	Tags []tag.ResourceTag `json:"tags,omitempty"`
	//
	// Whether to enable the Default Route Table Propagation function.
	// The default value is false, indicating that the function is disabled.
	EnableDefaultPropagation *bool `json:"enable_default_propagation,omitempty"`
	// Whether to enable the Default Route Table Association function.
	EnableDefaultAssociation *bool `json:"enable_default_association,omitempty"`
	// AZs where the enterprise router is located
	AvailabilityZoneIDs []string `json:"availability_zone_ids" required:"true"`
	// Whether to enable Auto Accept Shared Attachments.
	AutoAcceptSharedAttachments *bool `json:"auto_accept_shared_attachments,omitempty"`
	// Enterprise router CIDR block. This parameter is not supported for now.
	CidrBlocks []string `json:"cidr_blocks,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*RouterInstanceResp, error) {
	b, err := build.RequestBody(opts, "instance")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("enterprise-router", "instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res RouterInstanceResp
	return &res, extract.Into(raw.Body, &res)
}

type RouterInstanceResp struct {
	// Enterprise router
	Instance *RouterInstance `json:"instance"`
	// Request ID
	RequestID string `json:"request_id"`
}

type RouterInstance struct {
	// Enterprise router ID
	ID string `json:"id"`
	// Enterprise router name
	Name string `json:"name"`
	// Supplementary information about an enterprise router
	Description string `json:"description"`
	// Enterprise router status.
	State string `json:"state"`
	// Tag information
	Tags []tag.ResourceTag `json:"tags"`
	// Default: postPaid
	ChargeMode string `json:"charge_mode"`
	// Creation time in the format YYYY-MM-DDTHH:mm:ss.sssZ
	CreatedAt string `json:"created_at"`
	// Update time in the format YYYY-MM-DDTHH:mm:ss.sssZ
	UpdatedAt string `json:"updated_at"`
	// Project ID
	ProjectID string `json:"project_id"`
	// Enterprise router BGP ASN
	Asn float64 `json:"asn"`
	// Whether to enable the Default Route Table Propagation function.
	EnableDefaultPropagation bool `json:"enable_default_propagation"`
	// Whether to enable the Default Route Table Association function.
	EnableDefaultAssociation bool `json:"enable_default_association"`
	// Default propagation route table ID
	DefaultPropagationRouteTableID string `json:"default_propagation_route_table_id"`
	// Default association route table ID
	DefaultAssociationRouteTableID string `json:"default_association_route_table_id"`
	// AZs where the enterprise router is located
	AvailabilityZoneIDs []string `json:"availability_zone_ids"`
	// Whether to automatically accept shared attachments.
	AutoAcceptSharedAttachments bool `json:"auto_accept_shared_attachments"`
	// Enterprise router CIDR block. This parameter is not supported for now.
	CidrBlocks []string `json:"cidr_blocks"`
}
