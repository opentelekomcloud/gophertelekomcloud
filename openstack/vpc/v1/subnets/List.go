package subnets

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	Marker string `q:"marker,omitempty"`
	Limit  *int   `q:"limit,omitempty"`
	VpcID  string `q:"vpc_id,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Subnet, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.ProjectID, "subnets").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return SubnetPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractSubnets(pages)
}

type SubnetPage struct {
	pagination.NewPageResult
}

func (p SubnetPage) NewNextPageURL() (string, error) {
	subnets, err := ExtractSubnets(p)
	if err != nil || len(subnets) == 0 {
		return "", err
	}
	next := p.URL
	query := next.Query()
	query.Set("marker", subnets[len(subnets)-1].ID)
	next.RawQuery = query.Encode()
	return next.String(), nil
}

func (p SubnetPage) NewIsEmpty() (bool, error) {
	subnets, err := ExtractSubnets(p)
	return len(subnets) == 0, err
}

func ExtractSubnets(page pagination.NewPage) ([]Subnet, error) {
	var res struct {
		Subnets []Subnet `json:"subnets"`
	}
	err := extract.Into(bytes.NewReader(page.(SubnetPage).Body), &res)
	return res.Subnets, err
}
