package api

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListHistoryOpts struct {
	Offset  int64  `q:"offset"`
	Limit   int    `q:"limit"`
	EnvID   string `q:"env_id"`
	EnvName string `q:"env_name"`
}

func GetHistory(client *golangsdk.ServiceClient, gatewayID, apiID string, opts ListHistoryOpts) ([]VersionResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("apigw", "instances", gatewayID, "apis", "publish", apiID).
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return VersionPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractVersions(pages)
}

type VersionPage struct {
	pagination.NewSinglePageBase
}

func ExtractVersions(r pagination.NewPage) ([]VersionResp, error) {
	var s struct {
		Versions []VersionResp `json:"api_versions"`
	}
	err := extract.Into(bytes.NewReader((r.(VersionPage)).Body), &s)
	return s.Versions, err
}

type VersionResp struct {
	VersionID   string `json:"version_id"`
	Version     string `json:"version_no"`
	ApiID       string `json:"api_id"`
	EnvID       string `json:"env_id"`
	EnvName     string `json:"env_name"`
	Description string `json:"remark"`
	PublishTime string `json:"publish_time"`
	Status      int    `json:"status"`
}
