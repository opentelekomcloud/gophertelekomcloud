package route_table

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	RouterID    string  `json:"-" required:"true"`
	Name        string  `json:"name" required:"true"`
	Description *string `json:"description,omitempty"`
	// parameter not supported ☻
	// BgpOptions  *BgpOptions        `json:"bgp_options,omitempty"`
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

type BgpOptions struct {
	LoadBalancingAsPathIgnore *bool `json:"load_balancing_as_path_ignore,omitempty"`
	LoadBalancingAsPathRelax  *bool `json:"load_balancing_as_path_relax,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*RouteTable, error) {
	b, err := build.RequestBody(opts, "route_table")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("enterprise-router", opts.RouterID, "route-tables"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res RouteTable
	return &res, extract.IntoStructPtr(raw.Body, &res, "route_table")
}

type RouteTable struct {
	ID                   string             `json:"id"`
	Name                 string             `json:"name"`
	Description          string             `json:"description"`
	IsDefaultAssociation bool               `json:"is_default_association"`
	IsDefaultPropagation bool               `json:"is_default_propagation"`
	State                string             `json:"state"`
	Tags                 []tags.ResourceTag `json:"tags"`
	BgpOptions           *BgpOptions        `json:"bgp_options"`
	CreatedAt            string             `json:"created_at"`
	UpdatedAt            string             `json:"updated_at"`
}
