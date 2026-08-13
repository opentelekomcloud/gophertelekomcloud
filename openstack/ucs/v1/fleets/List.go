package fleets

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Limit   int    `q:"limit"`
	Offset  int    `q:"offset"`
	OrderBy string `q:"order_by"`
	Order   string `q:"order"`
}

type ListResponse struct {
	Items []ClusterGroup `json:"items"`
	Total int            `json:"total"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("clustergroups").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	return &res, extract.Into(raw.Body, &res)
}
