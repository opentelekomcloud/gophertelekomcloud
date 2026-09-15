package routetables

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type DisassociateOpts struct {
	Subnets []string `json:"disassociate" required:"true"`
}

func Disassociate(client *golangsdk.ServiceClient, id string, opts DisassociateOpts) (*RouteTable, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(client.ProjectID, "routetables", id, "action").Build()
	if err != nil {
		return nil, err
	}
	b, err := build.RequestBody(struct {
		Subnets DisassociateOpts `json:"subnets" required:"true"`
	}{Subnets: opts}, "routetable")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}
	var res struct {
		RouteTable RouteTable `json:"routetable"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.RouteTable, nil
}
