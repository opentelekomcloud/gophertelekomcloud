package policies

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type BatchOpts struct {
	ScalingPolicyId []string `json:"scaling_policy_id" required:"true"`
	// Specifies an action to be performed on AS policies in batches. The options are as follows:
	// delete: deletes AS policies.
	// resume: enables AS policies.
	// pause: disables AS policies.
	Action string `json:"action" required:"true"`
	// Specifies whether to forcibly delete an AS policy. If the value is set to no, in-progress AS policies cannot be deleted.
	// Options:
	// no (default): indicates that the AS policy is not forcibly deleted.
	// yes: indicates that the AS policy is forcibly deleted.
	// This parameter is available only when action is set to delete.
	ForceDelete string `json:"force_delete,omitempty"`
	// Specifies whether to delete the alarm rule used by the alarm policy. The value can be yes or no (default).
	// This parameter is available only when action is set to delete.
	DeleteAlarm string `json:"delete_alarm,omitempty"`
}

func BatchAction(client *golangsdk.ServiceClient, opts BatchOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /autoscaling-api/v1/{project_id}/scaling_policies/action
	_, err = client.Post(client.ServiceURL("scaling_policies", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return
}
