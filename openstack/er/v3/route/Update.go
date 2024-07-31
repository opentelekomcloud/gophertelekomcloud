package route

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	RouteTableId string `json:"-"`
	RouteId      string `json:"-"`
	AttachmentId string `json:"attachment_id,omitempty"`
	IsBlackhole  *bool  `json:"is_blackhole,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*Route, error) {
	b, err := build.RequestBody(opts, "route")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("enterprise-router", "route-tables", opts.RouteTableId, "static-routes", opts.RouteId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}

	var res Route
	return &res, extract.IntoStructPtr(raw.Body, &res, "route")
}
