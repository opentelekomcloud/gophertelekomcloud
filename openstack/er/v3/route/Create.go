package route

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	RouteTableId string `json:"-" required:"true"`
	Destination  string `json:"destination" required:"true"`
	AttachmentId string `json:"attachment_id,omitempty"`
	IsBlackhole  *bool  `json:"is_blackhole,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Route, error) {
	b, err := build.RequestBody(opts, "route")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("enterprise-router", "route-tables", opts.RouteTableId, "static-routes"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res Route
	return &res, extract.IntoStructPtr(raw.Body, &res, "route")
}

type Route struct {
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	State        string            `json:"state"`
	IsBlackhole  bool              `json:"is_blackhole"`
	Destination  string            `json:"destination"`
	RouteTableId string            `json:"route_table_id"`
	CreatedAt    string            `json:"created_at"`
	UpdatedAt    string            `json:"updated_at"`
	Attachments  []RouteAttachment `json:"attachments"`
}

type RouteAttachment struct {
	ResourceId   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
	AttachmentId string `json:"attachment_id"`
}
