package vpcs

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	ID                  string `q:"id,omitempty"`
	Marker              string `q:"marker,omitempty"`
	Limit               *int   `q:"limit,omitempty"`
	EnterpriseProjectID string `q:"enterprise_project_id,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Vpc, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.ProjectID, "vpcs").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return VpcPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractVpcs(pages)
}

type VpcPage struct {
	pagination.NewPageResult
}

func (p VpcPage) NewNextPageURL() (string, error) {
	vpcs, err := ExtractVpcs(p)
	if err != nil || len(vpcs) == 0 {
		return "", err
	}
	next := p.URL
	query := next.Query()
	query.Set("marker", vpcs[len(vpcs)-1].ID)
	next.RawQuery = query.Encode()
	return next.String(), nil
}

func (p VpcPage) NewIsEmpty() (bool, error) {
	vpcs, err := ExtractVpcs(p)
	return len(vpcs) == 0, err
}

func ExtractVpcs(page pagination.NewPage) ([]Vpc, error) {
	var res struct {
		VPCs []Vpc `json:"vpcs"`
	}
	err := extract.Into(bytes.NewReader(page.(VpcPage).Body), &res)
	return res.VPCs, err
}
