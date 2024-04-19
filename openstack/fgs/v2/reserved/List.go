package reserved

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Marker  string `q:"marker"`
	Limit   string `q:"limit"`
	FuncUrn string `q:"urn"`
}

func ListReservedInst(client *golangsdk.ServiceClient, opts ListOpts) (*FuncReservedResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("fgs", "functions", "reservedinstances").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res FuncReservedResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
