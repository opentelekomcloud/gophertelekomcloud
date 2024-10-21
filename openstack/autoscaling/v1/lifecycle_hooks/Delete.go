package lifecyclehooks

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// This function is used to delete a specified lifecycle hook.
// When a scaling action is being performed in an AS group, the lifecycle hooks of the AS group cannot be deleted.
func Delete(client *golangsdk.ServiceClient, asGroupId string, lifecycleHookName string) error {
	// DELETE /autoscaling-api/v1/{project_id}/scaling_lifecycle_hook/{scaling_group_id}/{lifecycle_hook_name}
	_, err := client.Delete(client.ServiceURL("scaling_lifecycle_hook", asGroupId, lifecycleHookName), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	if err != nil {
		return err
	}
	return nil
}
