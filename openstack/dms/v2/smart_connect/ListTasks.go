package smart_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type QueryTasksOpts struct {
	InstanceId string `json:"-"`
	Offset     *int   `q:"offset,omitempty"`
	Limit      *int   `q:"limit,omitempty"`
}

// This API is used to query Smart Connect tasks.
// GET /v2/{project_id}/instances/{instance_id}/connector/tasks
func ListTasks(client *golangsdk.ServiceClient, opts QueryTasksOpts) (*GetTaskResponce, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	url := client.ServiceURL("instances", opts.InstanceId, "connector", "tasks") + q.String()

	var res GetTaskResponce
	_, err = client.Get(url, &res, nil)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type GetTaskResponce struct {
	Tasks       []SmartConnectTaskEntity `json:"tasks"`
	TotalNumber int                      `json:"total_number"`
	MaxTasks    int                      `json:"max_tasks"`
	QuotaTasks  int                      `json:"quota_tasks"`
}

type SmartConnectTaskEntity struct {
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
