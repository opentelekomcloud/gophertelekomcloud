package lifecyclehooks

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to query lifecycle hooks by AS group ID.
func List(client *golangsdk.ServiceClient, asGroupId string) ([]LifecycleHook, error) {
	// GET https://{Endpoint}/autoscaling-api/v1/{project_id}/scaling_lifecycle_hook/e5d27f5c-dd76-4a61-b4bc-a67c5686719a/list
	raw, err := client.Get(client.ServiceURL("scaling_lifecycle_hook", asGroupId, "list"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListLifeCycleHooksResponse
	err = extract.Into(raw.Body, &res)
	return res.LifecycleHooks, err

}

type ListLifeCycleHooksResponse struct {
	// Specifies lifecycle hooks
	LifecycleHooks []LifecycleHook `json:"lifecycle_hooks"`
}
