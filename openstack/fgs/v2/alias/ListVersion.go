package alias

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
)

type ListVersionOpts struct {
	FuncUrn  string `q:"-"`
	Marker   string `q:"marker,omitempty"`
	Maxitems string `q:"maxitems,omitempty"`
}

func ListVersion(client *golangsdk.ServiceClient, opts ListVersionOpts) (*ListVersionResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("fgs", "functions", opts.FuncUrn, "versions").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListVersionResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListVersionResponse struct {
	Functions  []function.FuncGraph `json:"versions"`
	NextMarker int                  `json:"next_marker"`
	Count      int                  `json:"count"`
}
