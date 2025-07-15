package protectiongroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Create will create a new Group based on the values in CreateOpts.
func Delete(client *golangsdk.ServiceClient, ServerGroupId string) (*DeleteProtectionGroupResponse, error) {
	// DELETE /v1/{project_id}/server-groups/{server_group_id}
	raw, err := client.Delete(client.ServiceURL("server-groups", ServerGroupId), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res DeleteProtectionGroupResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type DeleteProtectionGroupResponse struct {
	// 	Specifies the job ID.
	JobID string `json:"job_id"`
}
