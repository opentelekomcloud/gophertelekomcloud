package channel

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	GatewayID string `json:"-"`
	// VPC channel ID.
	ID string `q:"id"`
	// Offset from which the query starts. If the value is less than 0, it is automatically converted to 20.
	Offset *int `q:"offset"`
	// Number of items displayed on each page.
	// A value less than or equal to 0 will be automatically converted to 20,
	// and a value greater than 500 will be automatically converted to 500.
	Limit int `q:"limit"`
	// VPC channel name.
	Name string `q:"name"`
	// Dictionary code of the VPC channel.
	// The value can contain letters, digits, hyphens (-), underscores (_), and periods (.).
	// This parameter is currently not supported.
	Code string `q:"dict_code"`
	// Parameter name for exact matching. Separate multiple parameter names with commas (,).
	// Currently, name and member_group_name are supported.
	PreciseSearch string `q:"precise_search"`
	// Backend service address. By default, exact match is used. Fuzzy match is not supported.
	MemberHost string `q:"member_host"`
	// Backend server port.
	MemberPort *int `q:"member_port"`
	// Backend server group name.
	MemberGroupName string `q:"member_group_name"`
	// Backend server group ID.
	MemberGroupId string `q:"member_group_id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]ChannelResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("apigw", "instances", opts.GatewayID, "vpc-channels").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ChannelPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractChannels(pages)
}

type ChannelPage struct {
	pagination.NewSinglePageBase
}

func ExtractChannels(r pagination.NewPage) ([]ChannelResp, error) {
	var s struct {
		Channels []ChannelResp `json:"vpc_channels"`
	}
	err := extract.Into(bytes.NewReader((r.(ChannelPage)).Body), &s)
	return s.Channels, err
}
