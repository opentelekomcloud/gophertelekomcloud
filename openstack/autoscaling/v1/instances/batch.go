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

func batch(client *golangsdk.ServiceClient, groupID string, opts BatchOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("scaling_group_instance", groupID, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}

func BatchAdd(client *golangsdk.ServiceClient, groupID string, instances []string) error {
	return batch(client, groupID, BatchOpts{
		Instances: instances,
		Action:    "ADD",
	})
}

func BatchDelete(client *golangsdk.ServiceClient, groupID string, instances []string, deleteEcs string) error {
	return batch(client, groupID, BatchOpts{
		Instances:   instances,
		IsDeleteEcs: deleteEcs,
		Action:      "REMOVE",
	})
}
