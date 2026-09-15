package routetables

import (
	"bytes"
	"fmt"
	"strconv"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

const defaultPageLimit = 2000

type ListOpts struct {
	Limit    *int   `q:"limit,omitempty"`
	Marker   string `q:"marker,omitempty"`
	ID       string `q:"id,omitempty"`
	VpcID    string `q:"vpc_id,omitempty"`
	SubnetID string `q:"subnet_id,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]RouteTable, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(client.ProjectID, "routetables").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return RouteTablePage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractRouteTables(pages)
}

type RouteTablePage struct {
	pagination.NewPageResult
}

func (p RouteTablePage) NewNextPageURL() (string, error) {
	tables, err := ExtractRouteTables(p)
	if err != nil || len(tables) == 0 {
		return "", err
	}
	limit := defaultPageLimit
	if value := p.URL.Query().Get("limit"); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil {
			return "", fmt.Errorf("invalid route table page limit %q: %w", value, err)
		}
	}
	if limit > 0 && len(tables) < limit {
		return "", nil
	}
	next := p.URL
	query := next.Query()
	query.Set("marker", tables[len(tables)-1].ID)
	next.RawQuery = query.Encode()
	return next.String(), nil
}

func (p RouteTablePage) NewIsEmpty() (bool, error) {
	tables, err := ExtractRouteTables(p)
	return len(tables) == 0, err
}

func ExtractRouteTables(page pagination.NewPage) ([]RouteTable, error) {
	var res struct {
		RouteTables []RouteTable `json:"routetables"`
	}
	err := extract.Into(bytes.NewReader(page.(RouteTablePage).Body), &res)
	return res.RouteTables, err
}
