package template

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type GetInstanceConfOpts struct {
	InstanceId string `json:"-"`
	Offset     int    `q:"offset"`
	Limit      int    `q:"limit"`
}

func GetInstanceConfigurations(client *golangsdk.ServiceClient, opts GetInstanceConfOpts) ([]ParameterValuesInfo, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", opts.InstanceId, "configurations").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return InstanceConfigurationPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractInstanceConfigurations(pages)
}

func ExtractInstanceConfigurations(r pagination.NewPage) ([]ParameterValuesInfo, error) {
	var s struct {
		ParameterValues []ParameterValuesInfo `json:"parameter_values"`
	}
	err := extract.Into(bytes.NewReader((r.(InstanceConfigurationPage)).Body), &s)
	return s.ParameterValues, err
}

type InstanceConfigurationPage struct {
	pagination.NewSinglePageBase
}

type ParameterValuesInfo struct {
	Name            string `json:"name"`
	Value           string `json:"value"`
	RestartRequired bool   `json:"restart_required"`
	Readonly        bool   `json:"readonly"`
	ValueRange      string `json:"value_range"`
	Type            string `json:"type"`
	Description     string `json:"description"`
}

type ParameterConfigurationInfo struct {
	DatastoreVersionName string `json:"datastore_version_name"`
	DatastoreName        string `json:"datastore_name"`
	Created              string `json:"created"`
	Updated              string `json:"updated"`
}
