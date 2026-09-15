package routetables

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name        string       `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	Routes      *RouteAction `json:"routes,omitempty"`
}

type RouteAction struct {
	Add []RouteOpts       `json:"add,omitempty"`
	Mod []RouteOpts       `json:"mod,omitempty"`
	Del []DeleteRouteOpts `json:"del,omitempty"`
}

type DeleteRouteOpts struct {
	Type        string  `json:"type,omitempty"`
	Destination string  `json:"destination" required:"true"`
	NextHop     string  `json:"nexthop,omitempty"`
	Description *string `json:"description,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*RouteTable, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(client.ProjectID, "routetables", id).Build()
	if err != nil {
		return nil, err
	}
	b, err := build.RequestBody(opts, "routetable")
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
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
