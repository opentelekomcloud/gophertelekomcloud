package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ChangeNodeClassOpts struct {
	// Resource specification code of the new node class
	SpecCode string `json:"spec_code" required:"true"`
	// This parameter is not required if the instance has only one node group.
	// Each instance has one node group by default.
	// If you need to create multiple node groups, set this parameter to the ID of the group whose node class you want to change.
	GroupId string `json:"group_id,omitempty"`
	// AutoPay specification
	IsAutoPay bool `json:"is_auto_pay,omitempty"`
}

func ChangeNodeClass(client *golangsdk.ServiceClient, instanceId string, opts ChangeNodeClassOpts) (*ChangeNodeClassResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/instances/{instance_id}/flavor
	raw, err := client.Put(client.ServiceURL("instances", instanceId, "flavor"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ChangeNodeClassResponse
	return &res, extract.Into(raw.Body, &res)
}

type ChangeNodeClassResponse struct {
	// ID of the task for changing node class.
	JobID string `json:"job_id"`
}
