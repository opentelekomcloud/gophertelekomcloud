package privateips

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
	Marker string `q:"marker,omitempty"`
	Limit  *int   `q:"limit,omitempty"`
}

func List(client *golangsdk.ServiceClient, subnetID string, opts ListOpts) ([]PrivateIP, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.ProjectID, "subnets", subnetID, "privateips").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return PrivateIPPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractPrivateIPs(pages)
}

type PrivateIPPage struct {
	pagination.NewPageResult
}

func (p PrivateIPPage) NewNextPageURL() (string, error) {
	privateIPs, err := ExtractPrivateIPs(p)
	if err != nil || len(privateIPs) == 0 {
		return "", err
	}

	limit := defaultPageLimit
	if value := p.URL.Query().Get("limit"); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil {
			return "", fmt.Errorf("invalid private IP page limit %q: %w", value, err)
		}
	}
	if limit > 0 && len(privateIPs) < limit {
		return "", nil
	}

	next := p.URL
	query := next.Query()
	query.Set("marker", privateIPs[len(privateIPs)-1].ID)
	next.RawQuery = query.Encode()
	return next.String(), nil
}

func (p PrivateIPPage) NewIsEmpty() (bool, error) {
	privateIPs, err := ExtractPrivateIPs(p)
	return len(privateIPs) == 0, err
}

func ExtractPrivateIPs(page pagination.NewPage) ([]PrivateIP, error) {
	var res struct {
		PrivateIPs []PrivateIP `json:"privateips"`
	}
	err := extract.Into(bytes.NewReader(page.(PrivateIPPage).Body), &res)
	return res.PrivateIPs, err
}
