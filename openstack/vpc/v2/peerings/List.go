package peerings

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	ID       string `q:"id,omitempty"`
	Name     string `q:"name,omitempty"`
	Status   string `q:"status,omitempty"`
	TenantID string `q:"tenant_id,omitempty"`
	VpcID    string `q:"vpc_id,omitempty"`
	Marker   string `q:"marker,omitempty"`
	Limit    *int   `q:"limit,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Peering, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("vpc", "peerings").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return PeeringPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractPeerings(pages)
}

type PeeringPage struct {
	pagination.NewPageResult
}

func (p PeeringPage) NewNextPageURL() (string, error) {
	var res struct {
		Links []golangsdk.Link `json:"peerings_links"`
	}
	if err := extract.Into(bytes.NewReader(p.Body), &res); err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(res.Links)
}

func (p PeeringPage) NewIsEmpty() (bool, error) {
	peerings, err := ExtractPeerings(p)
	return len(peerings) == 0, err
}

func ExtractPeerings(page pagination.NewPage) ([]Peering, error) {
	var res struct {
		Peerings []Peering `json:"peerings"`
	}
	err := extract.Into(bytes.NewReader(page.(PeeringPage).Body), &res)
	return res.Peerings, err
}
