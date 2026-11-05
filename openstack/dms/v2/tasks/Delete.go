package tasks

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete sends DELETE /v2/{project_id}/instances/{instance_id}/tasks/{task_id}
func Delete(client *golangsdk.ServiceClient, instanceID, taskID string) error {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "tasks", taskID).
		Build()
	if err != nil {
		return err
	}

	_, err = client.Delete(client.ServiceURL(url.String()), nil)
	return err
}
