package reserved

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListConfigOpts struct {
	Marker  string `q:"marker"`
	Limit   string `q:"limit"`
	FuncUrn string `q:"urn"`
}

func ListReservedInstConfigs(client *golangsdk.ServiceClient, opts ListConfigOpts) (*FuncReservedResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("fgs", "functions", "reservedinstanceconfigs").WithQueryParams(&opts).Build()
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

type FuncReservedResp struct {
	ReservedInstances []FuncReservedInstances `json:"reserved_instances"`
	PageInfo          *PageInfo               `json:"page_info"`
	Count             int                     `json:"count"`
}

type FuncReservedInstances struct {
	FuncUrn string `json:"func_urn"`
	Count   int    `json:"count"`
}

type PageInfo struct {
	NextMarker     int `json:"next_marker"`
	PreviousMarker int `json:"previous_marker"`
	CurrentCount   int `json:"current_count"`
}
