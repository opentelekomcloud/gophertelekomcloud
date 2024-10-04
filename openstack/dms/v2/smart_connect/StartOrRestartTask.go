package smart_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// PUT /v2/{project_id}/kafka/instances/{instance_id}/connector/tasks/{task_id}/restart
func StartOrRestartTask(client *golangsdk.ServiceClient, instanceId, taskId string) error {
	url := client.ServiceURL("kafka", "instances", instanceId, "connector", "tasks", taskId, "restart")

	_, err := client.Put(url, nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	if err != nil {
		return err
	}

	return nil
}
