package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type BatchOpts struct {
	// Specifies the ECS ID.
	Instances []string `json:"instances_id" required:"true"`
	// Specifies whether to delete an instance when it is removed from an AS group.
	// Options:
	// no (default): The instance will not be deleted.
	// yes: The instance will be deleted.
	// This parameter takes effect only when the action is set to REMOVE.
	IsDeleteEcs string `json:"instance_delete,omitempty"`
	// Specifies an action to be performed on instances in batches. The options are as follows:
	// ADD: adds instances to the AS group.
	// REMOVE: removes instances from the AS group.
	// PROTECT: enables instance protection.
	// UNPROTECT: disables instance protection.
	Action string `json:"action,omitempty"`
}

func BatchAction(client *golangsdk.ServiceClient, groupID string, opts BatchOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /autoscaling-api/v1/{project_id}/scaling_group_instance/{scaling_group_id}/action
	_, err = client.Post(client.ServiceURL("scaling_group_instance", groupID, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return
}
