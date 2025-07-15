package route_table

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	RouterID     string  `json:"-"`
	RouteTableId string  `json:"-"`
	Name         string  `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	// parameter not supported ☻
	// BgpOptions   *BgpOptions `json:"bgp_options,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*RouteTable, error) {
	b, err := build.RequestBody(opts, "route_table")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("enterprise-router", opts.RouterID, "route-tables", opts.RouteTableId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res RouteTable
	return &res, extract.IntoStructPtr(raw.Body, &res, "route_table")
}
