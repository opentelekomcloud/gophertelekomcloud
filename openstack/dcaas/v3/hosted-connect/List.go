package hosted_connect

import (
	"bytes"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// virtual gateway by ID
	ID []string `q:"id,omitempty"`
	// Specifies the number of records returned on each page. Value range: 1-2000
	Limit int `q:"limit,omitempty"`
	// Specifies the ID of the last resource record on the previous page. If this parameter is left blank, the first page is queried.
	// This parameter must be used together with limit.
	Marker string `q:"marker,omitempty"`
	// Specifies the list of fields to be displayed.
	Fields []interface{} `q:"fields,omitempty"`
	// Specifies the sorting order of returned results. The value can be asc (default) or desc.
	SortDir string `q:"sort_dir,omitempty"`
	// Specifies the field for sorting.
	SortKey string `q:"sort_key,omitempty"`
	// Specifies operations connection ID by which hosted connections are queried.
	HostingId []string `q:"hosting_id,omitempty"`
	// Specifies the resource name by which instances are queried. You can specify multiple names.
	Name []string `q:"name,omitempty"`
}

// List is used to obtain the virtual gateway list
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]HostedConnect, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("dcaas", "hosted-connects").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return HcPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractHCs(pages)
}

type HcPage struct {
	pagination.NewSinglePageBase
}

func ExtractHCs(r pagination.NewPage) ([]HostedConnect, error) {
	var s struct {
		Connects []HostedConnect `json:"hosted_connects"`
	}
	err := extract.Into(bytes.NewReader((r.(HcPage)).Body), &s)
	return s.Connects, err
}
