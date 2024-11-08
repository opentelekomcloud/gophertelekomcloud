package connection

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Specifies the number of records returned on each page during pagination query.
	Limit *int `q:"limit"`
	// Specifies the start flag for querying the current page. If this parameter is left blank, the first page is queried.
	// The marker for querying the next page is the next_marker in the page_info object returned on the current page.
	Marker string `q:"marker"`
	// Specifies an EIP ID or private IP address of a VPN gateway.
	VgwIp string `q:"vgw_ip"`
	// Specifies a VPN gateway ID.
	VgwId string `q:"vgw_id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Connection, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("vpn-connection").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ConnPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractConnections(pages)
}

type ConnPage struct {
	pagination.NewSinglePageBase
}

func ExtractConnections(r pagination.NewPage) ([]Connection, error) {
	var s struct {
		Connections []Connection `json:"vpn_connections"`
	}
	err := extract.Into(bytes.NewReader((r.(ConnPage)).Body), &s)
	return s.Connections, err
}
