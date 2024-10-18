package instances

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type QueryParametersOpts struct {
	// Specifies the Index offset.
	// The query starts from the next piece of data indexed by this parameter. The value is 0 by default.
	// The value must be a positive integer.
	Offset int `q:"offset"`
	// A maximum of parameters  to be queried.
	// Value range: 1 to 128.
	// If the parameter value is not specified, 10 parameters are queried by default.
	Limit int `q:"limit"`
}

// This function is used to query parameters of a specified DDM instance.
func QueryParameters(client *golangsdk.ServiceClient, instanceId string, opts QueryParametersOpts) ([]ConfigurationParameterList, error) {
	// GET /v3/{project_id}/instances/{instance_id}/configurations?offset={offset}&limit={limit}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceId, "configurations").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ParametersPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractParameters(pages)
}

func ExtractParameters(r pagination.NewPage) ([]ConfigurationParameterList, error) {
	var queryParametersResponse QueryParametersResponse
	err := extract.Into(bytes.NewReader((r.(ParametersPage)).Body), &queryParametersResponse)
	return queryParametersResponse.ConfigurationParameter, err
}

type ParametersPage struct {
	pagination.NewSinglePageBase
}

type QueryParametersResponse struct {
	// Time when DDM instance parameters are last updated
	Updated string `json:"updated"`
	// Information about DDM instance parameters
	ConfigurationParameter []ConfigurationParameterList `json:"configuration_parameter"`
	// Which page the server starts returning items.
	Offset int `json:"offset"`
	// Number of records displayed on each page
	Limit int `json:"limit"`
	// Total collections
	Total int `json:"total"`
}

type ConfigurationParameterList struct {
	// Parameter name
	Name string `json:"name"`
	// Parameter value
	Value string `json:"value"`
	// Whether the instance needs to be restarted
	NeedRestart string `json:"need_restart"`
	// Whether the parameter is read-only
	ReadOnly string `json:"read_only"`
	// Parameter value range
	ValueRange string `json:"value_range"`
	// Parameter type
	DataType string `json:"data_type"`
	// Parameter description
	Description string `json:"description"`
}
