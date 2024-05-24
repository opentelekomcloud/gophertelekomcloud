package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const systemTasksEndpoint = "system-tasks"

// GetSystemTask is used to query details about asynchronous tasks.
// Send request GET /v1/{project_id}/system-tasks/{task_id}
func GetSystemTask(client *golangsdk.ServiceClient, taskId, workspace string) (*SystemTaskResp, error) {

	var opts golangsdk.RequestOpts
	if workspace != "" {
		opts.MoreHeaders = map[string]string{HeaderWorkspace: workspace}
	}

	raw, err := client.Get(client.ServiceURL(systemTasksEndpoint, taskId), nil, &opts)
	if err != nil {
		return nil, err
	}

	var res SystemTaskResp
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type SystemTaskResp struct {
	// Task ID.
	Id string `json:"id"`
	// Name of the task.
	Name string `json:"name"`
	// Start time.
	StartTime int64 `json:"startTime"`
	// End time.
	EndTime int64 `json:"endTime,omitempty"`
	// Time when the task was last updated.
	LastUpdate int64 `json:"lastUpdate"`
	// Task status.
	//    RUNNING
	//    SUCCESSFUL
	//    FAILED
	Status string `json:"status"`
	// Project ID
	ProjectId string `json:"projectId,omitempty"`
	// Subtask.
	SubTasks []*SubTask `json:"subtasks,omitempty"`
}

type SubTask struct {
	// Subtask ID.
	Id string `json:"id"`
	// Name of the subtask.
	Name string `json:"name"`
	// Time when the task was last updated.
	LastUpdate int64 `json:"lastUpdate"`
	// Task status.
	//    RUNNING
	//    SUCCESSFUL
	//    FAILED
	Status string `json:"status"`
	// Task information.
	TaskId string `json:"taskId,omitempty"`
}
