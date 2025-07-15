package network

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	Namespace            string `json:"-"`
	AllowWatchBookmarks  *bool  `q:"allowWatchBookmarks,omitempty"`
	Continue             string `q:"continue,omitempty"`
	FieldSelector        string `q:"fieldSelector,omitempty"`
	LabelSelector        string `q:"labelSelector,omitempty"`
	Limit                *int   `q:"limit,omitempty"`
	ResourceVersion      string `q:"resourceVersion,omitempty"`
	ResourceVersionMatch string `q:"resourceVersionMatch,omitempty"`
	SendInitialEvents    *bool  `q:"sendInitialEvents,omitempty"`
	TimeoutSeconds       *int   `q:"timeoutSeconds,omitempty"`
	Watch                *bool  `q:"watch,omitempty"`
	Pretty               *bool  `q:"pretty,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]NetworkResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("namespaces", opts.Namespace, "networks").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return NetworksPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}

	return ExtractNetworks(pages)
}

type NetworksPage struct {
	pagination.NewSinglePageBase
}

func ExtractNetworks(r pagination.NewPage) ([]NetworkResp, error) {
	var s struct {
		Items []NetworkResp `json:"items"`
	}
	err := extract.Into(bytes.NewReader((r.(NetworksPage)).Body), &s)
	return s.Items, err
}
