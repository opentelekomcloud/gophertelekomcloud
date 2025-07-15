package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// ID of the last enterprise router on the previous page. If this parameter is left blank, the first page is queried.
	// This parameter must be used together with limit.
	Marker string `q:"marker"`
	// Number of records on each page. Value range: 0 to 2000
	Limit int `q:"limit"`
	// Enterprise router status. Value options: pending, available, modifying, deleting, deleted, failed and freezed
	State []string `q:"state"`
	// Query by resource ID. Multiple resources can be queried at a time.
	ID []string `q:"id"`
	// Attachment resource IDs
	ResourceID []string `q:"resource_id"`
	// Keyword for sorting. The keyword can be id, name, or state. By default, id is used.
	SortKey []string `q:"sort_key"`
	//
	// Sorting order. There are two value options: asc (ascending order) and desc (descending order).
	SortDir []string `q:"sort_dir"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListRouterInstanceResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("enterprise-router", "instances").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListRouterInstanceResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListRouterInstanceResp struct {
	// Enterprise routers
	Instances []RouterInstance `json:"instances"`
	// Request ID
	RequestID string `json:"request_id"`
	// Pagination query information
	PageInfo PageInfo `json:"page_info"`
}

type PageInfo struct {
	// Marker of the next page. The value is the resource UUID. If the value is empty, the resource is on the last page.
	NextMarker string `json:"next_marker"`
	// Number of resources in the list
	CurrentCount int `json:"current_count"`
}
