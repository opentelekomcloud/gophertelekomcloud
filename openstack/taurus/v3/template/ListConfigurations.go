package template

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListConfigurationsOpts struct {
	Offset int `q:"offset"`
	Limit  int `q:"limit"`
}

func ListConfigurations(client *golangsdk.ServiceClient, opts ListConfigurationsOpts) ([]ConfigurationSummary, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("configurations").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ConfigurationPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractConfigurations(pages)
}

func ExtractConfigurations(r pagination.NewPage) ([]ConfigurationSummary, error) {
	var s struct {
		Configurations []ConfigurationSummary `json:"configurations"`
	}
	err := extract.Into(bytes.NewReader((r.(ConfigurationPage)).Body), &s)
	return s.Configurations, err
}

type ConfigurationPage struct {
	pagination.NewSinglePageBase
}

type ConfigurationSummary struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	DatastoreVersionName string `json:"datastore_version_name"`
	DatastoreName        string `json:"datastore_name"`
	Created              string `json:"created"`
	Updated              string `json:"updated"`
	UserDefined          bool   `json:"user_defined"`
}
