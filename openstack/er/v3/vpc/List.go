package vpc

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
	// Attachment status
	State []string `q:"state"`
	// Query by resource ID. Multiple resources can be queried at a time.
	ID []string `q:"id"`
	// Keyword for sorting. The keyword can be id, name, or state. By default, id is used.
	SortKey []string `q:"sort_key"`
	// Sorting order
	SortDir []string `q:"sort_dir"`
	// Vpc ID
	VpcId string `q:"vpc_id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListVpcAttachmentDetails, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("enterprise-router", "instances").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListVpcAttachmentDetails
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListVpcAttachmentDetails struct {
	VpcAttachments []VpcAttachmentDetails `json:"vpc_attachments"`
	PageInfo       *PageInfo              `json:"page_info"`
	RequestId      string                 `json:"request_id"`
}

type PageInfo struct {
	NextMarker   string `json:"next_marker"`
	CurrentCount int    `json:"current_count"`
}
