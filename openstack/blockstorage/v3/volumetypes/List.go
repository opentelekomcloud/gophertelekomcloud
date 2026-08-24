package volumetypes

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// IsPublic filters volume types by public visibility.
	IsPublic *bool `q:"is_public"`
	// Sort is a comma-separated list of sort keys and optional sort directions
	// in the form <key>[:<direction>].
	Sort string `q:"sort"`
	// SortKey specifies the field used to sort results.
	SortKey string `q:"sort_key"`
	// SortDir specifies the sort direction.
	SortDir string `q:"sort_dir"`
	// Limit requests a page size of items.
	Limit int `q:"limit"`
	// Offset is where to start in the list.
	Offset int `q:"offset"`
	// Marker is the ID of the last-seen item.
	Marker string `q:"marker"`
}

type listResponse struct {
	VolumeTypes     []VolumeType     `json:"volume_types"`
	VolumeTypeLinks []golangsdk.Link `json:"volume_type_links"`
}

// List returns all volume types matching the provided options.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]VolumeType, error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("types") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return VolumeTypePage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}

	return extractVolumeTypes(pages)
}

type VolumeTypePage struct {
	pagination.NewPageResult
}

func (p VolumeTypePage) NewNextPageURL() (string, error) {
	var res listResponse
	if err := extract.Into(bytes.NewReader(p.Body), &res); err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(res.VolumeTypeLinks)
}

func (p VolumeTypePage) NewIsEmpty() (bool, error) {
	res, err := extractVolumeTypes(p)
	if err != nil {
		return false, err
	}
	return len(res) == 0, nil
}

func extractVolumeTypes(page pagination.NewPage) ([]VolumeType, error) {
	var res listResponse
	if err := extract.Into(bytes.NewReader(page.(VolumeTypePage).Body), &res); err != nil {
		return nil, err
	}
	if res.VolumeTypes == nil {
		res.VolumeTypes = []VolumeType{}
	}
	return res.VolumeTypes, nil
}
