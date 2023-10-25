package routetables

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ActionSubnetsOpts contains the subnets list that associate or disassociate with a route table
type ActionSubnetsOpts struct {
	Associate    []string `json:"associate,omitempty"`
	Disassociate []string `json:"disassociate,omitempty"`
}

// ActionOpts contains the values used when associating or disassociating subnets with a route table
type ActionOpts struct {
	Subnets ActionSubnetsOpts `json:"subnets" required:"true"`
}

// Action will associate or disassociate subnets with a particular route table based on its unique ID
func Action(client *golangsdk.ServiceClient, id string, opts ActionOpts) (*RouteTable, error) {
	b, err := build.RequestBody(opts, "routetable")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL(client.ProjectID, "routetables", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res RouteTable
	err = extract.IntoStructPtr(raw.Body, &res, "routetable")
	return &res, err
}
