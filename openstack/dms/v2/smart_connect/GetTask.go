package smart_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func GetTask(client *golangsdk.ServiceClient, instanceId, taskId string) (*TaskDetails, error) {
	// GET /v2/{project_id}/instances/{instance_id}/connector/tasks/{task_id}
	url := client.ServiceURL("instances", instanceId, "connector", "tasks", taskId)

	var res TaskDetails
	_, err := client.Get(url, &res, nil)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type TaskDetails struct {
	TaskName    string                        `json:"task_name"`
	Topics      string                        `json:"topics"`
	TopicsRegex string                        `json:"topics_regex"`
	SourceType  string                        `json:"source_type"`
	SourceTask  *SmartConnectTaskSourceConfig `json:"source_task"`
	SinkType    string                        `json:"sink_type"`
	SinkTask    *SmartConnectTaskSinkConfig   `json:"sink_task"`
	ID          string                        `json:"id"`
	Status      string                        `json:"status"`
	CreateTime  int64                         `json:"create_time"`
}
