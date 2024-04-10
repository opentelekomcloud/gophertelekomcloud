package function

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Marker      string `q:"marker"`
	MaxItems    string `q:"max_items"`
	PackageName string `q:"package_name"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListFuncResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("fgs", "functions").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListFuncResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListFuncResponse struct {
	Functions  []FuncGraph `json:"functions"`
	NextMarker int         `json:"next_marker"`
	Count      int         `json:"count"`
}
