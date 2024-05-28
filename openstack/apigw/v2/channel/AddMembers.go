package channel

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateMemberOpts struct {
	GatewayID string `json:"-"`
	ChannelID string `json:"-"`
	// Backend instances of the VPC channel.
	Members []Members `json:"members,omitempty"`
}

func AddMembers(client *golangsdk.ServiceClient, opts CreateMemberOpts) ([]MemberResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/apigw/instances/{instance_id}/vpc-channels/{vpc_channel_id}/members
	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "vpc-channels", opts.ChannelID, "members"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res []MemberResp

	err = extract.IntoSlicePtr(raw.Body, &res, "members")
	return res, err
}

type MemberResp struct {
	// ID of the backend server group of the VPC channel.
	ID string `json:"id"`
	// VPC channel ID.
	ChannelID string `json:"vpc_channel_id"`
	// Backend server group ID.
	MemberGroupID string `json:"member_group_id"`
	// Name of the VPC channel's backend server group.
	// It can contain 3 to 64 characters, starting with a letter.
	// Only letters, digits, underscores (_), hyphens (-), and periods (.) are allowed.
	Name string `json:"member_group_name"`
	// Weight of the backend server group.
	// If the server group contains servers and a weight has been set for it, the weight is automatically used to assign weights to servers in this group.
	Weight int `json:"member_group_weight"`
	// Indicates whether the backend service is a standby node.
	IsBackup string `json:"is_backup"`
	// Backend server status.
	// 1: available
	// 2: unavailable
	Status string `json:"status"`
	// Backend server port.
	Port int `json:"port"`
	// Backend server ID.
	EcsId string `json:"ecs_id"`
	// Backend server name.
	EcsName string `json:"ecs_name"`
	// Time when the backend server group is created.
	CreatedAt string `json:"create_time"`
}
