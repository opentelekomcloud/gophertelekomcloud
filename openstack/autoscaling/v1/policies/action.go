package policies

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ActionOpts struct {
	ScalingPolicyId string
	// Specifies the operation for an AS policy.
	// execute: immediately executes the AS policy.
	// resume: enables the AS group.
	// pause: disables the AS group.
	Action string `json:"action"`
}

func PolicyAction(client *golangsdk.ServiceClient, opts ActionOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /autoscaling-api/v1/{project_id}/scaling_policy/{scaling_policy_id}/action
	_, err = client.Post(client.ServiceURL("scaling_policy", opts.ScalingPolicyId, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return
}
