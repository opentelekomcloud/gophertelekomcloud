package instances

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type QueryEngineOpts struct {
	// Specifies the Index offset.
	// The query starts from the next piece of data indexed by this parameter. The value is 0 by default.
	// The value must be a positive integer.
	Offset int `q:"offset"`
	// A maximum of instances to be queried.
	// Value range: 1 to 128.
	// If the parameter value is not specified, 10 DDM instances are queried by default.
	Limit int `q:"limit"`
}

// QueryEngineInfo is used to query DDM instances.
func QueryEngineInfo(client *golangsdk.ServiceClient, opts QueryEngineOpts) ([]EngineGroupsInfo, error) {
	// GET /v2/{project_id}/engines?offset={offset}&limit={limit}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("engines").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return EnginesPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractEngineGroups(pages)
}

func ExtractEngineGroups(r pagination.NewPage) ([]EngineGroupsInfo, error) {
	var queryEnginesResponse QueryEnginesResponse
	err := extract.Into(bytes.NewReader((r.(EnginesPage)).Body), &queryEnginesResponse)
	return queryEnginesResponse.EngineGroups, err
}

type EnginesPage struct {
	pagination.NewSinglePageBase
}

type QueryEnginesResponse struct {
	// Information of available engines
	EngineGroups []EngineGroupsInfo `json:"engineGroups"`
	// Which page the server starts returning items
	Offset int `json:"offset"`
	// Number of records displayed on each page
	Limit int `json:"limit"`
	// Number of engine versions
	Total int `json:"total"`
}

type EngineGroupsInfo struct {
	// Engine ID
	ID string `json:"id"`
	// Engine name
	Name string `json:"name"`
	// Engine version
	Version string `json:"version"`
	// AZs (Array of SupportAzsInfo objects)
	SupportAzs []SupportAzsInfo `json:"supportAzs"`
}

type SupportAzsInfo struct {
	// AZ code
	Code string `json:"code"`
	// AZ name
	Name string `json:"name"`
	// Whether the current AZ is supported
	Favored bool `json:"favored"`
}
