package protectiongroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// ServerGroupInfo contains all the values needed to create a new group.
type CreateOpts struct {
	// Group Name
	Name string `json:"name" required:"true"`
	// Group Description
	Description string `json:"description,omitempty"`
	// The source AZ of a protection group
	SourceAZ string `json:"source_availability_zone" required:"true"`
	// The target AZ of a protection group
	TargetAZ string `json:"target_availability_zone" required:"true"`
	// An active-active domain
	DomainID string `json:"domain_id" required:"true"`
	// ID of the source VPC
	SourceVpcID string `json:"source_vpc_id" required:"true"`
	// Deployment model
	DrType string `json:"dr_type,omitempty"`
}

// Create will create a new Group based on the values in CreateOpts.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CreateProtectionGroupResponse, error) {
	b, err := build.RequestBody(opts, "server_group")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/server-groups
	raw, err := client.Post(client.ServiceURL("server-groups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CreateProtectionGroupResponse
	err = extract.Into(raw.Body, &res)
	return &res, err

}

type CreateProtectionGroupResponse struct {
	// 	Specifies the job ID.
	JobID string `json:"job_id"`
}
