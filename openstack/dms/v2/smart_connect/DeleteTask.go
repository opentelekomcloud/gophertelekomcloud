package smart_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// DELETE /v2/{project_id}/instances/{instance_id}/connector/tasks/{task_id}
func DeleteTask(client *golangsdk.ServiceClient, instanceId, taskId string) error {
	url := client.ServiceURL("instances", instanceId, "connector", "tasks", taskId)

	_, err := client.Delete(url, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	if err != nil {
		return err
	}

	return nil
}
