package volumetypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Comma-separated list of sort keys and optional sort directions in the
	// form of <key>[:<direction>].
	Sort string `q:"sort"`
	// Requests a page size of items.
	Limit int `q:"limit"`
	// Used in conjunction with limit to return a slice of items.
	Offset int `q:"offset"`
	// The ID of the last-seen item.
	Marker string `q:"marker"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("types")+query.String(), func(r pagination.PageResult) pagination.Page {
		return VolumeTypePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

type VolumeTypePage struct {
	pagination.LinkedPageBase
}

func (r VolumeTypePage) IsEmpty() (bool, error) {
	volumeTypes, err := ExtractVolumeTypes(r)
	return len(volumeTypes) == 0, err
}

func (r VolumeTypePage) NextPageURL() (string, error) {
	var s struct {
		Links []golangsdk.Link `json:"volume_type_links"`
	}

	err := extract.Into(r.Body, &s)
	if err != nil {
		return "", err
	}

	return golangsdk.ExtractNextURL(s.Links)
}

func ExtractVolumeTypes(r pagination.Page) ([]VolumeType, error) {
	var res struct {
		VolumeTypes []VolumeType `json:"volume_types"`
	}
	err := extract.Into(r.(VolumeTypePage).Result, &res)
	return res.VolumeTypes, err
}
