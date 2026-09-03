package bandwidths

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Specifies the resource ID of pagination query. If the parameter is
	// left blank, only resources on the first page are queried.
	Marker string `q:"marker,omitempty"`

	// Specifies the number of records returned on each page.
	Limit int `q:"limit,omitempty"`

	// Specifies the enterprise project ID. This field can be used to filter
	// bandwidths under an enterprise project.
	EnterpriseProjectId string `q:"enterprise_project_id,omitempty"`
}

// List retrieves bandwidths matching the given search criteria.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]BandWidth, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.ProjectID, "bandwidths").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return BandWidthPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractBandWidths(pages)
}

type BandWidthPage struct {
	pagination.NewPageResult
}

func (p BandWidthPage) NewNextPageURL() (string, error) {
	bandwidths, err := ExtractBandWidths(p)
	if err != nil || len(bandwidths) == 0 {
		return "", err
	}
	next := p.URL
	query := next.Query()
	query.Set("marker", bandwidths[len(bandwidths)-1].ID)
	next.RawQuery = query.Encode()
	return next.String(), nil
}

func (p BandWidthPage) NewIsEmpty() (bool, error) {
	bandwidths, err := ExtractBandWidths(p)
	return len(bandwidths) == 0, err
}

func ExtractBandWidths(page pagination.NewPage) ([]BandWidth, error) {
	var res struct {
		BandWidths []BandWidth `json:"bandwidths"`
	}
	err := extract.Into(bytes.NewReader(page.(BandWidthPage).Body), &res)
	return res.BandWidths, err
}
