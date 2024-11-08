package route_table

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	RouterId string   `json:"-"`
	Marker   string   `q:"marker"`
	Limit    int      `q:"limit"`
	State    []string `q:"state"`
	SortKey  []string `q:"sort_key"`
	SortDir  []string `q:"sort_dir"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListRouteTables, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("enterprise-router", opts.RouterId, "route-tables").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListRouteTables
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListRouteTables struct {
	RouteTables []RouteTable `json:"route_tables"`
	PageInfo    *PageInfo    `json:"page_info"`
	RequestId   string       `json:"request_id"`
}

type PageInfo struct {
	NextMarker   string `json:"next_marker"`
	CurrentCount int    `json:"current_count"`
}
