package smart_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// PUT /v2/{project_id}/instances/{instance_id}/connector/tasks/{task_id}/resume
func RestartTask(client *golangsdk.ServiceClient, instanceId, taskId string) error {
	url := client.ServiceURL("instances", instanceId, "connector", "tasks", taskId, "resume")

	_, err := client.Put(url, nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	if err != nil {
		return err
	}

	return nil
}
