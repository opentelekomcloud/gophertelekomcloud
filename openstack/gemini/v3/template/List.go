package template

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	Offset int `q:"offset"`
	Limit  int `q:"limit"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Configuration, error) {
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

func ExtractConfigurations(r pagination.NewPage) ([]Configuration, error) {
	var s struct {
		Configurations []Configuration `json:"configurations"`
	}
	err := extract.Into(bytes.NewReader((r.(ConfigurationPage)).Body), &s)
	return s.Configurations, err
}

type ConfigurationPage struct {
	pagination.NewSinglePageBase
}

type Configuration struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	DataStoreVersionName string `json:"datastore_version_name"`
	DataStoreName        string `json:"datastore_name"`
	Created              string `json:"created"`
	Updated              string `json:"updated"`
	Mode                 string `json:"mode"`
	UserDefined          bool   `json:"user_defined"`
}
