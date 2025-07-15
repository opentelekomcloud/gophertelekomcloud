package channel

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateMemberOpts struct {
	GatewayID string `json:"-"`
	ChannelID string `json:"-"`
	// Backend instances of the VPC channel.
	Members []Members `json:"members,omitempty"`
	// Backend server group to be modified.
	MemberGroupName string `json:"member_group_name,omitempty"`
}

func UpdateMembers(client *golangsdk.ServiceClient, opts UpdateMemberOpts) ([]MemberResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v2/{project_id}/apigw/instances/{instance_id}/vpc-channels/{vpc_channel_id}/members
	raw, err := client.Put(client.ServiceURL("apigw", "instances", opts.GatewayID, "vpc-channels", opts.ChannelID, "members"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []MemberResp

	err = extract.IntoSlicePtr(raw.Body, &res, "members")
	return res, err
}
