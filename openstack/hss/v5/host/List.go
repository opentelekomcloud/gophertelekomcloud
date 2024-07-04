package hss

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Offset from which the query starts. If the value is less than 0, it is automatically converted to 0.
	Offset *int `q:"offset"`
	// Number of items displayed on each page.
	// A value less than or equal to 0 will be automatically converted to 10,
	// and a value greater than 200 will be automatically converted to 200.
	Limit int `q:"limit"`
	// Access control policy name.
	Name string `q:"group_name"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]HostGroupResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("host-management", "groups").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v5/{project_id}/host-management/groups
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return HostGroupPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
		Headers: map[string]string{"region": client.RegionID},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractGroups(pages)
}

type HostGroupPage struct {
	pagination.NewSinglePageBase
}

func ExtractGroups(r pagination.NewPage) ([]HostGroupResp, error) {
	var s struct {
		Groups []HostGroupResp `json:"data_list"`
	}
	err := extract.Into(bytes.NewReader((r.(HostGroupPage)).Body), &s)
	return s.Groups, err
}

type HostGroupResp struct {
	ID               string   `json:"group_id"`
	Name             string   `json:"group_name"`
	HostIds          []string `json:"host_id_list"`
	HostNum          int      `json:"host_num"`
	RiskHostNum      int      `json:"risk_host_num"`
	UnprotectHostNum int      `json:"unprotect_host_num"`
}
