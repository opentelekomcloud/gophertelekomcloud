package route

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	RouteTableId string   `json:"-"`
	Marker       string   `q:"marker"`
	Limit        int      `q:"limit"`
	Destination  []string `q:"destination"`
	ResourceType []string `q:"resource_type"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListRoutes, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("enterprise-router", "route-tables", opts.RouteTableId, "routes").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListRoutes
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListRoutes struct {
	Routes    []EffectiveRoute `json:"routes"`
	PageInfo  *PageInfo        `json:"page_info"`
	RequestId string           `json:"request_id"`
}

type EffectiveRoute struct {
	RouteId        string            `json:"route_id"`
	Destination    string            `json:"destination"`
	NextHops       []RouteAttachment `json:"next_hops"`
	IsBlackhole    bool              `json:"is_blackhole"`
	RouteType      string            `json:"route_type"`
	AddressGroupId string            `json:"address_group_id"`
}
