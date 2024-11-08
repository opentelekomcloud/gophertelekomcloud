package protectedinstances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// DeleteOpts contains all the values needed to delete an Instance.
type DeleteOpts struct {
	// Delete Target Server
	DeleteTargetServer *bool `json:"delete_target_server,omitempty"`
	// Delete Target Eip
	DeleteTargetEip *bool `json:"delete_target_eip,omitempty"`
}

// Delete will permanently delete a particular Instance based on its unique ID.
func Delete(client *golangsdk.ServiceClient, instanceId string, opts DeleteOpts) (*DeleteProtectedInstanceResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	// DELETE /v1/{project_id}/protected-instances/{protected_instance_id}
	raw, err := client.DeleteWithBodyResp(client.ServiceURL("protected-instances", instanceId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res DeleteProtectedInstanceResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type DeleteProtectedInstanceResponse struct {
	// 	Specifies the job ID. For details about the task execution result, see the description in Querying the Job Status at:
	// https://docs-beta.sc.otc.t-systems.com/storage-disaster-recovery-service/api-ref/sdrs_apis/job/querying_the_job_status.html
	JobID string `json:"job_id"`
}
