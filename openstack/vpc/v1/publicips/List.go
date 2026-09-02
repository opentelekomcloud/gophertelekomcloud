package publicips

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
	Limit  int    `q:"limit,omitempty"`
	// IPVersion is not documented as a list query parameter in the reviewed
	// OTC documentation.
	IPVersion           int    `q:"ip_version,omitempty"`
	EnterpriseProjectId string `q:"enterprise_project_id,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]PublicIP, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.ProjectID, "publicips").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return PublicIPPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractPublicIPs(pages)
}

type PublicIPPage struct {
	pagination.NewPageResult
}

func (p PublicIPPage) NewNextPageURL() (string, error) {
	publicIPs, err := ExtractPublicIPs(p)
	if err != nil || len(publicIPs) == 0 {
		return "", err
	}
	limit := defaultPageLimit
	if value := p.URL.Query().Get("limit"); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil {
			return "", fmt.Errorf("invalid public IP page limit %q: %w", value, err)
		}
	}
	if limit > 0 && len(publicIPs) < limit {
		return "", nil
	}
	next := p.URL
	query := next.Query()
	query.Set("marker", publicIPs[len(publicIPs)-1].ID)
	next.RawQuery = query.Encode()
	return next.String(), nil
}

func (p PublicIPPage) NewIsEmpty() (bool, error) {
	publicIPs, err := ExtractPublicIPs(p)
	return len(publicIPs) == 0, err
}

func ExtractPublicIPs(page pagination.NewPage) ([]PublicIP, error) {
	var res struct {
		PublicIPs []PublicIP `json:"publicips"`
	}
	err := extract.Into(bytes.NewReader(page.(PublicIPPage).Body), &res)
	return res.PublicIPs, err
}
