package resource

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete is used to delete a resource.
// Send request DELETE /v1/{project_id}/resources/{resource_id}
func Delete(client *golangsdk.ServiceClient, resourceId, workspace string) error {
	var err error
	reqOpts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{HeaderContentType: ApplicationJson},
		OkCodes:     []int{204},
	}

	if workspace != "" {
		reqOpts.MoreHeaders[HeaderWorkspace] = workspace
	}

	_, err = client.Delete(client.ServiceURL(resourcesEndpoint, resourceId), reqOpts)
	if err != nil {
		return err
	}

	return nil
}
