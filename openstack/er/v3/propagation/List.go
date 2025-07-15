package propagation

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	RouterId     string   `json:"-"`
	RouteTableId string   `json:"-"`
	Limit        int      `q:"limit"`
	Marker       string   `q:"marker"`
	AttachmentId []string `q:"attachment_id"`
	ResourceType []string `q:"resource_type"`
	State        []string `q:"state"`
	SortKey      []string `q:"sort_key"`
	SortDir      []string `q:"sort_dir"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListAssociations, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("enterprise-router", opts.RouterId, "route-tables", opts.RouteTableId, "propagations").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListAssociations
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListAssociations struct {
	Propagations []Propagation `json:"propagations"`
	PageInfo     *PageInfo     `json:"page_info"`
	RequestId    string        `json:"request_id"`
}

type PageInfo struct {
	NextMarker   string `json:"next_marker"`
	CurrentCount int    `json:"current_count"`
}
