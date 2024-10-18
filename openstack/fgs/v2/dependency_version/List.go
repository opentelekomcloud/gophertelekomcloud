package dependency_version

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	DependId string `json:"-"`
	Marker   string `q:"marker"`
	MaxItems string `q:"max_items"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListDepVersionResp, error) {
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
	Dependencies []DepVersionResp `json:"dependencies"`
	NextMarker   int              `json:"next_marker"`
	Count        int              `json:"count"`
}
