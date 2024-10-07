package lifecyclehooks

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to query details about a specified lifecycle hook by AS group ID and lifecycle hook name.
func Get(client *golangsdk.ServiceClient, asGroupId string, lifecycleHookName string) (*LifecycleHook, error) {
	// GET /autoscaling-api/v1/{project_id}/scaling_lifecycle_hook/{scaling_group_id}/{lifecycle_hook_name}
	raw, err := client.Get(client.ServiceURL("scaling_lifecycle_hook", asGroupId, lifecycleHookName), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res LifecycleHook
	err = extract.Into(raw.Body, &res)
	return &res, err
}
