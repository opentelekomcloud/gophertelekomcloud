package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Enterprise router ID
	InstanceID string `json:"-"`
	// Enterprise router name. The value can contain 1 to 64 characters, including letters,
	// digits, underscores (_), hyphens (-), and periods (.).
	Name string `json:"name,omitempty"`
	// Supplementary information about an enterprise router
	Description string `json:"description,omitempty"`
	// Whether to enable Default Route Table Propagation.
	EnableDefaultPropagation *bool `json:"enable_default_propagation,omitempty"`
	// Whether to enable Default Route Table Association.
	EnableDefaultAssociation *bool `json:"enable_default_association,omitempty"`
	// Default propagation route table ID
	DefaultPropagationRouteTableID string `json:"default_propagation_route_table_id,omitempty"`
	// Default association route table ID
	DefaultAssociationRouteTableID string `json:"default_association_route_table_id,omitempty"`
	// Whether to automatically accept shared attachments.
	AutoAcceptSharedAttachments *bool `json:"auto_accept_shared_attachments,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*RouterInstanceResp, error) {
	b, err := build.RequestBody(opts, "instance")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("enterprise-router", "instances", opts.InstanceID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res RouterInstanceResp
	return &res, extract.Into(raw.Body, &res)
}
