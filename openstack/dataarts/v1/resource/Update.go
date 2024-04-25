package resource

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// Update is used to modify a specific resource. When modifying the resource, specify the resource ID. The resource type and directory cannot be modified.
// Send request PUT /v1/{project_id}/resources/{resource_id}
func Update(client *golangsdk.ServiceClient, resourceId, workspace string, resource Resource) error {

	b, err := build.RequestBody(resource, "")
	if err != nil {
		return err
	}

	reqOpts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{HeaderContentType: ApplicationJson},
		OkCodes:     []int{204},
	}

	if workspace != "" {
		reqOpts.MoreHeaders = map[string]string{HeaderWorkspace: workspace}
	}

	_, err = client.Put(client.ServiceURL(resourcesEndpoint, resourceId), b, nil, reqOpts)
	if err != nil {
		return err
	}

	return nil
}
