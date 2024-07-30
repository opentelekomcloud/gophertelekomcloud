package propagation

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	RouterID     string `json:"-" required:"true"`
	RouteTableID string `json:"-" required:"true"`
	AttachmentID string `json:"attachment_id,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Propagation, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("enterprise-router", opts.RouterID, "route-tables", opts.RouteTableID, "enable-propagations"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res Propagation
	return &res, extract.IntoStructPtr(raw.Body, &res, "propagation")
}

type Propagation struct {
	ID           string `json:"id"`
	ProjectID    string `json:"project_id"`
	ErID         string `json:"er_id"`
	RouteTableID string `json:"route_table_id"`
	AttachmentID string `json:"attachment_id"`
	ResourceType string `json:"resource_type"`
	ResourceID   string `json:"resource_id"`
	State        string `json:"state"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
