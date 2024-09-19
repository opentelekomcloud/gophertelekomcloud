package protectiongroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Specifies the information about a protection group.
	ServerGroup ServerGroupUpdateInfo `json:"server_group" required:"true"`
}

// UpdateOpts contains all the values needed to update a Group.
type ServerGroupUpdateInfo struct {
	// Group name
	Name string `json:"name" required:"true"`
}

// Update accepts a UpdateOpts struct and uses the values to update a Group.The response code from api is 200
func Update(client *golangsdk.ServiceClient, ServerGroupId string, opts UpdateOpts) (*ServerGroupResponseInfo, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	// PUT /v1/{project_id}/server-groups/{server_group_id}
	raw, err := client.Put(client.ServiceURL("server-groups", ServerGroupId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res UpdateResponse
	err = extract.Into(raw.Body, &res)
	return &res.ServerGroup, err
}

type UpdateResponse struct {
	// Specifies the information about a protection group.
	ServerGroup ServerGroupResponseInfo `json:"server_group"`
}
