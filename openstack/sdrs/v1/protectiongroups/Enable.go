package protectiongroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// EnableOpts contains all the values needed to enable protection for a Group.
type EnableOpts struct {
	// Enables protection for a protection group. The StartServerGroup is an empty object.
	StartServerGroup map[string]interface{} `json:"start-server-group" required:"true"`
}

// Enable will enable protection for a protection Group.
func Enable(client *golangsdk.ServiceClient, ServerGroupId string) (*EnableResponse, error) {
	opts := EnableOpts{StartServerGroup: map[string]interface{}{}}
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

	var res EnableResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type EnableResponse struct {
	// 	Specifies the job ID.
	JobID string `json:"job_id"`
}
