package protectiongroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// DisableOpts contains all the values needed to disable protection for a Group.
type DisableOpts struct {
	// Disables protection for a protection group. The StopServerGroup is an empty object.
	StopServerGroup StopServerGroupInfo `json:"stop-server-group" required:"true"`
}

type StopServerGroupInfo struct {
	// Empty
}

// Disable will Disable protection for a protection Group.
func Disable(client *golangsdk.ServiceClient, ServerGroupId string) (*DisableResponse, error) {
	opts := DisableOpts{
		StopServerGroup: StopServerGroupInfo{},
	}
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	// POST /v1/{project_id}/server-groups/{server_group_id}/action
	raw, err := client.Post(client.ServiceURL("server-groups", ServerGroupId, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res DisableResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type DisableResponse struct {
	// 	Specifies the job ID.
	JobID string `json:"job_id"`
}
