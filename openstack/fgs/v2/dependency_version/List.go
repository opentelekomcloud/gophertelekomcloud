package dependency_version

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListDependencyVersionOpts struct {
	DependId string `json:"-"`
	Marker   string `q:"marker"`
	MaxItems string `q:"max_items"`
}

func ListDependencyVersions(client *golangsdk.ServiceClient, opts ListDependencyVersionOpts) (*ListDepVersionResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("fgs", "dependencies", opts.DependId, "version").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListDepVersionResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListDepVersionResp struct {
	Dependencies []DepVersionListResp `json:"dependencies"`
	NextMarker   int                  `json:"next_marker"`
	Count        int                  `json:"count"`
}

type DepVersionListResp struct {
	Id           string `json:"id"`
	Owner        string `json:"owner"`
	Link         string `json:"link"`
	Runtime      string `json:"runtime"`
	Etag         string `json:"etag"`
	Size         int    `json:"size"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	FileName     string `json:"file_name"`
	Version      int    `json:"version"`
	DepId        string `json:"dep_id"`
	LastModified string `json:"last_modified"`
	DownloadLink string `json:"download_link"`
	IsShared     bool   `json:"is_shared"`
}
