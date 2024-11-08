package route

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListStaticOpts struct {
	RouteTableId string   `json:"-"`
	Marker       string   `q:"marker"`
	Limit        int      `q:"limit"`
	Destination  []string `q:"destination"`
	AttachmentId []string `q:"attachment_id"`
	ResourceType []string `q:"resource_type"`
	SortKey      []string `q:"sort_key"`
	SortDir      []string `q:"sort_dir"`
}

func ListStatic(client *golangsdk.ServiceClient, opts ListStaticOpts) (*ListStaticRoutes, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("enterprise-router", "route-tables", opts.RouteTableId, "static-routes").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListStaticRoutes
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListStaticRoutes struct {
	Routes    []Route   `json:"routes"`
	PageInfo  *PageInfo `json:"page_info"`
	RequestId string    `json:"request_id"`
}

type PageInfo struct {
	NextMarker   string `json:"next_marker"`
	CurrentCount int    `json:"current_count"`
}
