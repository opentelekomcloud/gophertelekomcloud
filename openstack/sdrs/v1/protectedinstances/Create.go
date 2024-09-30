package protectedinstances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts contains all the values needed to create a new instance.
type CreateOpts struct {
	// Group ID
	GroupID string `json:"server_group_id" required:"true"`
	// Server ID
	ServerID string `json:"server_id" required:"true"`
	// Instance Name
	Name string `json:"name" required:"true"`
	// Instance Description
	Description string `json:"description,omitempty"`
	// Subnet ID
	SubnetID string `json:"primary_subnet_id,omitempty"`
	// IP Address
	IpAddress string `json:"primary_ip_address,omitempty"`
	// Flavor ID of the DR site server
	Flavor string `json:"flavorRef,omitempty"`
}

// Create will create a new Instance based on the values in CreateOpts.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CreateProtectedInstanceResponse, error) {
	b, err := build.RequestBody(opts, "protected_instance")
	if err != nil {
		return nil, err
	}
	// POST /v1/{project_id}/protected-instances
	raw, err := client.Post(client.ServiceURL("protected-instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CreateProtectedInstanceResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CreateProtectedInstanceResponse struct {
	// 	Specifies the job ID. For details about the task execution result, see the description in Querying the Job Status at:
	// https://docs-beta.sc.otc.t-systems.com/storage-disaster-recovery-service/api-ref/sdrs_apis/job/querying_the_job_status.html
	JobID string `json:"job_id"`
}
