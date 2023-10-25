package routetables

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// The VPC ID that the route table belongs to
	VpcID string `json:"vpc_id" required:"true"`
	// The name of the route table. The name can contain a maximum of 64 characters,
	// which may consist of letters, digits, underscores (_), hyphens (-), and periods (.).
	// The name cannot contain spaces.
	Name string `json:"name" required:"true"`
	// The supplementary information about the route table. The description can contain
	// a maximum of 255 characters and cannot contain angle brackets (< or >).
	Description string `json:"description,omitempty"`
	// The route information
	Routes []RouteOpts `json:"routes,omitempty"`
}

// RouteOpts contains all the values needed to manage a vpc route
type RouteOpts struct {
	// The destination CIDR block. The destination of each route must be unique.
	// The destination cannot overlap with any subnet CIDR block in the VPC.
	Destination string `json:"destination" required:"true"`
	// the type of the next hop. value range:
	// ecs, eni, vip, nat, peering, vpn, dc, cc, egw
	Type string `json:"type" required:"true"`
	// the instance ID of next hop
	NextHop string `json:"nexthop" required:"true"`
	// The supplementary information about the route. The description can contain
	// a maximum of 255 characters and cannot contain angle brackets (< or >).
	Description *string `json:"description,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*RouteTable, error) {
	b, err := build.RequestBody(opts, "routetable")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(client.ProjectID, "routetables"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res RouteTable
	err = extract.IntoStructPtr(raw.Body, &res, "routetable")
	return &res, err
}

// RouteTable represents a route table
type RouteTable struct {
	// Name is the human-readable name for the route table
	Name string `json:"name"`
	// Description is the supplementary information about the route table
	Description string `json:"description"`
	// ID is the unique identifier for the route table
	ID string `json:"id"`
	// the VPC ID that the route table belongs to.
	VpcID string `json:"vpc_id"`
	// project id
	TenantID string `json:"tenant_id"`
	// Default indicates whether it is a default route table
	Default bool `json:"default"`
	// Routes is an array of static routes that the route table will host
	Routes []Route `json:"routes"`
	// Subnets is an array of subnets that associated with the route table
	Subnets []Subnet `json:"subnets"`
	// Specifies the time (UTC) when the route table is created.
	CreatedAt string `json:"created_at"`
	// Specifies the time (UTC) when the route table is updated.
	UpdatedAt string `json:"updated_at"`
}

// Route represents a route object in a route table
type Route struct {
	Type            string `json:"type"`
	DestinationCIDR string `json:"destination"`
	NextHop         string `json:"nexthop"`
	Description     string `json:"description"`
}

// Subnet represents a subnet object associated with a route table
type Subnet struct {
	ID string `json:"id"`
}
