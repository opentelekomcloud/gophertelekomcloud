package volumes

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts contains options for filtering, sorting, and paginating volumes.
type ListOpts struct {
	// AllTenants lists volumes for all projects. This option requires
	// administrator privileges.
	AllTenants bool `q:"all_tenants"`

	// TenantID filters volumes by project ID.
	TenantID string `q:"project_id"`

	// Name filters volumes by name. The maximum length is 255 bytes.
	Name string `q:"name"`

	// Status filters volumes by status.
	Status string `q:"status"`

	// Metadata filters volumes by metadata key-value pairs.
	Metadata map[string]string `q:"metadata"`

	// AvailabilityZone filters volumes by availability zone.
	AvailabilityZone string `q:"availability_zone"`

	// Sort is a comma-separated list of sort keys and optional sort directions
	// in the form <key>[:<direction>].
	Sort string `q:"sort"`

	// SortKey specifies the field used to sort results. Supported values are
	// id, status, size, and created_at.
	SortKey string `q:"sort_key"`

	// SortDir specifies the sort direction. Supported values are asc and desc.
	SortDir string `q:"sort_dir"`

	// Limit is the maximum number of results returned per page. Supported
	// values range from 1 to 1000.
	Limit int `q:"limit"`

	// Offset skips this number of results.
	Offset int `q:"offset"`

	// Marker is the ID of the last volume on the previous page.
	Marker string `q:"marker"`
}

// ListResponse contains volumes and pagination links.
type ListResponse struct {
	Volumes      []Volume         `json:"volumes"`
	VolumesLinks []golangsdk.Link `json:"volumes_links"`
}

// List returns all detailed volumes matching the provided options.
func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("volumes", "detail") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return VolumePage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}

	return extractVolumes(pages)
}

type VolumePage struct {
	pagination.NewPageResult
}

func (p VolumePage) NewNextPageURL() (string, error) {
	var res ListResponse
	if err := extract.Into(bytes.NewReader(p.Body), &res); err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(res.VolumesLinks)
}

func (p VolumePage) NewIsEmpty() (bool, error) {
	res, err := extractVolumes(p)
	if err != nil {
		return false, err
	}
	return len(res.Volumes) == 0, nil
}

func extractVolumes(page pagination.NewPage) (*ListResponse, error) {
	var res ListResponse
	if err := extract.Into(bytes.NewReader(page.(VolumePage).Body), &res); err != nil {
		return nil, err
	}
	if res.Volumes == nil {
		res.Volumes = []Volume{}
	}
	return &res, nil
}
