package channel

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListMembersOpts struct {
	GatewayID string `json:"-"`
	ChannelID string `json:"-"`
	// Offset from which the query starts. If the value is less than 0, it is automatically converted to 0.
	Offset string `q:"offset"`
	// Number of items displayed on each page.
	// A value less than or equal to 0 will be automatically converted to 20,
	// and a value greater than 500 will be automatically converted to 500.
	Limit int `q:"limit"`
	// Cloud server name.
	Name string `q:"name"`
	// Backend server group name.
	MemberGroupName string `q:"member_group_name"`
	// Backend server group ID.
	MemberGroupId string `q:"member_group_id"`
	// Parameter name for exact matching. Separate multiple parameter names with commas (,).
	// Currently, name and member_group_name are supported.
	PreciseSearch string `q:"precise_search"`
}

func ListMembers(client *golangsdk.ServiceClient, opts ListMembersOpts) ([]MemberResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("apigw", "instances", opts.GatewayID, "vpc-channels", opts.ChannelID, "members").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return MemberPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractMembers(pages)
}

type MemberPage struct {
	pagination.NewSinglePageBase
}

func ExtractMembers(r pagination.NewPage) ([]MemberResp, error) {
	var s struct {
		Members []MemberResp `json:"members"`
	}
	err := extract.Into(bytes.NewReader((r.(MemberPage)).Body), &s)
	return s.Members, err
}
